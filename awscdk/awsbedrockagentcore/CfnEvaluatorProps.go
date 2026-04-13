package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEvaluator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//
//   cfnEvaluatorProps := &CfnEvaluatorProps{
//   	EvaluatorConfig: &EvaluatorConfigProperty{
//   		CodeBased: &CodeBasedEvaluatorConfigProperty{
//   			LambdaConfig: &LambdaEvaluatorConfigProperty{
//   				LambdaArn: jsii.String("lambdaArn"),
//
//   				// the properties below are optional
//   				LambdaTimeoutInSeconds: jsii.Number(123),
//   			},
//   		},
//   		LlmAsAJudge: &LlmAsAJudgeEvaluatorConfigProperty{
//   			Instructions: jsii.String("instructions"),
//   			ModelConfig: &EvaluatorModelConfigProperty{
//   				BedrockEvaluatorModelConfig: &BedrockEvaluatorModelConfigProperty{
//   					ModelId: jsii.String("modelId"),
//
//   					// the properties below are optional
//   					AdditionalModelRequestFields: additionalModelRequestFields,
//   					InferenceConfig: &InferenceConfigurationProperty{
//   						MaxTokens: jsii.Number(123),
//   						Temperature: jsii.Number(123),
//   						TopP: jsii.Number(123),
//   					},
//   				},
//   			},
//   			RatingScale: &RatingScaleProperty{
//   				Categorical: []interface{}{
//   					&CategoricalScaleDefinitionProperty{
//   						Definition: jsii.String("definition"),
//   						Label: jsii.String("label"),
//   					},
//   				},
//   				Numerical: []interface{}{
//   					&NumericalScaleDefinitionProperty{
//   						Definition: jsii.String("definition"),
//   						Label: jsii.String("label"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	EvaluatorName: jsii.String("evaluatorName"),
//   	Level: jsii.String("level"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html
//
type CfnEvaluatorProps struct {
	// The configuration that defines how an evaluator assesses agent performance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-evaluatorconfig
	//
	EvaluatorConfig interface{} `field:"required" json:"evaluatorConfig" yaml:"evaluatorConfig"`
	// The name of the evaluator.
	//
	// Must be unique within your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-evaluatorname
	//
	EvaluatorName *string `field:"required" json:"evaluatorName" yaml:"evaluatorName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-level
	//
	Level *string `field:"required" json:"level" yaml:"level"`
	// The description of the evaluator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of tags to assign to the evaluator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

