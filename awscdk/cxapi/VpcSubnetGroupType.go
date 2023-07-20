package cxapi


// The type of subnet group.
//
// Same as SubnetType in the.
type VpcSubnetGroupType string

const (
	// Public subnet group type.
	VpcSubnetGroupType_PUBLIC VpcSubnetGroupType = "PUBLIC"
	// Private subnet group type.
	VpcSubnetGroupType_PRIVATE VpcSubnetGroupType = "PRIVATE"
	// Isolated subnet group type.
	VpcSubnetGroupType_ISOLATED VpcSubnetGroupType = "ISOLATED"
)

