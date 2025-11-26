package previewawssagemakermixins


// Represents the drift check explainability baselines that can be used when the model monitor is set using the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   driftCheckExplainabilityProperty := &DriftCheckExplainabilityProperty{
//   	ConfigFile: &FileSourceProperty{
//   		ContentDigest: jsii.String("contentDigest"),
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	Constraints: &MetricsSourceProperty{
//   		ContentDigest: jsii.String("contentDigest"),
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckexplainability.html
//
type CfnModelPackagePropsMixin_DriftCheckExplainabilityProperty struct {
	// The explainability config file for the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckexplainability.html#cfn-sagemaker-modelpackage-driftcheckexplainability-configfile
	//
	ConfigFile interface{} `field:"optional" json:"configFile" yaml:"configFile"`
	// The drift check explainability constraints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckexplainability.html#cfn-sagemaker-modelpackage-driftcheckexplainability-constraints
	//
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
}

