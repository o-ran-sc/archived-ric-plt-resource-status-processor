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
 * From ASN.1 module "X2AP-IEs"
 * 	found in "../../asnFiles/X2AP-IEs.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_CauseRadioNetwork_H_
#define	_CauseRadioNetwork_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum CauseRadioNetwork {
	CauseRadioNetwork_handover_desirable_for_radio_reasons	= 0,
	CauseRadioNetwork_time_critical_handover	= 1,
	CauseRadioNetwork_resource_optimisation_handover	= 2,
	CauseRadioNetwork_reduce_load_in_serving_cell	= 3,
	CauseRadioNetwork_partial_handover	= 4,
	CauseRadioNetwork_unknown_new_eNB_UE_X2AP_ID	= 5,
	CauseRadioNetwork_unknown_old_eNB_UE_X2AP_ID	= 6,
	CauseRadioNetwork_unknown_pair_of_UE_X2AP_ID	= 7,
	CauseRadioNetwork_ho_target_not_allowed	= 8,
	CauseRadioNetwork_tx2relocoverall_expiry	= 9,
	CauseRadioNetwork_trelocprep_expiry	= 10,
	CauseRadioNetwork_cell_not_available	= 11,
	CauseRadioNetwork_no_radio_resources_available_in_target_cell	= 12,
	CauseRadioNetwork_invalid_MME_GroupID	= 13,
	CauseRadioNetwork_unknown_MME_Code	= 14,
	CauseRadioNetwork_encryption_and_or_integrity_protection_algorithms_not_supported	= 15,
	CauseRadioNetwork_reportCharacteristicsEmpty	= 16,
	CauseRadioNetwork_noReportPeriodicity	= 17,
	CauseRadioNetwork_existingMeasurementID	= 18,
	CauseRadioNetwork_unknown_eNB_Measurement_ID	= 19,
	CauseRadioNetwork_measurement_temporarily_not_available	= 20,
	CauseRadioNetwork_unspecified	= 21,
	/*
	 * Enumeration is extensible
	 */
	CauseRadioNetwork_load_balancing	= 22,
	CauseRadioNetwork_handover_optimisation	= 23,
	CauseRadioNetwork_value_out_of_allowed_range	= 24,
	CauseRadioNetwork_multiple_E_RAB_ID_instances	= 25,
	CauseRadioNetwork_switch_off_ongoing	= 26,
	CauseRadioNetwork_not_supported_QCI_value	= 27,
	CauseRadioNetwork_measurement_not_supported_for_the_object	= 28,
	CauseRadioNetwork_tDCoverall_expiry	= 29,
	CauseRadioNetwork_tDCprep_expiry	= 30,
	CauseRadioNetwork_action_desirable_for_radio_reasons	= 31,
	CauseRadioNetwork_reduce_load	= 32,
	CauseRadioNetwork_resource_optimisation	= 33,
	CauseRadioNetwork_time_critical_action	= 34,
	CauseRadioNetwork_target_not_allowed	= 35,
	CauseRadioNetwork_no_radio_resources_available	= 36,
	CauseRadioNetwork_invalid_QoS_combination	= 37,
	CauseRadioNetwork_encryption_algorithms_not_aupported	= 38,
	CauseRadioNetwork_procedure_cancelled	= 39,
	CauseRadioNetwork_rRM_purpose	= 40,
	CauseRadioNetwork_improve_user_bit_rate	= 41,
	CauseRadioNetwork_user_inactivity	= 42,
	CauseRadioNetwork_radio_connection_with_UE_lost	= 43,
	CauseRadioNetwork_failure_in_the_radio_interface_procedure	= 44,
	CauseRadioNetwork_bearer_option_not_supported	= 45,
	CauseRadioNetwork_mCG_Mobility	= 46,
	CauseRadioNetwork_sCG_Mobility	= 47,
	CauseRadioNetwork_count_reaches_max_value	= 48,
	CauseRadioNetwork_unknown_old_en_gNB_UE_X2AP_ID	= 49,
	CauseRadioNetwork_pDCP_Overload	= 50
} e_CauseRadioNetwork;

/* CauseRadioNetwork */
typedef long	 CauseRadioNetwork_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_CauseRadioNetwork_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_CauseRadioNetwork;
extern const asn_INTEGER_specifics_t asn_SPC_CauseRadioNetwork_specs_1;
asn_struct_free_f CauseRadioNetwork_free;
asn_struct_print_f CauseRadioNetwork_print;
asn_constr_check_f CauseRadioNetwork_constraint;
ber_type_decoder_f CauseRadioNetwork_decode_ber;
der_type_encoder_f CauseRadioNetwork_encode_der;
xer_type_decoder_f CauseRadioNetwork_decode_xer;
xer_type_encoder_f CauseRadioNetwork_encode_xer;
per_type_decoder_f CauseRadioNetwork_decode_uper;
per_type_encoder_f CauseRadioNetwork_encode_uper;
per_type_decoder_f CauseRadioNetwork_decode_aper;
per_type_encoder_f CauseRadioNetwork_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _CauseRadioNetwork_H_ */
#include "asn_internal.h"
