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
 * From ASN.1 module "E2SM-gNB-X2-IEs"
 * 	found in "../../asnFiles/e2sm-gNB-X2-release-1-v041.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_InterfaceDirection_H_
#define	_InterfaceDirection_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum InterfaceDirection {
	InterfaceDirection_incoming	= 0,
	InterfaceDirection_outgoing	= 1
	/*
	 * Enumeration is extensible
	 */
} e_InterfaceDirection;

/* InterfaceDirection */
typedef long	 InterfaceDirection_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_InterfaceDirection_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_InterfaceDirection;
extern const asn_INTEGER_specifics_t asn_SPC_InterfaceDirection_specs_1;
asn_struct_free_f InterfaceDirection_free;
asn_struct_print_f InterfaceDirection_print;
asn_constr_check_f InterfaceDirection_constraint;
ber_type_decoder_f InterfaceDirection_decode_ber;
der_type_encoder_f InterfaceDirection_encode_der;
xer_type_decoder_f InterfaceDirection_decode_xer;
xer_type_encoder_f InterfaceDirection_encode_xer;
per_type_decoder_f InterfaceDirection_decode_uper;
per_type_encoder_f InterfaceDirection_encode_uper;
per_type_decoder_f InterfaceDirection_decode_aper;
per_type_encoder_f InterfaceDirection_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _InterfaceDirection_H_ */
#include "asn_internal.h"
