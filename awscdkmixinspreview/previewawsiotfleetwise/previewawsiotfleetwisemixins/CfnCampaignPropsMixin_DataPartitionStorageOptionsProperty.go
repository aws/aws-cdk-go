package previewawsiotfleetwisemixins


// Size, time, and location options for the data partition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPartitionStorageOptionsProperty := &DataPartitionStorageOptionsProperty{
//   	MaximumSize: &StorageMaximumSizeProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	MinimumTimeToLive: &StorageMinimumTimeToLiveProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	StorageLocation: jsii.String("storageLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartitionstorageoptions.html
//
type CfnCampaignPropsMixin_DataPartitionStorageOptionsProperty struct {
	// The maximum storage size of the data stored in the data partition.
	//
	// > Newer data overwrites older data when the partition reaches the maximum size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartitionstorageoptions.html#cfn-iotfleetwise-campaign-datapartitionstorageoptions-maximumsize
	//
	MaximumSize interface{} `field:"optional" json:"maximumSize" yaml:"maximumSize"`
	// The amount of time that data in this partition will be kept on disk.
	//
	// - After the designated amount of time passes, the data can be removed, but it's not guaranteed to be removed.
	// - Before the time expires, data in this partition can still be deleted if the partition reaches its configured maximum size.
	// - Newer data will overwrite older data when the partition reaches the maximum size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartitionstorageoptions.html#cfn-iotfleetwise-campaign-datapartitionstorageoptions-minimumtimetolive
	//
	MinimumTimeToLive interface{} `field:"optional" json:"minimumTimeToLive" yaml:"minimumTimeToLive"`
	// The folder name for the data partition under the campaign storage folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartitionstorageoptions.html#cfn-iotfleetwise-campaign-datapartitionstorageoptions-storagelocation
	//
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
}

