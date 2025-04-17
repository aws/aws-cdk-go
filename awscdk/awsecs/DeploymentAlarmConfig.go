package awsecs


// Configuration for deployment alarms.
//
// Example:
//   import cw "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var elbAlarm alarm
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   	DeploymentAlarms: &DeploymentAlarmConfig{
//   		AlarmNames: []*string{
//   			elbAlarm.AlarmName,
//   		},
//   		Behavior: ecs.AlarmBehavior_ROLLBACK_ON_ALARM,
//   	},
//   })
//
//   // Defining a deployment alarm after the service has been created
//   cpuAlarmName := "MyCpuMetricAlarm"
//   cw.NewAlarm(this, jsii.String("CPUAlarm"), &AlarmProps{
//   	AlarmName: cpuAlarmName,
//   	Metric: service.MetricCpuUtilization(),
//   	EvaluationPeriods: jsii.Number(2),
//   	Threshold: jsii.Number(80),
//   })
//   service.EnableDeploymentAlarms([]*string{
//   	cpuAlarmName,
//   }, &DeploymentAlarmOptions{
//   	Behavior: ecs.AlarmBehavior_FAIL_ON_ALARM,
//   })
//
type DeploymentAlarmConfig struct {
	// Default rollback on alarm.
	// Default: AlarmBehavior.ROLLBACK_ON_ALARM
	//
	Behavior AlarmBehavior `field:"optional" json:"behavior" yaml:"behavior"`
	// List of alarm names to monitor during deployments.
	AlarmNames *[]*string `field:"required" json:"alarmNames" yaml:"alarmNames"`
}

