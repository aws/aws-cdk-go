package awslex


// The Amazon CloudWatch Logs log group where the text and metadata logs are delivered.
//
// The log group must exist before you enable logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogGroupLogDestinationProperty := &CloudWatchLogGroupLogDestinationProperty{
//   	CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	LogPrefix: jsii.String("logPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-cloudwatchloggrouplogdestination.html
//
type CfnBot_CloudWatchLogGroupLogDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the log group where text and metadata logs are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-cloudwatchloggrouplogdestination.html#cfn-lex-bot-cloudwatchloggrouplogdestination-cloudwatchloggrouparn
	//
	CloudWatchLogGroupArn *string `field:"required" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// The prefix of the log stream name within the log group that you specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-cloudwatchloggrouplogdestination.html#cfn-lex-bot-cloudwatchloggrouplogdestination-logprefix
	//
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
}

