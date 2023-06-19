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
//   cfnVehicleProps := &CfnVehicleProps{
//   	DecoderManifestArn: jsii.String("decoderManifestArn"),
//   	ModelManifestArn: jsii.String("modelManifestArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AssociationBehavior: jsii.String("associationBehavior"),
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVehicleProps struct {
	// The Amazon Resource Name (ARN) of a decoder manifest associated with the vehicle to create.
	DecoderManifestArn *string `field:"required" json:"decoderManifestArn" yaml:"decoderManifestArn"`
	// The Amazon Resource Name (ARN) of the vehicle model (model manifest) to create the vehicle from.
	ModelManifestArn *string `field:"required" json:"modelManifestArn" yaml:"modelManifestArn"`
	// The unique ID of the vehicle.
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Optional) An option to create a new AWS IoT thing when creating a vehicle, or to validate an existing thing as a vehicle.
	AssociationBehavior *string `field:"optional" json:"associationBehavior" yaml:"associationBehavior"`
	// (Optional) Static information about a vehicle in a key-value pair.
	//
	// For example: `"engine Type"` : `"v6"`.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// (Optional) Metadata which can be used to manage the vehicle.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

