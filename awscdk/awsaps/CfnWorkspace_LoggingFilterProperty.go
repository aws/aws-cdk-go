package awsaps


// Filters for logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingFilterProperty := &LoggingFilterProperty{
//   	QspThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingfilter.html
//
type CfnWorkspace_LoggingFilterProperty struct {
	// Query logs with QSP above this limit are vended.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingfilter.html#cfn-aps-workspace-loggingfilter-qspthreshold
	//
	QspThreshold *float64 `field:"required" json:"qspThreshold" yaml:"qspThreshold"`
}

