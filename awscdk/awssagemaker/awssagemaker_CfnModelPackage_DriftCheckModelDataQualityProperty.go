package awssagemaker


// Represents the drift check data quality baselines that can be used when the model monitor is set using the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   driftCheckModelDataQualityProperty := &DriftCheckModelDataQualityProperty{
//   	Constraints: &MetricsSourceProperty{
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		ContentDigest: jsii.String("contentDigest"),
//   	},
//   	Statistics: &MetricsSourceProperty{
//   		ContentType: jsii.String("contentType"),
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		ContentDigest: jsii.String("contentDigest"),
//   	},
//   }
//
type CfnModelPackage_DriftCheckModelDataQualityProperty struct {
	// The drift check model data quality constraints.
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// The drift check model data quality statistics.
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
}

