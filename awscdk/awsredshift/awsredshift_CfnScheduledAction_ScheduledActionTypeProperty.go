package awsredshift


// The action type that specifies an Amazon Redshift API operation that is supported by the Amazon Redshift scheduler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledActionTypeProperty := &scheduledActionTypeProperty{
//   	pauseCluster: &pauseClusterMessageProperty{
//   		clusterIdentifier: jsii.String("clusterIdentifier"),
//   	},
//   	resizeCluster: &resizeClusterMessageProperty{
//   		clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   		// the properties below are optional
//   		classic: jsii.Boolean(false),
//   		clusterType: jsii.String("clusterType"),
//   		nodeType: jsii.String("nodeType"),
//   		numberOfNodes: jsii.Number(123),
//   	},
//   	resumeCluster: &resumeClusterMessageProperty{
//   		clusterIdentifier: jsii.String("clusterIdentifier"),
//   	},
//   }
//
type CfnScheduledAction_ScheduledActionTypeProperty struct {
	// An action that runs a `PauseCluster` API operation.
	PauseCluster interface{} `field:"optional" json:"pauseCluster" yaml:"pauseCluster"`
	// An action that runs a `ResizeCluster` API operation.
	ResizeCluster interface{} `field:"optional" json:"resizeCluster" yaml:"resizeCluster"`
	// An action that runs a `ResumeCluster` API operation.
	ResumeCluster interface{} `field:"optional" json:"resumeCluster" yaml:"resumeCluster"`
}

