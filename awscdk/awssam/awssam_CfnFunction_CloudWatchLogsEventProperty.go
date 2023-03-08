package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsEventProperty := &cloudWatchLogsEventProperty{
//   	filterPattern: jsii.String("filterPattern"),
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnFunction_CloudWatchLogsEventProperty struct {
	// `CfnFunction.CloudWatchLogsEventProperty.FilterPattern`.
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// `CfnFunction.CloudWatchLogsEventProperty.LogGroupName`.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

