package awsiot


// Thing group properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thingGroupPropertiesProperty := &ThingGroupPropertiesProperty{
//   	AttributePayload: &AttributePayloadProperty{
//   		Attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   	},
//   	ThingGroupDescription: jsii.String("thingGroupDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thinggroup-thinggroupproperties.html
//
type CfnThingGroup_ThingGroupPropertiesProperty struct {
	// The thing group attributes in JSON format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thinggroup-thinggroupproperties.html#cfn-iot-thinggroup-thinggroupproperties-attributepayload
	//
	AttributePayload interface{} `field:"optional" json:"attributePayload" yaml:"attributePayload"`
	// The thing group description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thinggroup-thinggroupproperties.html#cfn-iot-thinggroup-thinggroupproperties-thinggroupdescription
	//
	ThingGroupDescription *string `field:"optional" json:"thingGroupDescription" yaml:"thingGroupDescription"`
}

