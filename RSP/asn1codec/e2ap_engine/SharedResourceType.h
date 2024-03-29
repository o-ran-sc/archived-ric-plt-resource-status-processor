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

#ifndef	_SharedResourceType_H_
#define	_SharedResourceType_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum SharedResourceType_PR {
	SharedResourceType_PR_NOTHING,	/* No components present */
	SharedResourceType_PR_uLOnlySharing,
	SharedResourceType_PR_uLandDLSharing
	/* Extensions may appear below */
	
} SharedResourceType_PR;

/* Forward declarations */
struct ULOnlySharing;
struct ULandDLSharing;

/* SharedResourceType */
typedef struct SharedResourceType {
	SharedResourceType_PR present;
	union SharedResourceType_u {
		struct ULOnlySharing	*uLOnlySharing;
		struct ULandDLSharing	*uLandDLSharing;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} SharedResourceType_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_SharedResourceType;
extern asn_CHOICE_specifics_t asn_SPC_SharedResourceType_specs_1;
extern asn_TYPE_member_t asn_MBR_SharedResourceType_1[2];
extern asn_per_constraints_t asn_PER_type_SharedResourceType_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _SharedResourceType_H_ */
#include "asn_internal.h"
