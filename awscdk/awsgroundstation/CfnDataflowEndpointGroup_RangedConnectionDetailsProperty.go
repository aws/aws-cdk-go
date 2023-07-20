package awsgroundstation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-rangedconnectiondetails.html
//
type CfnDataflowEndpointGroup_RangedConnectionDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-rangedconnectiondetails.html#cfn-groundstation-dataflowendpointgroup-rangedconnectiondetails-mtu
	//
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-rangedconnectiondetails.html#cfn-groundstation-dataflowendpointgroup-rangedconnectiondetails-socketaddress
	//
	SocketAddress interface{} `field:"optional" json:"socketAddress" yaml:"socketAddress"`
}

