package awselasticloadbalancingv2


// Count of HTTP status originating from the targets.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var service fargateService
//   var blueTargetGroup applicationTargetGroup
//   var greenTargetGroup applicationTargetGroup
//   var listener iApplicationListener
//
//
//   // Alarm on the number of unhealthy ECS tasks in each target group
//   blueUnhealthyHosts := cloudwatch.NewAlarm(this, jsii.String("BlueUnhealthyHosts"), &AlarmProps{
//   	AlarmName: jsii.String(awscdk.stack_Of(this).stackName + "-Unhealthy-Hosts-Blue"),
//   	Metric: blueTargetGroup.MetricUnhealthyHostCount(),
//   	Threshold: jsii.Number(1),
//   	EvaluationPeriods: jsii.Number(2),
//   })
//
//   greenUnhealthyHosts := cloudwatch.NewAlarm(this, jsii.String("GreenUnhealthyHosts"), &AlarmProps{
//   	AlarmName: jsii.String(awscdk.stack_Of(this).stackName + "-Unhealthy-Hosts-Green"),
//   	Metric: greenTargetGroup.*MetricUnhealthyHostCount(),
//   	Threshold: jsii.Number(1),
//   	EvaluationPeriods: jsii.Number(2),
//   })
//
//   // Alarm on the number of HTTP 5xx responses returned by each target group
//   blueApiFailure := cloudwatch.NewAlarm(this, jsii.String("Blue5xx"), &AlarmProps{
//   	AlarmName: jsii.String(awscdk.stack_Of(this).stackName + "-Http-5xx-Blue"),
//   	Metric: blueTargetGroup.MetricHttpCodeTarget(elbv2.HttpCodeTarget_TARGET_5XX_COUNT, &MetricOptions{
//   		Period: awscdk.Duration_Minutes(jsii.Number(1)),
//   	}),
//   	Threshold: jsii.Number(1),
//   	EvaluationPeriods: jsii.Number(1),
//   })
//
//   greenApiFailure := cloudwatch.NewAlarm(this, jsii.String("Green5xx"), &AlarmProps{
//   	AlarmName: jsii.String(awscdk.stack_Of(this).stackName + "-Http-5xx-Green"),
//   	Metric: greenTargetGroup.*MetricHttpCodeTarget(elbv2.HttpCodeTarget_TARGET_5XX_COUNT, &MetricOptions{
//   		Period: awscdk.Duration_*Minutes(jsii.Number(1)),
//   	}),
//   	Threshold: jsii.Number(1),
//   	EvaluationPeriods: jsii.Number(1),
//   })
//
//   codedeploy.NewEcsDeploymentGroup(this, jsii.String("BlueGreenDG"), &EcsDeploymentGroupProps{
//   	// CodeDeploy will monitor these alarms during a deployment and automatically roll back
//   	Alarms: []iAlarm{
//   		blueUnhealthyHosts,
//   		greenUnhealthyHosts,
//   		blueApiFailure,
//   		greenApiFailure,
//   	},
//   	AutoRollback: &AutoRollbackConfig{
//   		// CodeDeploy will automatically roll back if a deployment is stopped
//   		StoppedDeployment: jsii.Boolean(true),
//   	},
//   	Service: Service,
//   	BlueGreenDeploymentConfig: &EcsBlueGreenDeploymentConfig{
//   		BlueTargetGroup: *BlueTargetGroup,
//   		GreenTargetGroup: *GreenTargetGroup,
//   		Listener: *Listener,
//   	},
//   	DeploymentConfig: codedeploy.EcsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
//   })
//
type HttpCodeTarget string

const (
	// The number of 2xx response codes from targets.
	HttpCodeTarget_TARGET_2XX_COUNT HttpCodeTarget = "TARGET_2XX_COUNT"
	// The number of 3xx response codes from targets.
	HttpCodeTarget_TARGET_3XX_COUNT HttpCodeTarget = "TARGET_3XX_COUNT"
	// The number of 4xx response codes from targets.
	HttpCodeTarget_TARGET_4XX_COUNT HttpCodeTarget = "TARGET_4XX_COUNT"
	// The number of 5xx response codes from targets.
	HttpCodeTarget_TARGET_5XX_COUNT HttpCodeTarget = "TARGET_5XX_COUNT"
)

