package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDecoderManifest`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkInterfaces interface{}
//   var signalDecoders interface{}
//
//   cfnDecoderManifestProps := &cfnDecoderManifestProps{
//   	modelManifestArn: jsii.String("modelManifestArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	networkInterfaces: []interface{}{
//   		networkInterfaces,
//   	},
//   	signalDecoders: []interface{}{
//   		signalDecoders,
//   	},
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDecoderManifestProps struct {
	// The ARN of a vehicle model (model manifest) associated with the decoder manifest.
	ModelManifestArn *string `field:"required" json:"modelManifestArn" yaml:"modelManifestArn"`
	// The name of the decoder manifest.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A brief description of the decoder manifest.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTFleetWise::DecoderManifest.NetworkInterfaces`.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// `AWS::IoTFleetWise::DecoderManifest.SignalDecoders`.
	SignalDecoders interface{} `field:"optional" json:"signalDecoders" yaml:"signalDecoders"`
	// The state of the decoder manifest.
	//
	// If the status is `ACTIVE` , the decoder manifest can't be edited. If the status is marked `DRAFT` , you can edit the decoder manifest.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `AWS::IoTFleetWise::DecoderManifest.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

