package awsgroundstation


// Information about the socket address.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   socketAddressProperty := &SocketAddressProperty{
//   	Name: jsii.String("name"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-socketaddress.html
//
type CfnDataflowEndpointGroupV2_SocketAddressProperty struct {
	// Name of a socket address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-socketaddress.html#cfn-groundstation-dataflowendpointgroupv2-socketaddress-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port of a socket address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-socketaddress.html#cfn-groundstation-dataflowendpointgroupv2-socketaddress-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

