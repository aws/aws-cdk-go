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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-connectiondetails.html
//
type CfnDataflowEndpointGroup_ConnectionDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-connectiondetails.html#cfn-groundstation-dataflowendpointgroup-connectiondetails-mtu
	//
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-connectiondetails.html#cfn-groundstation-dataflowendpointgroup-connectiondetails-socketaddress
	//
	SocketAddress interface{} `field:"optional" json:"socketAddress" yaml:"socketAddress"`
}

