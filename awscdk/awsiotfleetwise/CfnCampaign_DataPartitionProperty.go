package awsiotfleetwise


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartition.html#cfn-iotfleetwise-campaign-datapartition-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartition.html#cfn-iotfleetwise-campaign-datapartition-storageoptions
	//
	StorageOptions interface{} `field:"required" json:"storageOptions" yaml:"storageOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartition.html#cfn-iotfleetwise-campaign-datapartition-uploadoptions
	//
	UploadOptions interface{} `field:"optional" json:"uploadOptions" yaml:"uploadOptions"`
}

