package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// An AutoScalingGroup.
// Experimental.
type IAutoScalingGroup interface {
	awsiam.IGrantable
	awscdk.IResource
	// Send a message to either an SQS queue or SNS topic when instances launch or terminate.
	// Experimental.
	AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook
	// Add command to the startup script of fleet instances.
	//
	// The command must be in the scripting language supported by the fleet's OS (i.e. Linux/Windows).
	// Does nothing for imported ASGs.
	// Experimental.
	AddUserData(commands ...*string)
	// Add a pool of pre-initialized EC2 instances that sits alongside an Auto Scaling group.
	// Experimental.
	AddWarmPool(options *WarmPoolOptions) WarmPool
	// Scale out or in to achieve a target CPU utilization.
	// Experimental.
	ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in to achieve a target network ingress rate.
	// Experimental.
	ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in, in response to a metric.
	// Experimental.
	ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy
	// Scale out or in to achieve a target network egress rate.
	// Experimental.
	ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in based on time.
	// Experimental.
	ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction
	// Scale out or in in order to keep a metric around a target value.
	// Experimental.
	ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy
	// The arn of the AutoScalingGroup.
	// Experimental.
	AutoScalingGroupArn() *string
	// The name of the AutoScalingGroup.
	// Experimental.
	AutoScalingGroupName() *string
	// The operating system family that the instances in this auto-scaling group belong to.
	//
	// Is 'UNKNOWN' for imported ASGs.
	// Experimental.
	OsType() awsec2.OperatingSystemType
}

// The jsii proxy for IAutoScalingGroup
type jsiiProxy_IAutoScalingGroup struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAutoScalingGroup) AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook {
	if err := i.validateAddLifecycleHookParameters(id, props); err != nil {
		panic(err)
	}
	var returns LifecycleHook

	_jsii_.Invoke(
		i,
		"addLifecycleHook",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) AddUserData(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addUserData",
		args,
	)
}

func (i *jsiiProxy_IAutoScalingGroup) AddWarmPool(options *WarmPoolOptions) WarmPool {
	if err := i.validateAddWarmPoolParameters(options); err != nil {
		panic(err)
	}
	var returns WarmPool

	_jsii_.Invoke(
		i,
		"addWarmPool",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy {
	if err := i.validateScaleOnCpuUtilizationParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleOnCpuUtilization",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	if err := i.validateScaleOnIncomingBytesParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleOnIncomingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy {
	if err := i.validateScaleOnMetricParameters(id, props); err != nil {
		panic(err)
	}
	var returns StepScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleOnMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	if err := i.validateScaleOnOutgoingBytesParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleOnOutgoingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction {
	if err := i.validateScaleOnScheduleParameters(id, props); err != nil {
		panic(err)
	}
	var returns ScheduledAction

	_jsii_.Invoke(
		i,
		"scaleOnSchedule",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy {
	if err := i.validateScaleToTrackMetricParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		i,
		"scaleToTrackMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAutoScalingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IAutoScalingGroup) AutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) OsType() awsec2.OperatingSystemType {
	var returns awsec2.OperatingSystemType
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

