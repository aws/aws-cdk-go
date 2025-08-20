package awsivschat


// The S3DestinationConfiguration property type specifies an S3 location where chat logs will be stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationConfigurationProperty := &S3DestinationConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-loggingconfiguration-s3destinationconfiguration.html
//
type CfnLoggingConfiguration_S3DestinationConfigurationProperty struct {
	// Name of the Amazon S3 bucket where chat activity will be logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-loggingconfiguration-s3destinationconfiguration.html#cfn-ivschat-loggingconfiguration-s3destinationconfiguration-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

