package awscdkcloudassemblyschema


// Query to SSM Parameter Context Provider.
type SSMParameterContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Parameter name to query.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

