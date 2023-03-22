package awsiot


// The AttributePayload property specifies up to three attributes for an AWS IoT as key-value pairs.
//
// AttributePayload is a property of the [AWS::IoT::Thing](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributePayloadProperty := &attributePayloadProperty{
//   	attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   }
//
type CfnThing_AttributePayloadProperty struct {
	// A JSON string containing up to three key-value pair in JSON format. For example:.
	//
	// `{\"attributes\":{\"string1\":\"string2\"}}`.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
}

