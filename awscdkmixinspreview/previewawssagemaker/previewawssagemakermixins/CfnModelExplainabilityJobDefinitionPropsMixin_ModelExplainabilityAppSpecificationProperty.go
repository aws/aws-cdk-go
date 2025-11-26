package previewawssagemakermixins


// Docker container image configuration object for the model explainability job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modelExplainabilityAppSpecificationProperty := &ModelExplainabilityAppSpecificationProperty{
//   	ConfigUri: jsii.String("configUri"),
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ImageUri: jsii.String("imageUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityappspecification.html
//
type CfnModelExplainabilityJobDefinitionPropsMixin_ModelExplainabilityAppSpecificationProperty struct {
	// JSON formatted Amazon S3 file that defines explainability parameters.
	//
	// For more information on this JSON configuration file, see [Configure model explainability parameters](https://docs.aws.amazon.com/sagemaker/latest/dg/clarify-config-json-monitor-model-explainability-parameters.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityappspecification.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityappspecification-configuri
	//
	ConfigUri *string `field:"optional" json:"configUri" yaml:"configUri"`
	// Sets the environment variables in the Docker container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityappspecification.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityappspecification-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The container image to be run by the model explainability job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityappspecification.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityappspecification-imageuri
	//
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
}

