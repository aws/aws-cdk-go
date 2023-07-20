package awssagemaker


// Docker container image configuration object for the model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelBiasAppSpecificationProperty := &ModelBiasAppSpecificationProperty{
//   	ConfigUri: jsii.String("configUri"),
//   	ImageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasappspecification.html
//
type CfnModelBiasJobDefinition_ModelBiasAppSpecificationProperty struct {
	// JSON formatted S3 file that defines bias parameters.
	//
	// For more information on this JSON configuration file, see [Configure bias parameters](https://docs.aws.amazon.com/sagemaker/latest/dg/clarify-config-json-monitor-bias-parameters.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasappspecification.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasappspecification-configuri
	//
	ConfigUri *string `field:"required" json:"configUri" yaml:"configUri"`
	// The container image to be run by the model bias job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasappspecification.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasappspecification-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Sets the environment variables in the Docker container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasappspecification.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasappspecification-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
}

