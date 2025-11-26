package previewawsiotfleetwisemixins


// A list of information about available network interfaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterfacesItemsProperty := &NetworkInterfacesItemsProperty{
//   	CanInterface: &CanInterfaceProperty{
//   		Name: jsii.String("name"),
//   		ProtocolName: jsii.String("protocolName"),
//   		ProtocolVersion: jsii.String("protocolVersion"),
//   	},
//   	InterfaceId: jsii.String("interfaceId"),
//   	ObdInterface: &ObdInterfaceProperty{
//   		DtcRequestIntervalSeconds: jsii.String("dtcRequestIntervalSeconds"),
//   		HasTransmissionEcu: jsii.String("hasTransmissionEcu"),
//   		Name: jsii.String("name"),
//   		ObdStandard: jsii.String("obdStandard"),
//   		PidRequestIntervalSeconds: jsii.String("pidRequestIntervalSeconds"),
//   		RequestMessageId: jsii.String("requestMessageId"),
//   		UseExtendedIds: jsii.String("useExtendedIds"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html
//
type CfnDecoderManifestPropsMixin_NetworkInterfacesItemsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-caninterface
	//
	CanInterface interface{} `field:"optional" json:"canInterface" yaml:"canInterface"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-interfaceid
	//
	InterfaceId *string `field:"optional" json:"interfaceId" yaml:"interfaceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-obdinterface
	//
	ObdInterface interface{} `field:"optional" json:"obdInterface" yaml:"obdInterface"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

