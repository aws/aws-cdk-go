package previewawssagemakermixins


// Docker container image configuration object for the model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modelBiasAppSpecificationProperty := &ModelBiasAppSpecificationProperty{
//   	ConfigUri: jsii.String("configUri"),
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ImageUri: jsii.String("imageUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasappspecification.html
//
type CfnModelBiasJobDefinitionPropsMixin_ModelBiasAppSpecificationProperty struct {
	// JSON formatted S3 file that defines bias parameters.
	//
	// For more information on this JSON configuration file, see [Configure bias parameters](https://docs.aws.amazon.com/sagemaker/latest/dg/clarify-config-json-monitor-bias-parameters.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasappspecification.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasappspecification-configuri
	//
	ConfigUri *string `field:"optional" json:"configUri" yaml:"configUri"`
	// Sets the environment variables in the Docker container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasappspecification.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasappspecification-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The container image to be run by the model bias job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasappspecification.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasappspecification-imageuri
	//
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
}

