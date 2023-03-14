package awssagemaker


// Model quality statistics and constraints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelQualityProperty := &ModelQualityProperty{
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
type CfnModelPackage_ModelQualityProperty struct {
	// Model quality constraints.
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// Model quality statistics.
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
}

