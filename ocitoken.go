// Copyright 2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package amdtokens

type OciInstanceInfo struct {
	AvailabilityDomain	string		`cbor:"89199,keyasint" json:"availabilityDomain"`
	CanonicalRegionName	string		`cbor:"89198,keyasint" json:"canonicalRegionName"`
	CompartmentId		string		`cbor:"89197,keyasint" json:"compartmentId"`
	DefinedTags		OciDefinedTags	`cbor:"89196,keyasint" json:"definedTags"`
	DisplayName		string		`cbor:"89195,keyasint" json:"displayName"`
	FaultDomain		string		`cbor:"89194,keyasint" json:"faultDomain"`
	Hostname		string		`cbor:"89193,keyasint" json:"hostname"`
	Id			string		`cbor:"89192,keyasint" json:"id"`
	Image			string		`cbor:"89191,keyasint" json:"image"`
	Metadata		OciMetadata	`cbor:"89190,keyasint" json:"metadata"`
	ociAdName		string		`cbor:"89189,keyasint" json:"ociAdName"`
	Region			string		`cbor:"89188,keyasint" json:"region"`
	RegionInfo		OciRegionInfo	`cbor:"89187,keyasint" json:"regionInfo"`
	Shape			string		`cbor:"89186,keyasint" json:"shape"`
	ShapeConfig		OciShapeConfig	`cbor:"89185,keyasint" json:"shapeConfig"`
	State			string		`cbor:"89184,keyasint" json:"state"`
	TimeCreated		float64		`cbor:"89183,keyasint" json:"timeCreated"`
}

type OciDefinedTags struct {
	Operations		OciOperations			`cbor:"89129,keyasint" json:"Operations"`
	OracleRecommendedTags	OciOracleRecommendedTags	`cbor:"89128,keyasint" json:"Oracle-Recommended-Tags"`
	OracleTags		OciOracleTags			`cbor:"89127,keyasint" json:"Oracle-Tags"`
}

type OciOperations struct {
	CreateBy		string	`cbor:"89099,keyasint" json:"CreateBy"`
	CreatedDateTime		string	`cbor:"89098,keyasint" json:"CreatedDateTime"`
}

type OciOracleRecommendedTags struct {
	ResourceOwner		string	`cbor:"89089,keyasint" json:"ResourceOwner"`
	ResourceType		string	`cbor:"89088,keyasint" json:"ResourceType"`
}

type OciOracleTags struct {
	CreatedBy		string	`cbor:"89079,keyasint" json:"CreatedBy"`
	CreatedOn		string	`cbor:"89078,keyasint" json:"CreatedOn"`
}

type OciMetadata struct {
	SshAuthorizedKeys	string	`cbor:"89069,keyasint" json:"ssh_authorized_keys"`
}

type OciRegionInfo struct {
	RealmDomainComponent	string	`cbor:"89059,keyasint" json:"realmDomainComponent"`
	RealmKey		string	`cbor:"89058,keyasint" json:"realmKey"`
	RegionIdentifier	string	`cbor:"89057,keyasint" json:"regionIdentifier"`
	RegionKey		string	`cbor:"89056,keyasint" json:"regionKey"`
}

type OciShapeConfig struct {
	BaselineOcpuUtilization		string	`cbor:"89049,keyasint" json:"baselineOcpuUtilization"`
	MaxVnicAttachments		float64	`cbor:"89048,keyasint" json:"maxVnicAttachments"`
	MemoryInGBs			float64	`cbor:"89047,keyasint" json:"memoryInGBs"`
	NetworkingBandwidthInGbps	float64	`cbor:"89046,keyasint" json:"networkingBandwidthInGbps"`
	Ocpus				float64	`cbor:"89045,keyasint" json:"ocpus"`
}
