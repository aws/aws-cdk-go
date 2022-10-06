package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for creating an Amazon SageMaker endpoint configuration.
//
// Example:
//   tasks.NewSageMakerCreateEndpointConfig(this, jsii.String("SagemakerEndpointConfig"), &sageMakerCreateEndpointConfigProps{
//   	endpointConfigName: jsii.String("MyEndpointConfig"),
//   	productionVariants: []productionVariant{
//   		&productionVariant{
//   			initialInstanceCount: jsii.Number(2),
//   			instanceType: ec2.instanceType.of(ec2.instanceClass_M5, ec2.instanceSize_XLARGE),
//   			modelName: jsii.String("MyModel"),
//   			variantName: jsii.String("awesome-variant"),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
type SageMakerCreateEndpointConfigProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The name of the endpoint configuration.
	EndpointConfigName *string `field:"required" json:"endpointConfigName" yaml:"endpointConfigName"`
	// An list of ProductionVariant objects, one for each model that you want to host at this endpoint.
	//
	// Identifies a model that you want to host and the resources to deploy for hosting it.
	// If you are deploying multiple models, tell Amazon SageMaker how to distribute traffic among the models by specifying variant weights.
	ProductionVariants *[]*ProductionVariant `field:"required" json:"productionVariants" yaml:"productionVariants"`
	// AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Tags to be applied to the endpoint configuration.
	Tags awsstepfunctions.TaskInput `field:"optional" json:"tags" yaml:"tags"`
}

