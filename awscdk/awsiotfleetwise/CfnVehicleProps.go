package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVehicle`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var onChange interface{}
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
//   	StateTemplates: []interface{}{
//   		&StateTemplateAssociationProperty{
//   			Identifier: jsii.String("identifier"),
//   			StateTemplateUpdateStrategy: &StateTemplateUpdateStrategyProperty{
//   				OnChange: onChange,
//   				Periodic: &PeriodicStateTemplateUpdateStrategyProperty{
//   					StateTemplateUpdateRate: &TimePeriodProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-vehicle.html
//
type CfnVehicleProps struct {
	// The Amazon Resource Name (ARN) of a decoder manifest associated with the vehicle to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-vehicle.html#cfn-iotfleetwise-vehicle-decodermanifestarn
	//
	DecoderManifestArn *string `field:"required" json:"decoderManifestArn" yaml:"decoderManifestArn"`
	// The Amazon Resource Name (ARN) of the vehicle model (model manifest) to create the vehicle from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-vehicle.html#cfn-iotfleetwise-vehicle-modelmanifestarn
	//
	ModelManifestArn *string `field:"required" json:"modelManifestArn" yaml:"modelManifestArn"`
	// The unique ID of the vehicle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-vehicle.html#cfn-iotfleetwise-vehicle-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An option to create a new AWS IoT thing when creating a vehicle, or to validate an existing thing as a vehicle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-vehicle.html#cfn-iotfleetwise-vehicle-associationbehavior
	//
	AssociationBehavior *string `field:"optional" json:"associationBehavior" yaml:"associationBehavior"`
	// Static information about a vehicle in a key-value pair.
	//
	// For example: `"engine Type"` : `"v6"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-vehicle.html#cfn-iotfleetwise-vehicle-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// Associate state templates to track the state of the vehicle.
	//
	// State templates determine which signal updates the vehicle sends to the cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-vehicle.html#cfn-iotfleetwise-vehicle-statetemplates
	//
	StateTemplates interface{} `field:"optional" json:"stateTemplates" yaml:"stateTemplates"`
	// Metadata which can be used to manage the vehicle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-vehicle.html#cfn-iotfleetwise-vehicle-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

