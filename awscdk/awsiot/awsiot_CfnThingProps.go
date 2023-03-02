package awsiot


// Properties for defining a `CfnThing`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnThingProps := &cfnThingProps{
//   	attributePayload: &attributePayloadProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   	},
//   	thingName: jsii.String("thingName"),
//   }
//
type CfnThingProps struct {
	// A string that contains up to three key value pairs.
	//
	// Maximum length of 800. Duplicates not allowed.
	AttributePayload interface{} `field:"optional" json:"attributePayload" yaml:"attributePayload"`
	// The name of the thing to update.
	//
	// You can't change a thing's name. To change a thing's name, you must create a new thing, give it the new name, and then delete the old thing.
	ThingName *string `field:"optional" json:"thingName" yaml:"thingName"`
}

