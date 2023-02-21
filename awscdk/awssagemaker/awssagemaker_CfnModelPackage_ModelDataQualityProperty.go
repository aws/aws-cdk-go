package awssagemaker


// Data quality constraints and statistics for a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelDataQualityProperty := &modelDataQualityProperty{
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
type CfnModelPackage_ModelDataQualityProperty struct {
	// Data quality constraints for a model.
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// Data quality statistics for a model.
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
}

