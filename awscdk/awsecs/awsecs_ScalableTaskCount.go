package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// The scalable attribute representing task count.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
//   	minCapacity: jsii.Number(1),
//   	maxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.scaleOnCpuUtilization(jsii.String("CpuScaling"), &cpuUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
//   scalableTarget.scaleOnMemoryUtilization(jsii.String("MemoryScaling"), &memoryUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
// Experimental.
type ScalableTaskCount interface {
	awsapplicationautoscaling.BaseScalableAttribute
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Experimental.
	Props() *awsapplicationautoscaling.BaseScalableAttributeProps
	// Scale out or in based on a metric value.
	// Experimental.
	DoScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps)
	// Scale out or in based on time.
	// Experimental.
	DoScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule)
	// Scale out or in in order to keep a metric around a target value.
	// Experimental.
	DoScaleToTrackMetric(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Scales in or out to achieve a target CPU utilization.
	// Experimental.
	ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps)
	// Scales in or out to achieve a target memory utilization.
	// Experimental.
	ScaleOnMemoryUtilization(id *string, props *MemoryUtilizationScalingProps)
	// Scales in or out based on a specified metric value.
	// Experimental.
	ScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps)
	// Scales in or out to achieve a target Application Load Balancer request count per target.
	// Experimental.
	ScaleOnRequestCount(id *string, props *RequestCountScalingProps)
	// Scales in or out based on a specified scheduled time.
	// Experimental.
	ScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule)
	// Scales in or out to achieve a target on a custom metric.
	// Experimental.
	ScaleToTrackCustomMetric(id *string, props *TrackCustomMetricProps)
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for ScalableTaskCount
type jsiiProxy_ScalableTaskCount struct {
	internal.Type__awsapplicationautoscalingBaseScalableAttribute
}

func (j *jsiiProxy_ScalableTaskCount) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
// Experimental.
func NewScalableTaskCount(scope constructs.Construct, id *string, props *ScalableTaskCountProps) ScalableTaskCount {
	_init_.Initialize()

	j := jsiiProxy_ScalableTaskCount{}

	_jsii_.Create(
		"monocdk.aws_ecs.ScalableTaskCount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ScalableTaskCount class.
// Experimental.
func NewScalableTaskCount_Override(s ScalableTaskCount, scope constructs.Construct, id *string, props *ScalableTaskCountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.ScalableTaskCount",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ScalableTaskCount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.ScalableTaskCount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTaskCount) DoScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) {
	_jsii_.InvokeVoid(
		s,
		"doScaleOnMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) DoScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule) {
	_jsii_.InvokeVoid(
		s,
		"doScaleOnSchedule",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) DoScaleToTrackMetric(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps) {
	_jsii_.InvokeVoid(
		s,
		"doScaleToTrackMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScalableTaskCount) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_ScalableTaskCount) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTaskCount) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnCpuUtilization",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnMemoryUtilization(id *string, props *MemoryUtilizationScalingProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnMemoryUtilization",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnRequestCount(id *string, props *RequestCountScalingProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnRequestCount",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnSchedule",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) ScaleToTrackCustomMetric(id *string, props *TrackCustomMetricProps) {
	_jsii_.InvokeVoid(
		s,
		"scaleToTrackCustomMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableTaskCount) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
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

func (s *jsiiProxy_ScalableTaskCount) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

