package previewawsapsmixins


// Filtering criteria that determine which queries are logged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loggingFilterProperty := &LoggingFilterProperty{
//   	QspThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingfilter.html
//
type CfnWorkspacePropsMixin_LoggingFilterProperty struct {
	// The Query Samples Processed (QSP) threshold above which queries will be logged.
	//
	// Queries processing more samples than this threshold will be captured in logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingfilter.html#cfn-aps-workspace-loggingfilter-qspthreshold
	//
	QspThreshold *float64 `field:"optional" json:"qspThreshold" yaml:"qspThreshold"`
}

