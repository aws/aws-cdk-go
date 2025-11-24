package mixinsawssagemaker


// Represents the drift check data quality baselines that can be used when the model monitor is set using the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   driftCheckModelDataQualityProperty := &DriftCheckModelDataQualityProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckmodeldataquality.html
//
type CfnModelPackagePropsMixin_DriftCheckModelDataQualityProperty struct {
	// The drift check model data quality constraints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckmodeldataquality.html#cfn-sagemaker-modelpackage-driftcheckmodeldataquality-constraints
	//
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// The drift check model data quality statistics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckmodeldataquality.html#cfn-sagemaker-modelpackage-driftcheckmodeldataquality-statistics
	//
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
}

