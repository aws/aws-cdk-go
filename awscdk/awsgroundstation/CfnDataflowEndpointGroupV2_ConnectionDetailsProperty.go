package awsgroundstation


// Egress address of AgentEndpoint with an optional mtu.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionDetailsProperty := &ConnectionDetailsProperty{
//   	SocketAddress: &SocketAddressProperty{
//   		Name: jsii.String("name"),
//   		Port: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	Mtu: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-connectiondetails.html
//
type CfnDataflowEndpointGroupV2_ConnectionDetailsProperty struct {
	// A socket address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-connectiondetails.html#cfn-groundstation-dataflowendpointgroupv2-connectiondetails-socketaddress
	//
	SocketAddress interface{} `field:"required" json:"socketAddress" yaml:"socketAddress"`
	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-connectiondetails.html#cfn-groundstation-dataflowendpointgroupv2-connectiondetails-mtu
	//
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
}

