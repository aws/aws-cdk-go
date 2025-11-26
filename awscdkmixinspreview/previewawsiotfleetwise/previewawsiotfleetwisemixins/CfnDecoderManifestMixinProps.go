package previewawsiotfleetwisemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDecoderManifestPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDecoderManifestMixinProps := &CfnDecoderManifestMixinProps{
//   	DefaultForUnmappedSignals: jsii.String("defaultForUnmappedSignals"),
//   	Description: jsii.String("description"),
//   	ModelManifestArn: jsii.String("modelManifestArn"),
//   	Name: jsii.String("name"),
//   	NetworkInterfaces: []interface{}{
//   		&NetworkInterfacesItemsProperty{
//   			CanInterface: &CanInterfaceProperty{
//   				Name: jsii.String("name"),
//   				ProtocolName: jsii.String("protocolName"),
//   				ProtocolVersion: jsii.String("protocolVersion"),
//   			},
//   			InterfaceId: jsii.String("interfaceId"),
//   			ObdInterface: &ObdInterfaceProperty{
//   				DtcRequestIntervalSeconds: jsii.String("dtcRequestIntervalSeconds"),
//   				HasTransmissionEcu: jsii.String("hasTransmissionEcu"),
//   				Name: jsii.String("name"),
//   				ObdStandard: jsii.String("obdStandard"),
//   				PidRequestIntervalSeconds: jsii.String("pidRequestIntervalSeconds"),
//   				RequestMessageId: jsii.String("requestMessageId"),
//   				UseExtendedIds: jsii.String("useExtendedIds"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SignalDecoders: []interface{}{
//   		&SignalDecodersItemsProperty{
//   			CanSignal: &CanSignalProperty{
//   				Factor: jsii.String("factor"),
//   				IsBigEndian: jsii.String("isBigEndian"),
//   				IsSigned: jsii.String("isSigned"),
//   				Length: jsii.String("length"),
//   				MessageId: jsii.String("messageId"),
//   				Name: jsii.String("name"),
//   				Offset: jsii.String("offset"),
//   				SignalValueType: jsii.String("signalValueType"),
//   				StartBit: jsii.String("startBit"),
//   			},
//   			FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   			InterfaceId: jsii.String("interfaceId"),
//   			ObdSignal: &ObdSignalProperty{
//   				BitMaskLength: jsii.String("bitMaskLength"),
//   				BitRightShift: jsii.String("bitRightShift"),
//   				ByteLength: jsii.String("byteLength"),
//   				IsSigned: jsii.String("isSigned"),
//   				Offset: jsii.String("offset"),
//   				Pid: jsii.String("pid"),
//   				PidResponseLength: jsii.String("pidResponseLength"),
//   				Scaling: jsii.String("scaling"),
//   				ServiceMode: jsii.String("serviceMode"),
//   				SignalValueType: jsii.String("signalValueType"),
//   				StartByte: jsii.String("startByte"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html
//
type CfnDecoderManifestMixinProps struct {
	// Use default decoders for all unmapped signals in the model.
	//
	// You don't need to provide any detailed decoding information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html#cfn-iotfleetwise-decodermanifest-defaultforunmappedsignals
	//
	DefaultForUnmappedSignals *string `field:"optional" json:"defaultForUnmappedSignals" yaml:"defaultForUnmappedSignals"`
	// A brief description of the decoder manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html#cfn-iotfleetwise-decodermanifest-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of a vehicle model (model manifest) associated with the decoder manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html#cfn-iotfleetwise-decodermanifest-modelmanifestarn
	//
	ModelManifestArn *string `field:"optional" json:"modelManifestArn" yaml:"modelManifestArn"`
	// The name of the decoder manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html#cfn-iotfleetwise-decodermanifest-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of information about available network interfaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html#cfn-iotfleetwise-decodermanifest-networkinterfaces
	//
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// A list of information about signal decoders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html#cfn-iotfleetwise-decodermanifest-signaldecoders
	//
	SignalDecoders interface{} `field:"optional" json:"signalDecoders" yaml:"signalDecoders"`
	// The state of the decoder manifest.
	//
	// If the status is `ACTIVE` , the decoder manifest can't be edited. If the status is marked `DRAFT` , you can edit the decoder manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html#cfn-iotfleetwise-decodermanifest-status
	//
	// Default: - "DRAFT".
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Metadata that can be used to manage the decoder manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html#cfn-iotfleetwise-decodermanifest-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

