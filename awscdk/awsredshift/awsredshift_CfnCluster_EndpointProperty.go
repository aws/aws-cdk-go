package awsredshift


// Describes a connection endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &endpointProperty{
//   	address: jsii.String("address"),
//   	port: jsii.String("port"),
//   }
//
type CfnCluster_EndpointProperty struct {
	// The DNS address of the Cluster.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The port that the database engine is listening on.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

