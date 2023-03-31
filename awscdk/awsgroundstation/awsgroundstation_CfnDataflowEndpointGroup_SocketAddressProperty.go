package awsgroundstation


// The address of the endpoint, such as `192.168.1.1` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   socketAddressProperty := &socketAddressProperty{
//   	name: jsii.String("name"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataflowEndpointGroup_SocketAddressProperty struct {
	// The name of the endpoint, such as `Endpoint 1` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port of the endpoint, such as `55888` .
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

