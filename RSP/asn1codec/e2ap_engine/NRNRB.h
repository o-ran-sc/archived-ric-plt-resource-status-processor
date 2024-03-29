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

#ifndef	_NRNRB_H_
#define	_NRNRB_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum NRNRB {
	NRNRB_nrb11	= 0,
	NRNRB_nrb18	= 1,
	NRNRB_nrb24	= 2,
	NRNRB_nrb25	= 3,
	NRNRB_nrb31	= 4,
	NRNRB_nrb32	= 5,
	NRNRB_nrb38	= 6,
	NRNRB_nrb51	= 7,
	NRNRB_nrb52	= 8,
	NRNRB_nrb65	= 9,
	NRNRB_nrb66	= 10,
	NRNRB_nrb78	= 11,
	NRNRB_nrb79	= 12,
	NRNRB_nrb93	= 13,
	NRNRB_nrb106	= 14,
	NRNRB_nrb107	= 15,
	NRNRB_nrb121	= 16,
	NRNRB_nrb132	= 17,
	NRNRB_nrb133	= 18,
	NRNRB_nrb135	= 19,
	NRNRB_nrb160	= 20,
	NRNRB_nrb162	= 21,
	NRNRB_nrb189	= 22,
	NRNRB_nrb216	= 23,
	NRNRB_nrb217	= 24,
	NRNRB_nrb245	= 25,
	NRNRB_nrb264	= 26,
	NRNRB_nrb270	= 27,
	NRNRB_nrb273	= 28
	/*
	 * Enumeration is extensible
	 */
} e_NRNRB;

/* NRNRB */
typedef long	 NRNRB_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_NRNRB_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_NRNRB;
extern const asn_INTEGER_specifics_t asn_SPC_NRNRB_specs_1;
asn_struct_free_f NRNRB_free;
asn_struct_print_f NRNRB_print;
asn_constr_check_f NRNRB_constraint;
ber_type_decoder_f NRNRB_decode_ber;
der_type_encoder_f NRNRB_encode_der;
xer_type_decoder_f NRNRB_decode_xer;
xer_type_encoder_f NRNRB_encode_xer;
per_type_decoder_f NRNRB_decode_uper;
per_type_encoder_f NRNRB_encode_uper;
per_type_decoder_f NRNRB_decode_aper;
per_type_encoder_f NRNRB_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _NRNRB_H_ */
#include "asn_internal.h"
