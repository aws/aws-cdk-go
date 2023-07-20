package awslambda


// Properties for defining a `CfnCodeSigningConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeSigningConfigProps := &CfnCodeSigningConfigProps{
//   	AllowedPublishers: &AllowedPublishersProperty{
//   		SigningProfileVersionArns: []*string{
//   			jsii.String("signingProfileVersionArns"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	CodeSigningPolicies: &CodeSigningPoliciesProperty{
//   		UntrustedArtifactOnDeployment: jsii.String("untrustedArtifactOnDeployment"),
//   	},
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-codesigningconfig.html
//
type CfnCodeSigningConfigProps struct {
	// List of allowed publishers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-codesigningconfig.html#cfn-lambda-codesigningconfig-allowedpublishers
	//
	AllowedPublishers interface{} `field:"required" json:"allowedPublishers" yaml:"allowedPublishers"`
	// The code signing policy controls the validation failure action for signature mismatch or expiry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-codesigningconfig.html#cfn-lambda-codesigningconfig-codesigningpolicies
	//
	CodeSigningPolicies interface{} `field:"optional" json:"codeSigningPolicies" yaml:"codeSigningPolicies"`
	// Code signing configuration description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-codesigningconfig.html#cfn-lambda-codesigningconfig-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

