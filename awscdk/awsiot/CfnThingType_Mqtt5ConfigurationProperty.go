package awsiot


// The configuration to add user-defined properties to enrich MQTT 5 messages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mqtt5ConfigurationProperty := &Mqtt5ConfigurationProperty{
//   	PropagatingAttributes: []interface{}{
//   		&PropagatingAttributeProperty{
//   			UserPropertyKey: jsii.String("userPropertyKey"),
//
//   			// the properties below are optional
//   			ConnectionAttribute: jsii.String("connectionAttribute"),
//   			ThingAttribute: jsii.String("thingAttribute"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-mqtt5configuration.html
//
type CfnThingType_Mqtt5ConfigurationProperty struct {
	// An object that represents the connection attribute, the thing attribute, and the MQTT 5 user property key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thingtype-mqtt5configuration.html#cfn-iot-thingtype-mqtt5configuration-propagatingattributes
	//
	PropagatingAttributes interface{} `field:"optional" json:"propagatingAttributes" yaml:"propagatingAttributes"`
}

