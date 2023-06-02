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
type CfnThingGroup_ThingGroupPropertiesProperty struct {
	// The thing group attributes in JSON format.
	AttributePayload interface{} `field:"optional" json:"attributePayload" yaml:"attributePayload"`
	// The thing group description.
	ThingGroupDescription *string `field:"optional" json:"thingGroupDescription" yaml:"thingGroupDescription"`
}

