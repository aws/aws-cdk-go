package awsiotanalytics


// Where channel data is stored.
//
// You may choose one of `serviceManagedS3` , `customerManagedS3` storage. If not specified, the default is `serviceManagedS3` . This can't be changed after creation of the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceManagedS3 interface{}
//
//   channelStorageProperty := &ChannelStorageProperty{
//   	CustomerManagedS3: &CustomerManagedS3Property{
//   		Bucket: jsii.String("bucket"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		KeyPrefix: jsii.String("keyPrefix"),
//   	},
//   	ServiceManagedS3: serviceManagedS3,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-channelstorage.html
//
type CfnChannel_ChannelStorageProperty struct {
	// Used to store channel data in an S3 bucket that you manage.
	//
	// If customer managed storage is selected, the `retentionPeriod` parameter is ignored. You can't change the choice of S3 storage after the data store is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-channelstorage.html#cfn-iotanalytics-channel-channelstorage-customermanageds3
	//
	CustomerManagedS3 interface{} `field:"optional" json:"customerManagedS3" yaml:"customerManagedS3"`
	// Used to store channel data in an S3 bucket managed by AWS IoT Analytics .
	//
	// You can't change the choice of S3 storage after the data store is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-channelstorage.html#cfn-iotanalytics-channel-channelstorage-servicemanageds3
	//
	ServiceManagedS3 interface{} `field:"optional" json:"serviceManagedS3" yaml:"serviceManagedS3"`
}

