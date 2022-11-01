package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVehicle`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVehicleProps := &cfnVehicleProps{
//   	decoderManifestArn: jsii.String("decoderManifestArn"),
//   	modelManifestArn: jsii.String("modelManifestArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	associationBehavior: jsii.String("associationBehavior"),
//   	attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVehicleProps struct {
	// `AWS::IoTFleetWise::Vehicle.DecoderManifestArn`.
	DecoderManifestArn *string `field:"required" json:"decoderManifestArn" yaml:"decoderManifestArn"`
	// `AWS::IoTFleetWise::Vehicle.ModelManifestArn`.
	ModelManifestArn *string `field:"required" json:"modelManifestArn" yaml:"modelManifestArn"`
	// `AWS::IoTFleetWise::Vehicle.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::IoTFleetWise::Vehicle.AssociationBehavior`.
	AssociationBehavior *string `field:"optional" json:"associationBehavior" yaml:"associationBehavior"`
	// `AWS::IoTFleetWise::Vehicle.Attributes`.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// `AWS::IoTFleetWise::Vehicle.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

