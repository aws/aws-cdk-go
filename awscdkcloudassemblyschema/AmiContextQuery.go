package awscdkcloudassemblyschema


// Query to AMI context provider.
type AmiContextQuery struct {
	// Account to query.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Filters to DescribeImages call.
	Filters *map[string]*[]*string `field:"required" json:"filters" yaml:"filters"`
	// Region to query.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Owners to DescribeImages call.
	// Default: - All owners.
	//
	Owners *[]*string `field:"optional" json:"owners" yaml:"owners"`
}

