package awssagemaker


// Docker container image configuration object for the model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelBiasAppSpecificationProperty := &modelBiasAppSpecificationProperty{
//   	configUri: jsii.String("configUri"),
//   	imageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_ModelBiasAppSpecificationProperty struct {
	// JSON formatted S3 file that defines bias parameters.
	//
	// For more information on this JSON configuration file, see [Configure bias parameters](https://docs.aws.amazon.com/sagemaker/latest/json-bias-parameter-config.html) .
	ConfigUri *string `field:"required" json:"configUri" yaml:"configUri"`
	// The container image to be run by the model bias job.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Sets the environment variables in the Docker container.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
}

