package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for creating an Amazon SageMaker model.
//
// Example:
//   tasks.NewSageMakerCreateModel(this, jsii.String("Sagemaker"), &SageMakerCreateModelProps{
//   	ModelName: jsii.String("MyModel"),
//   	PrimaryContainer: tasks.NewContainerDefinition(&ContainerDefinitionOptions{
//   		Image: tasks.DockerImage_FromJsonExpression(sfn.JsonPath_StringAt(jsii.String("$.Model.imageName"))),
//   		Mode: tasks.Mode_SINGLE_MODEL,
//   		ModelS3Location: tasks.S3Location_FromJsonExpression(jsii.String("$.TrainingJob.ModelArtifacts.S3ModelArtifacts")),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
// Experimental.
type SageMakerCreateModelProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The name of the new model.
	// Experimental.
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// The definition of the primary docker image containing inference code, associated artifacts, and custom environment map that the inference code uses when the model is deployed for predictions.
	// Experimental.
	PrimaryContainer IContainerDefinition `field:"required" json:"primaryContainer" yaml:"primaryContainer"`
	// Specifies the containers in the inference pipeline.
	// Experimental.
	Containers *[]IContainerDefinition `field:"optional" json:"containers" yaml:"containers"`
	// Isolates the model container.
	//
	// No inbound or outbound network calls can be made to or from the model container.
	// Experimental.
	EnableNetworkIsolation *bool `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// An execution role that you can pass in a CreateModel API request.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The subnets of the VPC to which the hosted model is connected (Note this parameter is only used when VPC is provided).
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// Tags to be applied to the model.
	// Experimental.
	Tags awsstepfunctions.TaskInput `field:"optional" json:"tags" yaml:"tags"`
	// The VPC that is accessible by the hosted model.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

