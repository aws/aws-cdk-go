package awsmediaconnect


// The definition of a media stream that is associated with the output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnFlowOutput_DestinationConfigurationProperty struct {
	// The IP address where contents of the media stream will be sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-destinationconfiguration.html#cfn-mediaconnect-flowoutput-destinationconfiguration-destinationip
	//
	DestinationIp *string `field:"required" json:"destinationIp" yaml:"destinationIp"`
	// The port to use when the content of the media stream is distributed to the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-destinationconfiguration.html#cfn-mediaconnect-flowoutput-destinationconfiguration-destinationport
	//
	DestinationPort *float64 `field:"required" json:"destinationPort" yaml:"destinationPort"`
	// The VPC interface that is used for the media stream associated with the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-destinationconfiguration.html#cfn-mediaconnect-flowoutput-destinationconfiguration-interface
	//
	Interface interface{} `field:"required" json:"interface" yaml:"interface"`
}

