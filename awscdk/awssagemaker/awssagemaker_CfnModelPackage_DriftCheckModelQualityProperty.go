package awssagemaker


// Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   driftCheckModelQualityProperty := &driftCheckModelQualityProperty{
//   	constraints: &metricsSourceProperty{
//   		contentType: jsii.String("contentType"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		contentDigest: jsii.String("contentDigest"),
//   	},
//   	statistics: &metricsSourceProperty{
//   		contentType: jsii.String("contentType"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		contentDigest: jsii.String("contentDigest"),
//   	},
//   }
//
type CfnModelPackage_DriftCheckModelQualityProperty struct {
	// `CfnModelPackage.DriftCheckModelQualityProperty.Constraints`.
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// `CfnModelPackage.DriftCheckModelQualityProperty.Statistics`.
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
}

