package awscdkdsqlalpha


// Properties that describe an existing Aurora DSQL cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import dsql_alpha "github.com/aws/aws-cdk-go/awscdkdsqlalpha"
//
//   clusterAttributes := &ClusterAttributes{
//   	ClusterEndpoint: jsii.String("clusterEndpoint"),
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	VpcEndpointServiceName: jsii.String("vpcEndpointServiceName"),
//   }
//
// Experimental.
type ClusterAttributes struct {
	// Connection endpoint for the cluster.
	// Experimental.
	ClusterEndpoint *string `field:"required" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// VPC endpoint service name for the cluster.
	// Experimental.
	VpcEndpointServiceName *string `field:"required" json:"vpcEndpointServiceName" yaml:"vpcEndpointServiceName"`
}

