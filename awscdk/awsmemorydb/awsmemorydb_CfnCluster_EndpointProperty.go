package awsmemorydb


// Represents the information required for client programs to connect to the cluster and its nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &endpointProperty{
//   	address: jsii.String("address"),
//   	port: jsii.Number(123),
//   }
//
type CfnCluster_EndpointProperty struct {
	// The DNS hostname of the node.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The port number that the engine is listening on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

