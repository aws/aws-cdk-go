package awssagemaker


// An overview about the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelOverviewProperty := &ModelOverviewProperty{
//   	AlgorithmType: jsii.String("algorithmType"),
//   	InferenceEnvironment: &InferenceEnvironmentProperty{
//   		ContainerImage: []*string{
//   			jsii.String("containerImage"),
//   		},
//   	},
//   	ModelArtifact: []*string{
//   		jsii.String("modelArtifact"),
//   	},
//   	ModelCreator: jsii.String("modelCreator"),
//   	ModelDescription: jsii.String("modelDescription"),
//   	ModelId: jsii.String("modelId"),
//   	ModelName: jsii.String("modelName"),
//   	ModelOwner: jsii.String("modelOwner"),
//   	ModelVersion: jsii.Number(123),
//   	ProblemType: jsii.String("problemType"),
//   }
//
type CfnModelCard_ModelOverviewProperty struct {
	// The algorithm used to solve the problem.
	AlgorithmType *string `field:"optional" json:"algorithmType" yaml:"algorithmType"`
	// An overview about model inference.
	InferenceEnvironment interface{} `field:"optional" json:"inferenceEnvironment" yaml:"inferenceEnvironment"`
	// The location of the model artifact.
	ModelArtifact *[]*string `field:"optional" json:"modelArtifact" yaml:"modelArtifact"`
	// The creator of the model.
	ModelCreator *string `field:"optional" json:"modelCreator" yaml:"modelCreator"`
	// A description of the model.
	ModelDescription *string `field:"optional" json:"modelDescription" yaml:"modelDescription"`
	// The SageMaker Model ARN or non- SageMaker Model ID.
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// The name of the model.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The owner of the model.
	ModelOwner *string `field:"optional" json:"modelOwner" yaml:"modelOwner"`
	// The version of the model.
	ModelVersion *float64 `field:"optional" json:"modelVersion" yaml:"modelVersion"`
	// The problem being solved with the model.
	ProblemType *string `field:"optional" json:"problemType" yaml:"problemType"`
}

