package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A target tracking scaling policy that automatically adjusts the capacity provider's compute resources to maintain a specified target value by tracking the required CloudWatch metric.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"))
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//
//   capacityProvider := lambda.NewCapacityProvider(this, jsii.String("MyCapacityProvider"), &CapacityProviderProps{
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	ScalingOptions: lambda.ScalingOptions_Manual([]TargetTrackingScalingPolicy{
//   		lambda.TargetTrackingScalingPolicy_CpuUtilization(jsii.Number(70)),
//   	}),
//   })
//
type TargetTrackingScalingPolicy interface {
	// The predefined metric type.
	MetricType() *string
	// The predefined metric type for this scaling policy.
	PredefinedMetricType() *string
	// The target value for the specified metric as a percentage.
	//
	// The capacity provider will scale resources to maintain this target value.
	TargetValue() *float64
	// The target value for the metric.
	Value() *float64
}

// The jsii proxy struct for TargetTrackingScalingPolicy
type jsiiProxy_TargetTrackingScalingPolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) MetricType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) PredefinedMetricType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedMetricType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates a target tracking scaling policy for CPU utilization.
func TargetTrackingScalingPolicy_CpuUtilization(targetCpuUtilization *float64) TargetTrackingScalingPolicy {
	_init_.Initialize()

	if err := validateTargetTrackingScalingPolicy_CpuUtilizationParameters(targetCpuUtilization); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.TargetTrackingScalingPolicy",
		"cpuUtilization",
		[]interface{}{targetCpuUtilization},
		&returns,
	)

	return returns
}

