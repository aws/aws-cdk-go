package mixinsawsmediaconnect


// The transport parameters that you want to associate with an outbound media stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   destinationConfigurationProperty := &DestinationConfigurationProperty{
//   	DestinationIp: jsii.String("destinationIp"),
//   	DestinationPort: jsii.Number(123),
//   	Interface: &InterfaceProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-destinationconfiguration.html
//
type CfnFlowOutputPropsMixin_DestinationConfigurationProperty struct {
	// The IP address where you want MediaConnect to send contents of the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-destinationconfiguration.html#cfn-mediaconnect-flowoutput-destinationconfiguration-destinationip
	//
	DestinationIp *string `field:"optional" json:"destinationIp" yaml:"destinationIp"`
	// The port that you want MediaConnect to use when it distributes the media stream to the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-destinationconfiguration.html#cfn-mediaconnect-flowoutput-destinationconfiguration-destinationport
	//
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// The VPC interface that you want to use for the media stream associated with the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-destinationconfiguration.html#cfn-mediaconnect-flowoutput-destinationconfiguration-interface
	//
	Interface interface{} `field:"optional" json:"interface" yaml:"interface"`
}

