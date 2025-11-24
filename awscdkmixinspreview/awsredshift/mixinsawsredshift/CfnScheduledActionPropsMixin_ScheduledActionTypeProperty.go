package mixinsawsredshift


// The action type that specifies an Amazon Redshift API operation that is supported by the Amazon Redshift scheduler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scheduledActionTypeProperty := &ScheduledActionTypeProperty{
//   	PauseCluster: &PauseClusterMessageProperty{
//   		ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	},
//   	ResizeCluster: &ResizeClusterMessageProperty{
//   		Classic: jsii.Boolean(false),
//   		ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		ClusterType: jsii.String("clusterType"),
//   		NodeType: jsii.String("nodeType"),
//   		NumberOfNodes: jsii.Number(123),
//   	},
//   	ResumeCluster: &ResumeClusterMessageProperty{
//   		ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-scheduledaction-scheduledactiontype.html
//
type CfnScheduledActionPropsMixin_ScheduledActionTypeProperty struct {
	// An action that runs a `PauseCluster` API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-scheduledaction-scheduledactiontype.html#cfn-redshift-scheduledaction-scheduledactiontype-pausecluster
	//
	PauseCluster interface{} `field:"optional" json:"pauseCluster" yaml:"pauseCluster"`
	// An action that runs a `ResizeCluster` API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-scheduledaction-scheduledactiontype.html#cfn-redshift-scheduledaction-scheduledactiontype-resizecluster
	//
	ResizeCluster interface{} `field:"optional" json:"resizeCluster" yaml:"resizeCluster"`
	// An action that runs a `ResumeCluster` API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-scheduledaction-scheduledactiontype.html#cfn-redshift-scheduledaction-scheduledactiontype-resumecluster
	//
	ResumeCluster interface{} `field:"optional" json:"resumeCluster" yaml:"resumeCluster"`
}

