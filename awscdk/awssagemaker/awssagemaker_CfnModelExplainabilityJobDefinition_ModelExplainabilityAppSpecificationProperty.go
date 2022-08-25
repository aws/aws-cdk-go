package awssagemaker


// Docker container image configuration object for the model explainability job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelExplainabilityAppSpecificationProperty := &modelExplainabilityAppSpecificationProperty{
//   	configUri: jsii.String("configUri"),
//   	imageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_ModelExplainabilityAppSpecificationProperty struct {
	// JSON formatted S3 file that defines explainability parameters.
	//
	// For more information on this JSON configuration file, see [Configure model explainability parameters](https://docs.aws.amazon.com/sagemaker/latest/json-model-explainability-parameter-config.html) .
	ConfigUri *string `field:"required" json:"configUri" yaml:"configUri"`
	// The container image to be run by the model explainability job.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Sets the environment variables in the Docker container.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
}

