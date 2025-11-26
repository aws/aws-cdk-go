package previewawsiotfleetwisemixins


// The maximum storage size for the data partition.
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageMaximumSizeProperty := &StorageMaximumSizeProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-storagemaximumsize.html
//
type CfnCampaignPropsMixin_StorageMaximumSizeProperty struct {
	// The data type of the data to store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-storagemaximumsize.html#cfn-iotfleetwise-campaign-storagemaximumsize-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The maximum amount of time to store data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-storagemaximumsize.html#cfn-iotfleetwise-campaign-storagemaximumsize-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

