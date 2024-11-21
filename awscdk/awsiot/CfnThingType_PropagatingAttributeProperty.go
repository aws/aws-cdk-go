package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propagatingAttributeProperty := &PropagatingAttributeProperty{
//   	UserPropertyKey: jsii.String("userPropertyKey"),
//
//   	// the properties below are optional
//   	ConnectionAttribute: jsii.String("connectionAttribute"),
//   	ThingAttribute: jsii.String("thingAttribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-propagatingattribute.html
//
type CfnThingType_PropagatingAttributeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-propagatingattribute.html#cfn-iot-thingtype-propagatingattribute-userpropertykey
	//
	UserPropertyKey *string `field:"required" json:"userPropertyKey" yaml:"userPropertyKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-propagatingattribute.html#cfn-iot-thingtype-propagatingattribute-connectionattribute
	//
	ConnectionAttribute *string `field:"optional" json:"connectionAttribute" yaml:"connectionAttribute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-propagatingattribute.html#cfn-iot-thingtype-propagatingattribute-thingattribute
	//
	ThingAttribute *string `field:"optional" json:"thingAttribute" yaml:"thingAttribute"`
}

