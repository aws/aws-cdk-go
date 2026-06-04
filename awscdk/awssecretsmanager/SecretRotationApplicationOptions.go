package awssecretsmanager


// Options for a SecretRotationApplication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretRotationApplicationOptions := &SecretRotationApplicationOptions{
//   	AdditionalSemanticVersions: map[string]*string{
//   		"additionalSemanticVersionsKey": jsii.String("additionalSemanticVersions"),
//   	},
//   	IsMultiUser: jsii.Boolean(false),
//   }
//
type SecretRotationApplicationOptions struct {
	// Semantic versions for partitions other than 'aws'.
	//
	// If not specified, it is assumed that non aws partitions (eg aws-cn, aws-us-gov) are not supported.
	// Default: - no additional partition versions (only 'aws' partition is supported).
	//
	AdditionalSemanticVersions *map[string]*string `field:"optional" json:"additionalSemanticVersions" yaml:"additionalSemanticVersions"`
	// Whether the rotation application uses the multi user scheme.
	// Default: false.
	//
	IsMultiUser *bool `field:"optional" json:"isMultiUser" yaml:"isMultiUser"`
}

