package awsredshift


// Describes a resume cluster operation.
//
// For example, a scheduled action to run the `ResumeCluster` API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resumeClusterMessageProperty := &resumeClusterMessageProperty{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   }
//
type CfnScheduledAction_ResumeClusterMessageProperty struct {
	// The identifier of the cluster to be resumed.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

