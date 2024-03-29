/*
 *
 * Copyright 2019 AT&T Intellectual Property
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */


/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-PDU-Descriptions"
 * 	found in "../../asnFiles/e2ap-v031.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_SuccessfulOutcome_H_
#define	_SuccessfulOutcome_H_


#include "asn_application.h"

/* Including external dependencies */
#include "ProcedureCode.h"
#include "Criticality.h"
#include "ANY.h"
#include "asn_ioc.h"
#include "RICsubscriptionRequest.h"
#include "RICsubscriptionResponse.h"
#include "RICsubscriptionFailure.h"
#include "RICsubscriptionDeleteRequest.h"
#include "RICsubscriptionDeleteResponse.h"
#include "RICsubscriptionDeleteFailure.h"
#include "RICserviceUpdate.h"
#include "RICserviceUpdateAcknowledge.h"
#include "RICserviceUpdateFailure.h"
#include "RICcontrolRequest.h"
#include "RICcontrolAcknowledge.h"
#include "RICcontrolFailure.h"
#include "X2SetupRequest.h"
#include "X2SetupResponse.h"
#include "X2SetupFailure.h"
#include "ENDCX2SetupRequest.h"
#include "ENDCX2SetupResponse.h"
#include "ENDCX2SetupFailure.h"
#include "ResourceStatusRequest.h"
#include "ResourceStatusResponse.h"
#include "ResourceStatusFailure.h"
#include "ENBConfigurationUpdate.h"
#include "ENBConfigurationUpdateAcknowledge.h"
#include "ENBConfigurationUpdateFailure.h"
#include "ENDCConfigurationUpdate.h"
#include "ENDCConfigurationUpdateAcknowledge.h"
#include "ENDCConfigurationUpdateFailure.h"
#include "ResetRequest.h"
#include "ResetResponse.h"
#include "RICindication.h"
#include "RICserviceQuery.h"
#include "LoadInformation.h"
#include "GNBStatusIndication.h"
#include "ResourceStatusUpdate.h"
#include "ErrorIndication.h"
#include "OPEN_TYPE.h"
#include "constr_CHOICE.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum SuccessfulOutcome__value_PR {
	SuccessfulOutcome__value_PR_NOTHING,	/* No components present */
	SuccessfulOutcome__value_PR_RICsubscriptionResponse,
	SuccessfulOutcome__value_PR_RICsubscriptionDeleteResponse,
	SuccessfulOutcome__value_PR_RICserviceUpdateAcknowledge,
	SuccessfulOutcome__value_PR_RICcontrolAcknowledge,
	SuccessfulOutcome__value_PR_X2SetupResponse,
	SuccessfulOutcome__value_PR_ENDCX2SetupResponse,
	SuccessfulOutcome__value_PR_ResourceStatusResponse,
	SuccessfulOutcome__value_PR_ENBConfigurationUpdateAcknowledge,
	SuccessfulOutcome__value_PR_ENDCConfigurationUpdateAcknowledge,
	SuccessfulOutcome__value_PR_ResetResponse
} SuccessfulOutcome__value_PR;

/* SuccessfulOutcome */
typedef struct SuccessfulOutcome {
	ProcedureCode_t	 procedureCode;
	Criticality_t	 criticality;
	struct SuccessfulOutcome__value {
		SuccessfulOutcome__value_PR present;
		union SuccessfulOutcome__value_u {
			RICsubscriptionResponse_t	 RICsubscriptionResponse;
			RICsubscriptionDeleteResponse_t	 RICsubscriptionDeleteResponse;
			RICserviceUpdateAcknowledge_t	 RICserviceUpdateAcknowledge;
			RICcontrolAcknowledge_t	 RICcontrolAcknowledge;
			X2SetupResponse_t	 X2SetupResponse;
			ENDCX2SetupResponse_t	 ENDCX2SetupResponse;
			ResourceStatusResponse_t	 ResourceStatusResponse;
			ENBConfigurationUpdateAcknowledge_t	 ENBConfigurationUpdateAcknowledge;
			ENDCConfigurationUpdateAcknowledge_t	 ENDCConfigurationUpdateAcknowledge;
			ResetResponse_t	 ResetResponse;
		} choice;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} value;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} SuccessfulOutcome_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_SuccessfulOutcome;
extern asn_SEQUENCE_specifics_t asn_SPC_SuccessfulOutcome_specs_1;
extern asn_TYPE_member_t asn_MBR_SuccessfulOutcome_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _SuccessfulOutcome_H_ */
#include "asn_internal.h"
