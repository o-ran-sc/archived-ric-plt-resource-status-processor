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

#include "SharedResourceType.h"

#include "ULOnlySharing.h"
#include "ULandDLSharing.h"
asn_per_constraints_t asn_PER_type_SharedResourceType_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  1,  1,  0,  1 }	/* (0..1,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
asn_TYPE_member_t asn_MBR_SharedResourceType_1[] = {
	{ ATF_POINTER, 0, offsetof(struct SharedResourceType, choice.uLOnlySharing),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ULOnlySharing,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"uLOnlySharing"
		},
	{ ATF_POINTER, 0, offsetof(struct SharedResourceType, choice.uLandDLSharing),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ULandDLSharing,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"uLandDLSharing"
		},
};
static const asn_TYPE_tag2member_t asn_MAP_SharedResourceType_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* uLOnlySharing */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 } /* uLandDLSharing */
};
asn_CHOICE_specifics_t asn_SPC_SharedResourceType_specs_1 = {
	sizeof(struct SharedResourceType),
	offsetof(struct SharedResourceType, _asn_ctx),
	offsetof(struct SharedResourceType, present),
	sizeof(((struct SharedResourceType *)0)->present),
	asn_MAP_SharedResourceType_tag2el_1,
	2,	/* Count of tags in the map */
	0, 0,
	2	/* Extensions start */
};
asn_TYPE_descriptor_t asn_DEF_SharedResourceType = {
	"SharedResourceType",
	"SharedResourceType",
	&asn_OP_CHOICE,
	0,	/* No effective tags (pointer) */
	0,	/* No effective tags (count) */
	0,	/* No tags (pointer) */
	0,	/* No tags (count) */
	{ 0, &asn_PER_type_SharedResourceType_constr_1, CHOICE_constraint },
	asn_MBR_SharedResourceType_1,
	2,	/* Elements count */
	&asn_SPC_SharedResourceType_specs_1	/* Additional specs */
};

