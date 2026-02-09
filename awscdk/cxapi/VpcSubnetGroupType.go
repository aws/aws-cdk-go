package cxapi


// The type of subnet group.
//
// Same as SubnetType in the aws-cdk-lib/aws-ec2 package,
// but we can't use that because of cyclical dependencies.
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type VpcSubnetGroupType string

const (
	// Public subnet group type.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	VpcSubnetGroupType_PUBLIC VpcSubnetGroupType = "PUBLIC"
	// Private subnet group type.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	VpcSubnetGroupType_PRIVATE VpcSubnetGroupType = "PRIVATE"
	// Isolated subnet group type.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	VpcSubnetGroupType_ISOLATED VpcSubnetGroupType = "ISOLATED"
)

