// Copyright 2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package amdtokens

/**
 * The following contains a description of the Attestation token for SEV-SNP
 */

type SnpToken struct {
	Evidence	SnpEvidence	`cbor:"89989,keyasint" json:"evidence"`
	ReferenceId	string		`cbor:"89988,keyasint" json:"reference-id"`
	TrustAnchorId	string		`cbor:"89987,keyasint" json:"trust-anchor-id"`
}

type SnpEvidence struct {
	Nonce			[]byte			`cbor:"89979,keyasint" json:"snp-nonce"`
	UniqueId		string			`cbor:"89978,keyasint" json:"snp-uniqueid"`
	Keys			SnpKeys			`cbor:"89977,keyasint" json:"snp-keys"`
	Core			SnpCoreComponents	`cbor:"89976,keyasint" json:"snp-core-components"`
	EvidenceTimestamp	string			`cbor:"89975,keyasint" json:"snp-evidence-timestamp"`
	SwInfo			SnpSwInfo		`cbor:"89974,keyasint" json:"snp-sw-info"`
	InstanceInfo		OciInstanceInfo		`cbor:"89972,keyasint" json:"snp-instance-info"`
}

type SnpKeys struct {
	Ark	string	`cbor:"89499,keyasint" json:"ark"`
	Ask	string	`cbor:"89498,keyasint" json:"ask"`
	Vcek	string	`cbor:"89497,keyasint" json:"vcek"`
}

type SnpCoreComponents struct {
	HwModel			string	`cbor:"89399,keyasint" json:"hw-model"`
	AttestationReport	[]byte	`cbor:"89398,keyasint" json:"attestation-report"`
}

type SnpSwInfo struct {
	Kernel	string	`cbor:"89349,keyasint" json:"hw-model"`
}
