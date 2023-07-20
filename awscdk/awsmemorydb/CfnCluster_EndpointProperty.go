package awsmemorydb


// Represents the information required for client programs to connect to the cluster and its nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &EndpointProperty{
//   	Address: jsii.String("address"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-cluster-endpoint.html
//
type CfnCluster_EndpointProperty struct {
	// The DNS hostname of the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-cluster-endpoint.html#cfn-memorydb-cluster-endpoint-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The port number that the engine is listening on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-cluster-endpoint.html#cfn-memorydb-cluster-endpoint-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

