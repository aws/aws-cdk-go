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
//   pauseClusterMessageProperty := &PauseClusterMessageProperty{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-scheduledaction-pauseclustermessage.html
//
type CfnScheduledAction_PauseClusterMessageProperty struct {
	// The identifier of the cluster to be paused.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-scheduledaction-pauseclustermessage.html#cfn-redshift-scheduledaction-pauseclustermessage-clusteridentifier
	//
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

