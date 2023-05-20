package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnThingType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnThingTypeProps := &CfnThingTypeProps{
//   	DeprecateThingType: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThingTypeName: jsii.String("thingTypeName"),
//   	ThingTypeProperties: &ThingTypePropertiesProperty{
//   		SearchableAttributes: []*string{
//   			jsii.String("searchableAttributes"),
//   		},
//   		ThingTypeDescription: jsii.String("thingTypeDescription"),
//   	},
//   }
//
type CfnThingTypeProps struct {
	// Deprecates a thing type. You can not associate new things with deprecated thing type.
	//
	// Requires permission to access the [DeprecateThingType](https://docs.aws.amazon.com//service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions) action.
	DeprecateThingType interface{} `field:"optional" json:"deprecateThingType" yaml:"deprecateThingType"`
	// Metadata which can be used to manage the thing type.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the thing type.
	ThingTypeName *string `field:"optional" json:"thingTypeName" yaml:"thingTypeName"`
	// The thing type properties for the thing type to create.
	//
	// It contains information about the new thing type including a description, and a list of searchable thing attribute names. `ThingTypeProperties` can't be updated after the initial creation of the `ThingType` .
	ThingTypeProperties interface{} `field:"optional" json:"thingTypeProperties" yaml:"thingTypeProperties"`
}

