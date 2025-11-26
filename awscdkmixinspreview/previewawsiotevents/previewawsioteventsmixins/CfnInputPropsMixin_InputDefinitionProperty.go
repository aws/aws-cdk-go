package previewawsioteventsmixins


// The definition of the input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputDefinitionProperty := &InputDefinitionProperty{
//   	Attributes: []interface{}{
//   		&AttributeProperty{
//   			JsonPath: jsii.String("jsonPath"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-input-inputdefinition.html
//
type CfnInputPropsMixin_InputDefinitionProperty struct {
	// The attributes from the JSON payload that are made available by the input.
	//
	// Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage` . Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-input-inputdefinition.html#cfn-iotevents-input-inputdefinition-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
}

