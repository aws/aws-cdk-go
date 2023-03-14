package awsredshift


// Describes a resize cluster operation.
//
// For example, a scheduled action to run the `ResizeCluster` API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resizeClusterMessageProperty := &resizeClusterMessageProperty{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	classic: jsii.Boolean(false),
//   	clusterType: jsii.String("clusterType"),
//   	nodeType: jsii.String("nodeType"),
//   	numberOfNodes: jsii.Number(123),
//   }
//
type CfnScheduledAction_ResizeClusterMessageProperty struct {
	// The unique identifier for the cluster to resize.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// A boolean value indicating whether the resize operation is using the classic resize process.
	//
	// If you don't provide this parameter or set the value to `false` , the resize type is elastic.
	Classic interface{} `field:"optional" json:"classic" yaml:"classic"`
	// The new cluster type for the specified cluster.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// The new node type for the nodes you are adding.
	//
	// If not specified, the cluster's current node type is used.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// The new number of nodes for the cluster.
	//
	// If not specified, the cluster's current number of nodes is used.
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
}

