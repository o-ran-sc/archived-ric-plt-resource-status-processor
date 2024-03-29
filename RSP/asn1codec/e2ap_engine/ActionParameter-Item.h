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

#ifndef	_ActionParameter_Item_H_
#define	_ActionParameter_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "ActionParameter-ID.h"
#include "ActionParameter-Value.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* ActionParameter-Item */
typedef struct ActionParameter_Item {
	ActionParameter_ID_t	 actionParameter_ID;
	ActionParameter_Value_t	 actionParameter_Value;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ActionParameter_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ActionParameter_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_ActionParameter_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_ActionParameter_Item_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _ActionParameter_Item_H_ */
#include "asn_internal.h"
