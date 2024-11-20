package awsgamelift


// A policy the location and provider of logs from the fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	LogDestination: jsii.String("logDestination"),
//   	S3BucketName: jsii.String("s3BucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-logconfiguration.html
//
type CfnContainerFleet_LogConfigurationProperty struct {
	// Configures the service that provides logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-logconfiguration.html#cfn-gamelift-containerfleet-logconfiguration-logdestination
	//
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// The name of the S3 bucket to pull logs from if S3 is the LogDestination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-logconfiguration.html#cfn-gamelift-containerfleet-logconfiguration-s3bucketname
	//
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
}

