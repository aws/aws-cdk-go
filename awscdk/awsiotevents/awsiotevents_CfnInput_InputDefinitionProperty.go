package awsiotevents


// The definition of the input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDefinitionProperty := &inputDefinitionProperty{
//   	attributes: []interface{}{
//   		&attributeProperty{
//   			jsonPath: jsii.String("jsonPath"),
//   		},
//   	},
//   }
//
type CfnInput_InputDefinitionProperty struct {
	// The attributes from the JSON payload that are made available by the input.
	//
	// Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage` . Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
}

