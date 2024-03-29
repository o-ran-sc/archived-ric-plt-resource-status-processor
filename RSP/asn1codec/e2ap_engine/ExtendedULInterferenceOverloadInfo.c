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

#include "ExtendedULInterferenceOverloadInfo.h"

#include "ProtocolExtensionContainer.h"
static int
memb_associatedSubframes_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const BIT_STRING_t *st = (const BIT_STRING_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	if(st->size > 0) {
		/* Size in bits */
		size = 8 * st->size - (st->bits_unused & 0x07);
	} else {
		size = 0;
	}
	
	if((size == 5)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static asn_per_constraints_t asn_PER_memb_associatedSubframes_constr_2 CC_NOTUSED = {
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	{ APC_CONSTRAINED,	 0,  0,  5,  5 }	/* (SIZE(5..5)) */,
	0, 0	/* No PER value map */
};
asn_TYPE_member_t asn_MBR_ExtendedULInterferenceOverloadInfo_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct ExtendedULInterferenceOverloadInfo, associatedSubframes),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_BIT_STRING,
		0,
		{ 0, &asn_PER_memb_associatedSubframes_constr_2,  memb_associatedSubframes_constraint_1 },
		0, 0, /* No default value */
		"associatedSubframes"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct ExtendedULInterferenceOverloadInfo, extended_ul_InterferenceOverloadIndication),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_UL_InterferenceOverloadIndication,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"extended-ul-InterferenceOverloadIndication"
		},
	{ ATF_POINTER, 1, offsetof(struct ExtendedULInterferenceOverloadInfo, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_170P144,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"iE-Extensions"
		},
};
static const int asn_MAP_ExtendedULInterferenceOverloadInfo_oms_1[] = { 2 };
static const ber_tlv_tag_t asn_DEF_ExtendedULInterferenceOverloadInfo_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_ExtendedULInterferenceOverloadInfo_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* associatedSubframes */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* extended-ul-InterferenceOverloadIndication */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_ExtendedULInterferenceOverloadInfo_specs_1 = {
	sizeof(struct ExtendedULInterferenceOverloadInfo),
	offsetof(struct ExtendedULInterferenceOverloadInfo, _asn_ctx),
	asn_MAP_ExtendedULInterferenceOverloadInfo_tag2el_1,
	3,	/* Count of tags in the map */
	asn_MAP_ExtendedULInterferenceOverloadInfo_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	3,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_ExtendedULInterferenceOverloadInfo = {
	"ExtendedULInterferenceOverloadInfo",
	"ExtendedULInterferenceOverloadInfo",
	&asn_OP_SEQUENCE,
	asn_DEF_ExtendedULInterferenceOverloadInfo_tags_1,
	sizeof(asn_DEF_ExtendedULInterferenceOverloadInfo_tags_1)
		/sizeof(asn_DEF_ExtendedULInterferenceOverloadInfo_tags_1[0]), /* 1 */
	asn_DEF_ExtendedULInterferenceOverloadInfo_tags_1,	/* Same as above */
	sizeof(asn_DEF_ExtendedULInterferenceOverloadInfo_tags_1)
		/sizeof(asn_DEF_ExtendedULInterferenceOverloadInfo_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_ExtendedULInterferenceOverloadInfo_1,
	3,	/* Elements count */
	&asn_SPC_ExtendedULInterferenceOverloadInfo_specs_1	/* Additional specs */
};

