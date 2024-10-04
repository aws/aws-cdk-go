package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
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
	// A list of tags to add to the code signing configuration.
	//
	// > You must have the `lambda:TagResource` , `lambda:UntagResource` , and `lambda:ListTags` permissions for your [IAM principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) to manage the AWS CloudFormation stack. If you don't have these permissions, there might be unexpected behavior with stack-level tags propagating to the resource during resource creation and update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-codesigningconfig.html#cfn-lambda-codesigningconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

