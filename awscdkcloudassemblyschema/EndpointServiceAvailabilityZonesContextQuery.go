package awscdkcloudassemblyschema


// Query to endpoint service context provider.
type EndpointServiceAvailabilityZonesContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Query service name.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

