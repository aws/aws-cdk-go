package previewawspipesmixins


// Represents the Amazon CloudWatch Logs logging configuration settings for the pipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudwatchLogsLogDestinationProperty := &CloudwatchLogsLogDestinationProperty{
//   	LogGroupArn: jsii.String("logGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-cloudwatchlogslogdestination.html
//
type CfnPipePropsMixin_CloudwatchLogsLogDestinationProperty struct {
	// The AWS Resource Name (ARN) for the CloudWatch log group to which EventBridge sends the log records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-cloudwatchlogslogdestination.html#cfn-pipes-pipe-cloudwatchlogslogdestination-loggrouparn
	//
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}

