package previewawsfismixins


// Describes the output destinations of the experiment report.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputsProperty := &OutputsProperty{
//   	ExperimentReportS3Configuration: &ExperimentReportS3ConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-outputs.html
//
type CfnExperimentTemplatePropsMixin_OutputsProperty struct {
	// The S3 destination for the experiment report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-outputs.html#cfn-fis-experimenttemplate-outputs-experimentreports3configuration
	//
	ExperimentReportS3Configuration interface{} `field:"optional" json:"experimentReportS3Configuration" yaml:"experimentReportS3Configuration"`
}

