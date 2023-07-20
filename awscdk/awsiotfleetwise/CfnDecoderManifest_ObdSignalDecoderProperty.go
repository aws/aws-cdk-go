package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   obdSignalDecoderProperty := &ObdSignalDecoderProperty{
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	InterfaceId: jsii.String("interfaceId"),
//   	ObdSignal: &ObdSignalProperty{
//   		ByteLength: jsii.String("byteLength"),
//   		Offset: jsii.String("offset"),
//   		Pid: jsii.String("pid"),
//   		PidResponseLength: jsii.String("pidResponseLength"),
//   		Scaling: jsii.String("scaling"),
//   		ServiceMode: jsii.String("serviceMode"),
//   		StartByte: jsii.String("startByte"),
//
//   		// the properties below are optional
//   		BitMaskLength: jsii.String("bitMaskLength"),
//   		BitRightShift: jsii.String("bitRightShift"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html
//
type CfnDecoderManifest_ObdSignalDecoderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html#cfn-iotfleetwise-decodermanifest-obdsignaldecoder-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html#cfn-iotfleetwise-decodermanifest-obdsignaldecoder-interfaceid
	//
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html#cfn-iotfleetwise-decodermanifest-obdsignaldecoder-obdsignal
	//
	ObdSignal interface{} `field:"required" json:"obdSignal" yaml:"obdSignal"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html#cfn-iotfleetwise-decodermanifest-obdsignaldecoder-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

