package awsredshift


// Properties for defining a `CfnScheduledAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledActionProps := &CfnScheduledActionProps{
//   	ScheduledActionName: jsii.String("scheduledActionName"),
//
//   	// the properties below are optional
//   	Enable: jsii.Boolean(false),
//   	EndTime: jsii.String("endTime"),
//   	IamRole: jsii.String("iamRole"),
//   	Schedule: jsii.String("schedule"),
//   	ScheduledActionDescription: jsii.String("scheduledActionDescription"),
//   	StartTime: jsii.String("startTime"),
//   	TargetAction: &ScheduledActionTypeProperty{
//   		PauseCluster: &PauseClusterMessageProperty{
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   		ResizeCluster: &ResizeClusterMessageProperty{
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//
//   			// the properties below are optional
//   			Classic: jsii.Boolean(false),
//   			ClusterType: jsii.String("clusterType"),
//   			NodeType: jsii.String("nodeType"),
//   			NumberOfNodes: jsii.Number(123),
//   		},
//   		ResumeCluster: &ResumeClusterMessageProperty{
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   	},
//   }
//
type CfnScheduledActionProps struct {
	// The name of the scheduled action.
	ScheduledActionName *string `field:"required" json:"scheduledActionName" yaml:"scheduledActionName"`
	// If true, the schedule is enabled.
	//
	// If false, the scheduled action does not trigger. For more information about `state` of the scheduled action, see `ScheduledAction` .
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// The end time in UTC when the schedule is no longer active.
	//
	// After this time, the scheduled action does not trigger.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The IAM role to assume to run the scheduled action.
	//
	// This IAM role must have permission to run the Amazon Redshift API operation in the scheduled action. This IAM role must allow the Amazon Redshift scheduler (Principal scheduler.redshift.amazonaws.com) to assume permissions on your behalf. For more information about the IAM role to use with the Amazon Redshift scheduler, see [Using Identity-Based Policies for Amazon Redshift](https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html) in the *Amazon Redshift Cluster Management Guide* .
	IamRole *string `field:"optional" json:"iamRole" yaml:"iamRole"`
	// The schedule for a one-time (at format) or recurring (cron format) scheduled action.
	//
	// Schedule invocations must be separated by at least one hour.
	//
	// Format of at expressions is " `at(yyyy-mm-ddThh:mm:ss)` ". For example, " `at(2016-03-04T17:27:00)` ".
	//
	// Format of cron expressions is " `cron(Minutes Hours Day-of-month Month Day-of-week Year)` ". For example, " `cron(0 10 ? * MON *)` ". For more information, see [Cron Expressions](https://docs.aws.amazon.com//AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions) in the *Amazon CloudWatch Events User Guide* .
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// The description of the scheduled action.
	ScheduledActionDescription *string `field:"optional" json:"scheduledActionDescription" yaml:"scheduledActionDescription"`
	// The start time in UTC when the schedule is active.
	//
	// Before this time, the scheduled action does not trigger.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// A JSON format string of the Amazon Redshift API operation with input parameters.
	//
	// " `{\"ResizeCluster\":{\"NodeType\":\"ds2.8xlarge\",\"ClusterIdentifier\":\"my-test-cluster\",\"NumberOfNodes\":3}}` ".
	TargetAction interface{} `field:"optional" json:"targetAction" yaml:"targetAction"`
}

