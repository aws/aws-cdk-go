package awscdkcloudassemblyschema


// Query input for looking up a security group.
type SecurityGroupContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Security group id.
	// Default: - None.
	//
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// Security group name.
	// Default: - None.
	//
	SecurityGroupName *string `field:"optional" json:"securityGroupName" yaml:"securityGroupName"`
	// VPC ID.
	// Default: - None.
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

