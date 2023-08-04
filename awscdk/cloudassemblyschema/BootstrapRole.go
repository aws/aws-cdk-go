package cloudassemblyschema


// Information needed to access an IAM role created as part of the bootstrap process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bootstrapRole := &BootstrapRole{
//   	Arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   	RequiresBootstrapStackVersion: jsii.Number(123),
//   }
//
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

