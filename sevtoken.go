// Copyright 2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package amdtokens

/**
 * The following contains a description of the Attestation Token for SEV
 */

type SevToken struct {
	Evidence	SevEvidence	`cbor:"99989,keyasint" json:"evidence"`
	ReferenceId	string		`cbor:"99988,keyasint" json:"reference-id"`
	TrustAnchorId	string		`cbor:"99987,keyasint" json:"trust-anchor-id"`
}

type SevEvidence struct {
	Nonce			[]byte			`cbor:"99979,keyasint" json:"sev-nonce"`
	UniqueId		string			`cbor:"99978,keyasint" json:"sev-uniqueid"`
	Keys			SevKeys			`cbor:"99977,keyasint" json:"sev-keys"`
	Core			SevCoreComponents	`cbor:"99976,keyasint" json:"sev-core-components"`
	SwInfo			SevSwInfo		`cbor:"99975,keyasint" json:"sev-sw-info"`
	InstanceInfo		OciInstanceInfo		`cbor:"99973,keyasint" json:"sev-instance-info"`
}

type SevKeys struct {
	Cek	[]byte	`cbor:"99499,keyasint" json:"cek"`
	Pek	[]byte	`cbor:"99498,keyasint" json:"pek"`
	Oca	[]byte	`cbor:"99497,keyasint" json:"oca"`
	Pdh	[]byte	`cbor:"99496,keyasint" json:"pdh"`
}

type SevCoreComponents struct {
	HwModel			string	`cbor:"99399,keyasint" json:"hw-model"`
	AttestationReport	[]byte	`cbor:"99398,keyasint" json:"attestation-report"`
}

type SevSwInfo struct {
	Kernel	string	`cbor:"99349,keyasint" json:"hw-model"`
}
