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

#include "ABSInformation.h"

#include "ABSInformationFDD.h"
#include "ABSInformationTDD.h"
asn_per_constraints_t asn_PER_type_ABSInformation_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  2,  2,  0,  2 }	/* (0..2,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
asn_TYPE_member_t asn_MBR_ABSInformation_1[] = {
	{ ATF_POINTER, 0, offsetof(struct ABSInformation, choice.fdd),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ABSInformationFDD,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"fdd"
		},
	{ ATF_POINTER, 0, offsetof(struct ABSInformation, choice.tdd),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ABSInformationTDD,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"tdd"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct ABSInformation, choice.abs_inactive),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NULL,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"abs-inactive"
		},
};
static const asn_TYPE_tag2member_t asn_MAP_ABSInformation_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* fdd */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* tdd */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 } /* abs-inactive */
};
asn_CHOICE_specifics_t asn_SPC_ABSInformation_specs_1 = {
	sizeof(struct ABSInformation),
	offsetof(struct ABSInformation, _asn_ctx),
	offsetof(struct ABSInformation, present),
	sizeof(((struct ABSInformation *)0)->present),
	asn_MAP_ABSInformation_tag2el_1,
	3,	/* Count of tags in the map */
	0, 0,
	3	/* Extensions start */
};
asn_TYPE_descriptor_t asn_DEF_ABSInformation = {
	"ABSInformation",
	"ABSInformation",
	&asn_OP_CHOICE,
	0,	/* No effective tags (pointer) */
	0,	/* No effective tags (count) */
	0,	/* No tags (pointer) */
	0,	/* No tags (count) */
	{ 0, &asn_PER_type_ABSInformation_constr_1, CHOICE_constraint },
	asn_MBR_ABSInformation_1,
	3,	/* Elements count */
	&asn_SPC_ABSInformation_specs_1	/* Additional specs */
};

