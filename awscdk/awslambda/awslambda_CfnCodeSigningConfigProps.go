package awslambda


// Properties for defining a `CfnCodeSigningConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeSigningConfigProps := &cfnCodeSigningConfigProps{
//   	allowedPublishers: &allowedPublishersProperty{
//   		signingProfileVersionArns: []*string{
//   			jsii.String("signingProfileVersionArns"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	codeSigningPolicies: &codeSigningPoliciesProperty{
//   		untrustedArtifactOnDeployment: jsii.String("untrustedArtifactOnDeployment"),
//   	},
//   	description: jsii.String("description"),
//   }
//
type CfnCodeSigningConfigProps struct {
	// List of allowed publishers.
	AllowedPublishers interface{} `field:"required" json:"allowedPublishers" yaml:"allowedPublishers"`
	// The code signing policy controls the validation failure action for signature mismatch or expiry.
	CodeSigningPolicies interface{} `field:"optional" json:"codeSigningPolicies" yaml:"codeSigningPolicies"`
	// Code signing configuration description.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

