package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDecoderManifest`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDecoderManifestProps := &CfnDecoderManifestProps{
//   	ModelManifestArn: jsii.String("modelManifestArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	NetworkInterfaces: []interface{}{
//   		&NetworkInterfacesItemsProperty{
//   			InterfaceId: jsii.String("interfaceId"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			CanInterface: &CanInterfaceProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				ProtocolName: jsii.String("protocolName"),
//   				ProtocolVersion: jsii.String("protocolVersion"),
//   			},
//   			ObdInterface: &ObdInterfaceProperty{
//   				Name: jsii.String("name"),
//   				RequestMessageId: jsii.String("requestMessageId"),
//
//   				// the properties below are optional
//   				DtcRequestIntervalSeconds: jsii.String("dtcRequestIntervalSeconds"),
//   				HasTransmissionEcu: jsii.String("hasTransmissionEcu"),
//   				ObdStandard: jsii.String("obdStandard"),
//   				PidRequestIntervalSeconds: jsii.String("pidRequestIntervalSeconds"),
//   				UseExtendedIds: jsii.String("useExtendedIds"),
//   			},
//   		},
//   	},
//   	SignalDecoders: []interface{}{
//   		&SignalDecodersItemsProperty{
//   			FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   			InterfaceId: jsii.String("interfaceId"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			CanSignal: &CanSignalProperty{
//   				Factor: jsii.String("factor"),
//   				IsBigEndian: jsii.String("isBigEndian"),
//   				IsSigned: jsii.String("isSigned"),
//   				Length: jsii.String("length"),
//   				MessageId: jsii.String("messageId"),
//   				Offset: jsii.String("offset"),
//   				StartBit: jsii.String("startBit"),
//
//   				// the properties below are optional
//   				Name: jsii.String("name"),
//   			},
//   			ObdSignal: &ObdSignalProperty{
//   				ByteLength: jsii.String("byteLength"),
//   				Offset: jsii.String("offset"),
//   				Pid: jsii.String("pid"),
//   				PidResponseLength: jsii.String("pidResponseLength"),
//   				Scaling: jsii.String("scaling"),
//   				ServiceMode: jsii.String("serviceMode"),
//   				StartByte: jsii.String("startByte"),
//
//   				// the properties below are optional
//   				BitMaskLength: jsii.String("bitMaskLength"),
//   				BitRightShift: jsii.String("bitRightShift"),
//   			},
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDecoderManifestProps struct {
	// The Amazon Resource Name (ARN) of a vehicle model (model manifest) associated with the decoder manifest.
	ModelManifestArn *string `field:"required" json:"modelManifestArn" yaml:"modelManifestArn"`
	// The name of the decoder manifest.
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Optional) A brief description of the decoder manifest.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// (Optional) A list of information about available network interfaces.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// (Optional) A list of information about signal decoders.
	SignalDecoders interface{} `field:"optional" json:"signalDecoders" yaml:"signalDecoders"`
	// (Optional) The state of the decoder manifest.
	//
	// If the status is `ACTIVE` , the decoder manifest can't be edited. If the status is marked `DRAFT` , you can edit the decoder manifest.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// (Optional) Metadata that can be used to manage the decoder manifest.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

