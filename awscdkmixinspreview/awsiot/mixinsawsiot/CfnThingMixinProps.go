package mixinsawsiot


// Properties for CfnThingPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnThingMixinProps := &CfnThingMixinProps{
//   	AttributePayload: &AttributePayloadProperty{
//   		Attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   	},
//   	ThingName: jsii.String("thingName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html
//
type CfnThingMixinProps struct {
	// A string that contains up to three key value pairs.
	//
	// Maximum length of 800. Duplicates not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html#cfn-iot-thing-attributepayload
	//
	AttributePayload interface{} `field:"optional" json:"attributePayload" yaml:"attributePayload"`
	// The name of the thing to update.
	//
	// You can't change a thing's name. To change a thing's name, you must create a new thing, give it the new name, and then delete the old thing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html#cfn-iot-thing-thingname
	//
	ThingName *string `field:"optional" json:"thingName" yaml:"thingName"`
}

