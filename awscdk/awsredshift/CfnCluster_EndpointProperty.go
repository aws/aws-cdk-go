package awsredshift


// Describes a connection endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &EndpointProperty{
//   	Address: jsii.String("address"),
//   	Port: jsii.String("port"),
//   }
//
type CfnCluster_EndpointProperty struct {
	// The DNS address of the cluster.
	//
	// This property is read only.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The port that the database engine is listening on.
	//
	// This property is read only.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

