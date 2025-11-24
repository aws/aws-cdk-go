package mixinsawsiot


// An object that represents the connection attribute, the thing attribute, and the MQTT 5 user property key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   propagatingAttributeProperty := &PropagatingAttributeProperty{
//   	ConnectionAttribute: jsii.String("connectionAttribute"),
//   	ThingAttribute: jsii.String("thingAttribute"),
//   	UserPropertyKey: jsii.String("userPropertyKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-propagatingattribute.html
//
type CfnThingTypePropsMixin_PropagatingAttributeProperty struct {
	// The attribute associated with the connection details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-propagatingattribute.html#cfn-iot-thingtype-propagatingattribute-connectionattribute
	//
	ConnectionAttribute *string `field:"optional" json:"connectionAttribute" yaml:"connectionAttribute"`
	// The thing attribute that is propagating for MQTT 5 message enrichment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-propagatingattribute.html#cfn-iot-thingtype-propagatingattribute-thingattribute
	//
	ThingAttribute *string `field:"optional" json:"thingAttribute" yaml:"thingAttribute"`
	// The key of the MQTT 5 user property, which is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-propagatingattribute.html#cfn-iot-thingtype-propagatingattribute-userpropertykey
	//
	UserPropertyKey *string `field:"optional" json:"userPropertyKey" yaml:"userPropertyKey"`
}

