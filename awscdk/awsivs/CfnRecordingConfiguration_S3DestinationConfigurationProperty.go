package awsivs


// The S3DestinationConfiguration property type describes an S3 location where recorded videos will be stored.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-s3destinationconfiguration.html
//
type CfnRecordingConfiguration_S3DestinationConfigurationProperty struct {
	// Location (S3 bucket name) where recorded videos will be stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-s3destinationconfiguration.html#cfn-ivs-recordingconfiguration-s3destinationconfiguration-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

