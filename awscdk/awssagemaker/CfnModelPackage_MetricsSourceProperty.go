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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-metricssource.html
//
type CfnModelPackage_MetricsSourceProperty struct {
	// The metric source content type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-metricssource.html#cfn-sagemaker-modelpackage-metricssource-contenttype
	//
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// The S3 URI for the metrics source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-metricssource.html#cfn-sagemaker-modelpackage-metricssource-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The hash key used for the metrics source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-metricssource.html#cfn-sagemaker-modelpackage-metricssource-contentdigest
	//
	ContentDigest *string `field:"optional" json:"contentDigest" yaml:"contentDigest"`
}

