package awsredshift


// Describes a pause cluster operation.
//
// For example, a scheduled action to run the `PauseCluster` API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pauseClusterMessageProperty := &pauseClusterMessageProperty{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   }
//
type CfnScheduledAction_PauseClusterMessageProperty struct {
	// The identifier of the cluster to be paused.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

