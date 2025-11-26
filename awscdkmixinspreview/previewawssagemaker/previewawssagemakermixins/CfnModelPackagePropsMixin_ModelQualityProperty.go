package previewawssagemakermixins


// Model quality statistics and constraints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modelQualityProperty := &ModelQualityProperty{
//   	Constraints: &MetricsSourceProperty{
//   		ContentDigest: jsii.String("contentDigest"),
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	Statistics: &MetricsSourceProperty{
//   		ContentDigest: jsii.String("contentDigest"),
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelquality.html
//
type CfnModelPackagePropsMixin_ModelQualityProperty struct {
	// Model quality constraints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelquality.html#cfn-sagemaker-modelpackage-modelquality-constraints
	//
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// Model quality statistics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelquality.html#cfn-sagemaker-modelpackage-modelquality-statistics
	//
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
}

