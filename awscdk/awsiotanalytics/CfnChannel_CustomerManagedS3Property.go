package awsiotanalytics


// Used to store channel data in an S3 bucket that you manage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedS3Property := &CustomerManagedS3Property{
//   	Bucket: jsii.String("bucket"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	KeyPrefix: jsii.String("keyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-customermanageds3.html
//
type CfnChannel_CustomerManagedS3Property struct {
	// The name of the S3 bucket in which channel data is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-customermanageds3.html#cfn-iotanalytics-channel-customermanageds3-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-customermanageds3.html#cfn-iotanalytics-channel-customermanageds3-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// (Optional) The prefix used to create the keys of the channel data objects.
	//
	// Each object in an S3 bucket has a key that is its unique identifier within the bucket (each object in a bucket has exactly one key). The prefix must end with a forward slash (/).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-customermanageds3.html#cfn-iotanalytics-channel-customermanageds3-keyprefix
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

