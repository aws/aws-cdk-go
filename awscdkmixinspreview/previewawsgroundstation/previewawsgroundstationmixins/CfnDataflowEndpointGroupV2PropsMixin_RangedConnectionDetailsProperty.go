package previewawsgroundstationmixins


// Ingress address of AgentEndpoint with a port range and an optional mtu.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rangedConnectionDetailsProperty := &RangedConnectionDetailsProperty{
//   	Mtu: jsii.Number(123),
//   	SocketAddress: &RangedSocketAddressProperty{
//   		Name: jsii.String("name"),
//   		PortRange: &IntegerRangeProperty{
//   			Maximum: jsii.Number(123),
//   			Minimum: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-rangedconnectiondetails.html
//
type CfnDataflowEndpointGroupV2PropsMixin_RangedConnectionDetailsProperty struct {
	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-rangedconnectiondetails.html#cfn-groundstation-dataflowendpointgroupv2-rangedconnectiondetails-mtu
	//
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// A ranged socket address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-rangedconnectiondetails.html#cfn-groundstation-dataflowendpointgroupv2-rangedconnectiondetails-socketaddress
	//
	SocketAddress interface{} `field:"optional" json:"socketAddress" yaml:"socketAddress"`
}

