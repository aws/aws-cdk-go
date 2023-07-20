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
//   resumeClusterMessageProperty := &ResumeClusterMessageProperty{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-scheduledaction-resumeclustermessage.html
//
type CfnScheduledAction_ResumeClusterMessageProperty struct {
	// The identifier of the cluster to be resumed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-scheduledaction-resumeclustermessage.html#cfn-redshift-scheduledaction-resumeclustermessage-clusteridentifier
	//
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

