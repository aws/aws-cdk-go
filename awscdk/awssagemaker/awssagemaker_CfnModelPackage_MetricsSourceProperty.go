package awssagemaker


// Details about the metrics source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsSourceProperty := &MetricsSourceProperty{
//   	ContentType: jsii.String("contentType"),
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	ContentDigest: jsii.String("contentDigest"),
//   }
//
type CfnModelPackage_MetricsSourceProperty struct {
	// The metric source content type.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// The S3 URI for the metrics source.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The hash key used for the metrics source.
	ContentDigest *string `field:"optional" json:"contentDigest" yaml:"contentDigest"`
}

