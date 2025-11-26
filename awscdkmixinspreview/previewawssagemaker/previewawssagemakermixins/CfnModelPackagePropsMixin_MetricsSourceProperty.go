package previewawssagemakermixins


// Details about the metrics source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricsSourceProperty := &MetricsSourceProperty{
//   	ContentDigest: jsii.String("contentDigest"),
//   	ContentType: jsii.String("contentType"),
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-metricssource.html
//
type CfnModelPackagePropsMixin_MetricsSourceProperty struct {
	// The hash key used for the metrics source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-metricssource.html#cfn-sagemaker-modelpackage-metricssource-contentdigest
	//
	ContentDigest *string `field:"optional" json:"contentDigest" yaml:"contentDigest"`
	// The metric source content type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-metricssource.html#cfn-sagemaker-modelpackage-metricssource-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The S3 URI for the metrics source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-metricssource.html#cfn-sagemaker-modelpackage-metricssource-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

