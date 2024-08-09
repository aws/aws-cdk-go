package awscdkcloudassemblyschema


// Information needed to access an IAM role created as part of the bootstrap process.
type BootstrapRole struct {
	// The ARN of the IAM role created as part of bootrapping e.g. lookupRoleArn.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// External ID to use when assuming the bootstrap role.
	// Default: - No external ID.
	//
	AssumeRoleExternalId *string `field:"optional" json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// Name of SSM parameter with bootstrap stack version.
	// Default: - Discover SSM parameter by reading stack.
	//
	BootstrapStackVersionSsmParameter *string `field:"optional" json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// Version of bootstrap stack required to use this role.
	// Default: - No bootstrap stack required.
	//
	RequiresBootstrapStackVersion *float64 `field:"optional" json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
}

