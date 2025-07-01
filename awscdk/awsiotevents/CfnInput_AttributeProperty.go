package awsiotevents


// The attributes from the JSON payload that are made available by the input.
//
// Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage` . Each such message contains a JSON payload. Those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeProperty := &AttributeProperty{
//   	JsonPath: jsii.String("jsonPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-input-attribute.html
//
type CfnInput_AttributeProperty struct {
	// An expression that specifies an attribute-value pair in a JSON structure.
	//
	// Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events ( `BatchPutMessage` ). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.
	//
	// Syntax: `<field-name>.<field-name>...`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-input-attribute.html#cfn-iotevents-input-attribute-jsonpath
	//
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
}

