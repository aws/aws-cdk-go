package awsmediaconnect


// The transport parameters associated with an incoming media stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputConfigurationProperty := &InputConfigurationProperty{
//   	InputPort: jsii.Number(123),
//   	Interface: &InterfaceProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-inputconfiguration.html
//
type CfnFlow_InputConfigurationProperty struct {
	// The port that the flow listens on for an incoming media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-inputconfiguration.html#cfn-mediaconnect-flow-inputconfiguration-inputport
	//
	InputPort *float64 `field:"required" json:"inputPort" yaml:"inputPort"`
	// The VPC interface where the media stream comes in from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-inputconfiguration.html#cfn-mediaconnect-flow-inputconfiguration-interface
	//
	Interface interface{} `field:"required" json:"interface" yaml:"interface"`
}

