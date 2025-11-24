package mixinsawsredshift


// Properties for CfnScheduledActionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScheduledActionMixinProps := &CfnScheduledActionMixinProps{
//   	Enable: jsii.Boolean(false),
//   	EndTime: jsii.String("endTime"),
//   	IamRole: jsii.String("iamRole"),
//   	Schedule: jsii.String("schedule"),
//   	ScheduledActionDescription: jsii.String("scheduledActionDescription"),
//   	ScheduledActionName: jsii.String("scheduledActionName"),
//   	StartTime: jsii.String("startTime"),
//   	TargetAction: &ScheduledActionTypeProperty{
//   		PauseCluster: &PauseClusterMessageProperty{
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   		ResizeCluster: &ResizeClusterMessageProperty{
//   			Classic: jsii.Boolean(false),
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html
//
type CfnScheduledActionMixinProps struct {
	// If true, the schedule is enabled.
	//
	// If false, the scheduled action does not trigger. For more information about `state` of the scheduled action, see `ScheduledAction` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html#cfn-redshift-scheduledaction-enable
	//
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// The end time in UTC when the schedule is no longer active.
	//
	// After this time, the scheduled action does not trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html#cfn-redshift-scheduledaction-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The IAM role to assume to run the scheduled action.
	//
	// This IAM role must have permission to run the Amazon Redshift API operation in the scheduled action. This IAM role must allow the Amazon Redshift scheduler (Principal scheduler.redshift.amazonaws.com) to assume permissions on your behalf. For more information about the IAM role to use with the Amazon Redshift scheduler, see [Using Identity-Based Policies for Amazon Redshift](https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html) in the *Amazon Redshift Cluster Management Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html#cfn-redshift-scheduledaction-iamrole
	//
	IamRole *string `field:"optional" json:"iamRole" yaml:"iamRole"`
	// The schedule for a one-time (at format) or recurring (cron format) scheduled action.
	//
	// Schedule invocations must be separated by at least one hour.
	//
	// Format of at expressions is " `at(yyyy-mm-ddThh:mm:ss)` ". For example, " `at(2016-03-04T17:27:00)` ".
	//
	// Format of cron expressions is " `cron(Minutes Hours Day-of-month Month Day-of-week Year)` ". For example, " `cron(0 10 ? * MON *)` ". For more information, see [Cron Expressions](https://docs.aws.amazon.com//AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions) in the *Amazon CloudWatch Events User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html#cfn-redshift-scheduledaction-schedule
	//
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// The description of the scheduled action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html#cfn-redshift-scheduledaction-scheduledactiondescription
	//
	ScheduledActionDescription *string `field:"optional" json:"scheduledActionDescription" yaml:"scheduledActionDescription"`
	// The name of the scheduled action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html#cfn-redshift-scheduledaction-scheduledactionname
	//
	ScheduledActionName *string `field:"optional" json:"scheduledActionName" yaml:"scheduledActionName"`
	// The start time in UTC when the schedule is active.
	//
	// Before this time, the scheduled action does not trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html#cfn-redshift-scheduledaction-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// A JSON format string of the Amazon Redshift API operation with input parameters.
	//
	// " `{\"ResizeCluster\":{\"NodeType\":\"ra3.4xlarge\",\"ClusterIdentifier\":\"my-test-cluster\",\"NumberOfNodes\":3}}` ".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html#cfn-redshift-scheduledaction-targetaction
	//
	TargetAction interface{} `field:"optional" json:"targetAction" yaml:"targetAction"`
}

