package awsgamelift


// A method for collecting container logs for the fleet.
//
// Amazon GameLift Servers saves all standard output for each container in logs, including game session logs. You can select from the following methods:
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	LogDestination: jsii.String("logDestination"),
//   	LogGroupArn: jsii.String("logGroupArn"),
//   	S3BucketName: jsii.String("s3BucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-logconfiguration.html
//
type CfnContainerFleet_LogConfigurationProperty struct {
	// The type of log collection to use for a fleet.
	//
	// - `CLOUDWATCH` -- (default value) Send logs to an Amazon CloudWatch log group that you define. Each container emits a log stream, which is organized in the log group.
	// - `S3` -- Store logs in an Amazon S3 bucket that you define. This bucket must reside in the fleet's home AWS Region.
	// - `NONE` -- Don't collect container logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-logconfiguration.html#cfn-gamelift-containerfleet-logconfiguration-logdestination
	//
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// If log destination is `CLOUDWATCH` , logs are sent to the specified log group in Amazon CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-logconfiguration.html#cfn-gamelift-containerfleet-logconfiguration-loggrouparn
	//
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// If log destination is `S3` , logs are sent to the specified Amazon S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-logconfiguration.html#cfn-gamelift-containerfleet-logconfiguration-s3bucketname
	//
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
}

