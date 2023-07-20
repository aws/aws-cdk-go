package awsiot


// The attribute payload.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributePayloadProperty := &AttributePayloadProperty{
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thinggroup-attributepayload.html
//
type CfnThingGroup_AttributePayloadProperty struct {
	// A JSON string containing up to three key-value pair in JSON format. For example:.
	//
	// `{\"attributes\":{\"string1\":\"string2\"}}`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thinggroup-attributepayload.html#cfn-iot-thinggroup-attributepayload-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
}

