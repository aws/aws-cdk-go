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
//   cloudWatchLogGroupLogDestinationProperty := &cloudWatchLogGroupLogDestinationProperty{
//   	cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	logPrefix: jsii.String("logPrefix"),
//   }
//
type CfnBotAlias_CloudWatchLogGroupLogDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the log group where text and metadata logs are delivered.
	CloudWatchLogGroupArn *string `field:"required" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// The prefix of the log stream name within the log group that you specified.
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
}

