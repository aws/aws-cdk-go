package mixinsawssagemaker


// Data quality constraints and statistics for a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modelDataQualityProperty := &ModelDataQualityProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modeldataquality.html
//
type CfnModelPackagePropsMixin_ModelDataQualityProperty struct {
	// Data quality constraints for a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modeldataquality.html#cfn-sagemaker-modelpackage-modeldataquality-constraints
	//
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// Data quality statistics for a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modeldataquality.html#cfn-sagemaker-modelpackage-modeldataquality-statistics
	//
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
}

