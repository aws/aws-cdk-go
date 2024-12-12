package awsiotfleetwise


// The configuration for signal data storage and upload options.
//
// You can only specify these options when the campaign's spooling mode is `TO_DISK` .
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPartitionProperty := &DataPartitionProperty{
//   	Id: jsii.String("id"),
//   	StorageOptions: &DataPartitionStorageOptionsProperty{
//   		MaximumSize: &StorageMaximumSizeProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		MinimumTimeToLive: &StorageMinimumTimeToLiveProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		StorageLocation: jsii.String("storageLocation"),
//   	},
//
//   	// the properties below are optional
//   	UploadOptions: &DataPartitionUploadOptionsProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		ConditionLanguageVersion: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartition.html
//
type CfnCampaign_DataPartitionProperty struct {
	// The ID of the data partition.
	//
	// The data partition ID must be unique within a campaign. You can establish a data partition as the default partition for a campaign by using `default` as the ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartition.html#cfn-iotfleetwise-campaign-datapartition-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The storage options for a data partition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartition.html#cfn-iotfleetwise-campaign-datapartition-storageoptions
	//
	StorageOptions interface{} `field:"required" json:"storageOptions" yaml:"storageOptions"`
	// The upload options for the data partition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartition.html#cfn-iotfleetwise-campaign-datapartition-uploadoptions
	//
	UploadOptions interface{} `field:"optional" json:"uploadOptions" yaml:"uploadOptions"`
}

