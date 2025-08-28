package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplicationInferenceProfile`.
//
// Example:
//   // Create or reference an existing L1 CfnApplicationInferenceProfile
//   cfnProfile := awscdk.Aws_bedrock.NewCfnApplicationInferenceProfile(this, jsii.String("CfnProfile"), &CfnApplicationInferenceProfileProps{
//   	InferenceProfileName: jsii.String("my-cfn-profile"),
//   	ModelSource: &InferenceProfileModelSourceProperty{
//   		CopyFrom: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0().InvokableArn,
//   	},
//   	Description: jsii.String("Profile created via L1 construct"),
//   })
//
//   // Import the L1 construct as an L2 ApplicationInferenceProfile
//   importedFromCfn := bedrock.ApplicationInferenceProfile_FromCfnApplicationInferenceProfile(cfnProfile)
//
//   // Grant permissions to use the imported profile
//   lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_11(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("def handler(event, context): return \"Hello\"")),
//   })
//
//   importedFromCfn.GrantProfileUsage(lambdaFunction)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html
//
type CfnApplicationInferenceProfileProps struct {
	// The name of the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-inferenceprofilename
	//
	InferenceProfileName *string `field:"required" json:"inferenceProfileName" yaml:"inferenceProfileName"`
	// The description of the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Contains configurations for the inference profile to copy as the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-modelsource
	//
	ModelSource interface{} `field:"optional" json:"modelSource" yaml:"modelSource"`
	// A list of tags associated with the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

