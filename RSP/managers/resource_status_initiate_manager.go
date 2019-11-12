package managers

import (
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
	"rsp/e2pdus"
	"rsp/enums"
	"rsp/logger"
	"rsp/models"
	"rsp/rmrcgo"
	"rsp/rsperrors"
	"rsp/services"
	"rsp/services/rmrsender"
)

type ResourceStatusInitiateManager struct {
	logger          *logger.Logger
	rnibDataService services.RNibDataService
	rmrSender       *rmrsender.RmrSender
}

type IResourceStatusInitiateManager interface {
	Execute(inventoryName string, resourceStatusInitiateRequestParams *e2pdus.ResourceStatusRequestData) error
}

func NewResourceStatusInitiateManager(logger *logger.Logger, rnibDataService services.RNibDataService, rmrSender *rmrsender.RmrSender) *ResourceStatusInitiateManager {
	return &ResourceStatusInitiateManager{
		logger:          logger,
		rnibDataService: rnibDataService,
		rmrSender:       rmrSender,
	}
}

func (m *ResourceStatusInitiateManager) Execute(inventoryName string, resourceStatusInitiateRequestParams *e2pdus.ResourceStatusRequestData) error {

	nodebInfo, err := m.rnibDataService.GetNodeb(inventoryName)

	if err != nil {
		m.logger.Errorf("#ResourceStatusInitiateManager.Execute - RAN name: %s - Error fetching RAN from rNib: %v", inventoryName, err)
		return rsperrors.NewRnibDbError()
	}

	m.logger.Infof("#ResourceStatusInitiateManager.Execute - RAN name: %s, connection status: %s", nodebInfo.GetRanName(), nodebInfo.GetConnectionStatus())

	if nodebInfo.GetConnectionStatus() != entities.ConnectionStatus_CONNECTED {
		m.logger.Errorf("#ResourceStatusInitiateManager.Execute - RAN name: %s - RAN's connection status isn't CONNECTED", inventoryName)
		return rsperrors.NewWrongStateError("Resource Status Initiate", entities.ConnectionStatus_name[int32(nodebInfo.ConnectionStatus)])
	}

	m.sendResourceStatusInitiatePerCell(nodebInfo, *resourceStatusInitiateRequestParams)

	return nil
}

func (m *ResourceStatusInitiateManager) sendResourceStatusInitiatePerCell(nodebInfo *entities.NodebInfo, requestParams e2pdus.ResourceStatusRequestData) {
	enb, _ := nodebInfo.Configuration.(*entities.NodebInfo_Enb)
	cells := enb.Enb.ServedCells

	for index, cellInfo := range cells {
		requestParams.CellID = cellInfo.CellId
		requestParams.MeasurementID = e2pdus.Measurement_ID(index + 1)

		m.logger.Infof("#ResourceStatusInitiateManager.sendResourceStatusInitiatePerCell - RAN name: %s, Going to send request for cell id %s, measurement id %d", nodebInfo.RanName, requestParams.CellID, requestParams.MeasurementID)

		payload, payloadAsString, err := e2pdus.BuildPackedResourceStatusRequest(enums.Registration_Request_start, &requestParams, e2pdus.MaxAsn1PackedBufferSize, e2pdus.MaxAsn1CodecMessageBufferSize, m.logger.DebugEnabled())
		if err != nil {
			m.logger.Errorf("#ResourceStatusInitiateManager.sendResourceStatusInitiatePerCell - RAN name: %s. Failed to build and pack the resource status initiate request for cell id %s, measurement id %d. error: %s", nodebInfo.RanName, requestParams.CellID, requestParams.MeasurementID, err)
			continue
		}

		m.logger.Debugf("#ResourceStatusInitiateManager.sendResourceStatusInitiatePerCell - RAN name: %s, cell id: %s, measurement id: %d, payload: %s", nodebInfo.RanName, requestParams.CellID, requestParams.MeasurementID, payloadAsString)

		rmrMsg := models.NewRmrMessage(rmrcgo.RicResStatusReq, nodebInfo.RanName, payload)
		go m.rmrSender.Send(rmrMsg)
	}
}
