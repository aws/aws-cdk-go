package awsgroundstation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionDetailsProperty := &ConnectionDetailsProperty{
//   	Mtu: jsii.Number(123),
//   	SocketAddress: &SocketAddressProperty{
//   		Name: jsii.String("name"),
//   		Port: jsii.Number(123),
//   	},
//   }
//
type CfnDataflowEndpointGroup_ConnectionDetailsProperty struct {
	// `CfnDataflowEndpointGroup.ConnectionDetailsProperty.Mtu`.
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// `CfnDataflowEndpointGroup.ConnectionDetailsProperty.SocketAddress`.
	SocketAddress interface{} `field:"optional" json:"socketAddress" yaml:"socketAddress"`
}

