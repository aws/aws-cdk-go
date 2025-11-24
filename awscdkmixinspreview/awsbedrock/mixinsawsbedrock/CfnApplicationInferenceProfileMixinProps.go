package mixinsawsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationInferenceProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationInferenceProfileMixinProps := &CfnApplicationInferenceProfileMixinProps{
//   	Description: jsii.String("description"),
//   	InferenceProfileName: jsii.String("inferenceProfileName"),
//   	ModelSource: &InferenceProfileModelSourceProperty{
//   		CopyFrom: jsii.String("copyFrom"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html
//
type CfnApplicationInferenceProfileMixinProps struct {
	// The description of the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-inferenceprofilename
	//
	InferenceProfileName *string `field:"optional" json:"inferenceProfileName" yaml:"inferenceProfileName"`
	// Contains configurations for the inference profile to copy as the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-modelsource
	//
	ModelSource interface{} `field:"optional" json:"modelSource" yaml:"modelSource"`
	// A list of tags associated with the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

