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
//   		Mqtt5Configuration: &Mqtt5ConfigurationProperty{
//   			PropagatingAttributes: []interface{}{
//   				&PropagatingAttributeProperty{
//   					UserPropertyKey: jsii.String("userPropertyKey"),
//
//   					// the properties below are optional
//   					ConnectionAttribute: jsii.String("connectionAttribute"),
//   					ThingAttribute: jsii.String("thingAttribute"),
//   				},
//   			},
//   		},
//   		SearchableAttributes: []*string{
//   			jsii.String("searchableAttributes"),
//   		},
//   		ThingTypeDescription: jsii.String("thingTypeDescription"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingtype.html
//
type CfnThingTypeProps struct {
	// Deprecates a thing type.
	//
	// You can not associate new things with deprecated thing type. You cannot update `ThingTypeProperties` if the thing type is deprecated.
	//
	// Requires permission to access the [DeprecateThingType](https://docs.aws.amazon.com//service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions) action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingtype.html#cfn-iot-thingtype-deprecatethingtype
	//
	DeprecateThingType interface{} `field:"optional" json:"deprecateThingType" yaml:"deprecateThingType"`
	// Metadata which can be used to manage the thing type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingtype.html#cfn-iot-thingtype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the thing type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingtype.html#cfn-iot-thingtype-thingtypename
	//
	ThingTypeName *string `field:"optional" json:"thingTypeName" yaml:"thingTypeName"`
	// The thing type properties for the thing type to create.
	//
	// It contains information about the new thing type including a description, a list of searchable thing attribute names, and a list of propagating attributes. After a thing type is created, you can only update `Mqtt5Configuration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingtype.html#cfn-iot-thingtype-thingtypeproperties
	//
	ThingTypeProperties interface{} `field:"optional" json:"thingTypeProperties" yaml:"thingTypeProperties"`
}

