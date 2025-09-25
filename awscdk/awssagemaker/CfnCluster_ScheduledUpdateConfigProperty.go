package awssagemaker


// The configuration object of the schedule that SageMaker follows when updating the AMI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledUpdateConfigProperty := &ScheduledUpdateConfigProperty{
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//
//   	// the properties below are optional
//   	DeploymentConfig: &DeploymentConfigProperty{
//   		AutoRollbackConfiguration: []interface{}{
//   			&AlarmDetailsProperty{
//   				AlarmName: jsii.String("alarmName"),
//   			},
//   		},
//   		RollingUpdatePolicy: &RollingUpdatePolicyProperty{
//   			MaximumBatchSize: &CapacitySizeConfigProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			RollbackMaximumBatchSize: &CapacitySizeConfigProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   		WaitIntervalInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-scheduledupdateconfig.html
//
type CfnCluster_ScheduledUpdateConfigProperty struct {
	// A cron expression that specifies the schedule that SageMaker follows when updating the AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-scheduledupdateconfig.html#cfn-sagemaker-cluster-scheduledupdateconfig-scheduleexpression
	//
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The configuration to use when updating the AMI versions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-scheduledupdateconfig.html#cfn-sagemaker-cluster-scheduledupdateconfig-deploymentconfig
	//
	DeploymentConfig interface{} `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

