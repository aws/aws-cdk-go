package awscdkcloudassemblyschema


// Query input for looking up a KMS Key.
type KeyContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Alias name used to search the Key.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

