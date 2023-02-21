package awsroute53recoverycontrol


// A cluster endpoint.
//
// Specify an endpoint when you want to set or retrieve a routing control state in the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterEndpointProperty := &clusterEndpointProperty{
//   	endpoint: jsii.String("endpoint"),
//   	region: jsii.String("region"),
//   }
//
type CfnCluster_ClusterEndpointProperty struct {
	// A cluster endpoint.
	//
	// Specify an endpoint and AWS Region when you want to set or retrieve a routing control state in the cluster.
	//
	// To get or update the routing control state, see the Amazon Route 53 Application Recovery Controller Routing Control Actions.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The AWS Region for a cluster endpoint.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

