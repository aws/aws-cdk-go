package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplicationInferenceProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationInferenceProfileProps := &CfnApplicationInferenceProfileProps{
//   	InferenceProfileName: jsii.String("inferenceProfileName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ModelSource: &InferenceProfileModelSourceProperty{
//   		CopyFrom: jsii.String("copyFrom"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html
//
type CfnApplicationInferenceProfileProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-inferenceprofilename
	//
	InferenceProfileName *string `field:"required" json:"inferenceProfileName" yaml:"inferenceProfileName"`
	// Description of the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Various ways to encode a list of models in a CreateInferenceProfile request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-modelsource
	//
	ModelSource interface{} `field:"optional" json:"modelSource" yaml:"modelSource"`
	// List of Tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

