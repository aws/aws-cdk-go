package previewawsiotfleetwisemixins


// Information about a signal decoder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   signalDecodersItemsProperty := &SignalDecodersItemsProperty{
//   	CanSignal: &CanSignalProperty{
//   		Factor: jsii.String("factor"),
//   		IsBigEndian: jsii.String("isBigEndian"),
//   		IsSigned: jsii.String("isSigned"),
//   		Length: jsii.String("length"),
//   		MessageId: jsii.String("messageId"),
//   		Name: jsii.String("name"),
//   		Offset: jsii.String("offset"),
//   		SignalValueType: jsii.String("signalValueType"),
//   		StartBit: jsii.String("startBit"),
//   	},
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	InterfaceId: jsii.String("interfaceId"),
//   	ObdSignal: &ObdSignalProperty{
//   		BitMaskLength: jsii.String("bitMaskLength"),
//   		BitRightShift: jsii.String("bitRightShift"),
//   		ByteLength: jsii.String("byteLength"),
//   		IsSigned: jsii.String("isSigned"),
//   		Offset: jsii.String("offset"),
//   		Pid: jsii.String("pid"),
//   		PidResponseLength: jsii.String("pidResponseLength"),
//   		Scaling: jsii.String("scaling"),
//   		ServiceMode: jsii.String("serviceMode"),
//   		SignalValueType: jsii.String("signalValueType"),
//   		StartByte: jsii.String("startByte"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html
//
type CfnDecoderManifestPropsMixin_SignalDecodersItemsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-cansignal
	//
	CanSignal interface{} `field:"optional" json:"canSignal" yaml:"canSignal"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"optional" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-interfaceid
	//
	InterfaceId *string `field:"optional" json:"interfaceId" yaml:"interfaceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-obdsignal
	//
	ObdSignal interface{} `field:"optional" json:"obdSignal" yaml:"obdSignal"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

