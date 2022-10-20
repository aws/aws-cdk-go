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
	// `AWS::IoTFleetWise::DecoderManifest.ModelManifestArn`.
	ModelManifestArn *string `field:"required" json:"modelManifestArn" yaml:"modelManifestArn"`
	// `AWS::IoTFleetWise::DecoderManifest.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::IoTFleetWise::DecoderManifest.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTFleetWise::DecoderManifest.NetworkInterfaces`.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// `AWS::IoTFleetWise::DecoderManifest.SignalDecoders`.
	SignalDecoders interface{} `field:"optional" json:"signalDecoders" yaml:"signalDecoders"`
	// `AWS::IoTFleetWise::DecoderManifest.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `AWS::IoTFleetWise::DecoderManifest.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

