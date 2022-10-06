package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsSourceProperty := &metricsSourceProperty{
//   	contentType: jsii.String("contentType"),
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	contentDigest: jsii.String("contentDigest"),
//   }
//
type CfnModelPackage_MetricsSourceProperty struct {
	// `CfnModelPackage.MetricsSourceProperty.ContentType`.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// `CfnModelPackage.MetricsSourceProperty.S3Uri`.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// `CfnModelPackage.MetricsSourceProperty.ContentDigest`.
	ContentDigest *string `field:"optional" json:"contentDigest" yaml:"contentDigest"`
}

