package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The scalable attribute representing task count.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	DesiredCount: jsii.Number(1),
//   	Cpu: jsii.Number(512),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.Service.AutoScaleTaskCount(&EnableScalingProps{
//   	MinCapacity: jsii.Number(1),
//   	MaxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.ScaleOnCpuUtilization(jsii.String("CpuScaling"), &CpuUtilizationScalingProps{
//   	TargetUtilizationPercent: jsii.Number(50),
//   })
//
//   scalableTarget.ScaleOnMemoryUtilization(jsii.String("MemoryScaling"), &MemoryUtilizationScalingProps{
//   	TargetUtilizationPercent: jsii.Number(50),
//   })
//
type ScalableTaskCount interface {
	awsapplicationautoscaling.BaseScalableAttribute
	// The tree node.
	Node() constructs.Node
	Props() *awsapplicationautoscaling.BaseScalableAttributeProps
	// Scale out or in based on a metric value.
	DoScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps)
	// Scale out or in based on time.
	DoScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule)
	// Scale out or in in order to keep a metric around a target value.
	DoScaleToTrackMetric(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps)
	// Scales in or out to achieve a target CPU utilization.
	ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps)
	// Scales in or out to achieve a target memory utilization.
	ScaleOnMemoryUtilization(id *string, props *MemoryUtilizationScalingProps)
	// Scales in or out based on a specified metric value.
	ScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps)
	// Scales in or out to achieve a target Application Load Balancer request count per target.
	ScaleOnRequestCount(id *string, props *RequestCountScalingProps)
	// Scales in or out based on a specified scheduled time.
	ScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule)
	// Scales in or out to achieve a target on a custom metric.
	ScaleToTrackCustomMetric(id *string, props *TrackCustomMetricProps)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ScalableTaskCount
type jsiiProxy_ScalableTaskCount struct {
	internal.Type__awsapplicationautoscalingBaseScalableAttribute
}

func (j *jsiiProxy_ScalableTaskCount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTaskCount) Props() *awsapplicationautoscaling.BaseScalableAttributeProps {
	var returns *awsapplicationautoscaling.BaseScalableAttributeProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ScalableTaskCount class.
func NewScalableTaskCount(scope constructs.Construct, id *string, props *ScalableTaskCountProps) ScalableTaskCount {
	_init_.Initialize()

	if err := validateNewScalableTaskCountParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScalableTaskCount{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ScalableTaskCount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ScalableTaskCount class.
func NewScalableTaskCount_Override(s ScalableTaskCount, scope constructs.Construct, id *string, props *ScalableTaskCountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ScalableTaskCount",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ScalableTaskCount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScalableTaskCount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ScalableTaskCount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTaskCount) DoScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) {
	if err := s.validateDoScaleOnMetricParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"doScaleOnMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) DoScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule) {
	if err := s.validateDoScaleOnScheduleParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"doScaleOnSchedule",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) DoScaleToTrackMetric(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps) {
	if err := s.validateDoScaleToTrackMetricParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"doScaleToTrackMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) {
	if err := s.validateScaleOnCpuUtilizationParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"scaleOnCpuUtilization",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnMemoryUtilization(id *string, props *MemoryUtilizationScalingProps) {
	if err := s.validateScaleOnMemoryUtilizationParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"scaleOnMemoryUtilization",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) {
	if err := s.validateScaleOnMetricParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"scaleOnMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnRequestCount(id *string, props *RequestCountScalingProps) {
	if err := s.validateScaleOnRequestCountParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"scaleOnRequestCount",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule) {
	if err := s.validateScaleOnScheduleParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"scaleOnSchedule",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleToTrackCustomMetric(id *string, props *TrackCustomMetricProps) {
	if err := s.validateScaleToTrackCustomMetricParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"scaleToTrackCustomMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

