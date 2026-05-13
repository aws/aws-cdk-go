package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEvaluatorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var additionalModelRequestFields interface{}
//
//   cfnEvaluatorMixinProps := &CfnEvaluatorMixinProps{
//   	Description: jsii.String("description"),
//   	EvaluatorConfig: &EvaluatorConfigProperty{
//   		CodeBased: &CodeBasedEvaluatorConfigProperty{
//   			LambdaConfig: &LambdaEvaluatorConfigProperty{
//   				LambdaArn: jsii.String("lambdaArn"),
//   				LambdaTimeoutInSeconds: jsii.Number(123),
//   			},
//   		},
//   		LlmAsAJudge: &LlmAsAJudgeEvaluatorConfigProperty{
//   			Instructions: jsii.String("instructions"),
//   			ModelConfig: &EvaluatorModelConfigProperty{
//   				BedrockEvaluatorModelConfig: &BedrockEvaluatorModelConfigProperty{
//   					AdditionalModelRequestFields: additionalModelRequestFields,
//   					InferenceConfig: &InferenceConfigurationProperty{
//   						MaxTokens: jsii.Number(123),
//   						Temperature: jsii.Number(123),
//   						TopP: jsii.Number(123),
//   					},
//   					ModelId: jsii.String("modelId"),
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
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Level: jsii.String("level"),
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
type CfnEvaluatorMixinProps struct {
	// The description of the evaluator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration that defines how an evaluator assesses agent performance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-evaluatorconfig
	//
	EvaluatorConfig interface{} `field:"optional" json:"evaluatorConfig" yaml:"evaluatorConfig"`
	// The name of the evaluator.
	//
	// Must be unique within your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-evaluatorname
	//
	EvaluatorName *string `field:"optional" json:"evaluatorName" yaml:"evaluatorName"`
	// The ARN of the KMS key used to encrypt evaluator data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
	// A list of tags to assign to the evaluator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html#cfn-bedrockagentcore-evaluator-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

