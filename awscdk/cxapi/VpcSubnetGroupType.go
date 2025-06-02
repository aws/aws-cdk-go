package cxapi


// The type of subnet group.
//
// Same as SubnetType in the aws-cdk-lib/aws-ec2 package,
// but we can't use that because of cyclical dependencies.
type VpcSubnetGroupType string

const (
	// Public subnet group type.
	VpcSubnetGroupType_PUBLIC VpcSubnetGroupType = "PUBLIC"
	// Private subnet group type.
	VpcSubnetGroupType_PRIVATE VpcSubnetGroupType = "PRIVATE"
	// Isolated subnet group type.
	VpcSubnetGroupType_ISOLATED VpcSubnetGroupType = "ISOLATED"
)

