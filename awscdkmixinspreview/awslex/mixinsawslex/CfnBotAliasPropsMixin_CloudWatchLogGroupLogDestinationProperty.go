package mixinsawslex


// The Amazon CloudWatch Logs log group where the text and metadata logs are delivered.
//
// The log group must exist before you enable logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchLogGroupLogDestinationProperty := &CloudWatchLogGroupLogDestinationProperty{
//   	CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	LogPrefix: jsii.String("logPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-cloudwatchloggrouplogdestination.html
//
type CfnBotAliasPropsMixin_CloudWatchLogGroupLogDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the log group where text and metadata logs are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-cloudwatchloggrouplogdestination.html#cfn-lex-botalias-cloudwatchloggrouplogdestination-cloudwatchloggrouparn
	//
	CloudWatchLogGroupArn *string `field:"optional" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// The prefix of the log stream name within the log group that you specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-cloudwatchloggrouplogdestination.html#cfn-lex-botalias-cloudwatchloggrouplogdestination-logprefix
	//
	LogPrefix *string `field:"optional" json:"logPrefix" yaml:"logPrefix"`
}

