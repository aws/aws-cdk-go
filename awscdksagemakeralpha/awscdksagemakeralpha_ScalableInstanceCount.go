// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A scalable sagemaker endpoint attribute.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var model model
//
//
//   variantName := "my-variant"
//   endpointConfig := sagemaker.NewEndpointConfig(this, jsii.String("EndpointConfig"), &endpointConfigProps{
//   	instanceProductionVariants: []instanceProductionVariantProps{
//   		&instanceProductionVariantProps{
//   			model: model,
//   			variantName: variantName,
//   		},
//   	},
//   })
//
//   endpoint := sagemaker.NewEndpoint(this, jsii.String("Endpoint"), &endpointProps{
//   	endpointConfig: endpointConfig,
//   })
//   productionVariant := endpoint.findInstanceProductionVariant(variantName)
//   instanceCount := productionVariant.autoScaleInstanceCount(&enableScalingProps{
//   	maxCapacity: jsii.Number(3),
//   })
//   instanceCount.scaleOnInvocations(jsii.String("LimitRPS"), &invocationsScalingProps{
//   	maxRequestsPerSecond: jsii.Number(30),
//   })
//
// Experimental.
type ScalableInstanceCount interface {
	awsapplicationautoscaling.BaseScalableAttribute
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Scales in or out to achieve a target requests per second per instance.
	// Experimental.
	ScaleOnInvocations(id *string, props *InvocationsScalingProps)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ScalableInstanceCount
type jsiiProxy_ScalableInstanceCount struct {
	internal.Type__awsapplicationautoscalingBaseScalableAttribute
}

func (j *jsiiProxy_ScalableInstanceCount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableInstanceCount) Props() *awsapplicationautoscaling.BaseScalableAttributeProps {
	var returns *awsapplicationautoscaling.BaseScalableAttributeProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ScalableInstanceCount class.
// Experimental.
func NewScalableInstanceCount(scope constructs.Construct, id *string, props *ScalableInstanceCountProps) ScalableInstanceCount {
	_init_.Initialize()

	if err := validateNewScalableInstanceCountParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScalableInstanceCount{}

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.ScalableInstanceCount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ScalableInstanceCount class.
// Experimental.
func NewScalableInstanceCount_Override(s ScalableInstanceCount, scope constructs.Construct, id *string, props *ScalableInstanceCountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.ScalableInstanceCount",
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
// Experimental.
func ScalableInstanceCount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScalableInstanceCount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.ScalableInstanceCount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableInstanceCount) DoScaleOnMetric(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) {
	if err := s.validateDoScaleOnMetricParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"doScaleOnMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableInstanceCount) DoScaleOnSchedule(id *string, props *awsapplicationautoscaling.ScalingSchedule) {
	if err := s.validateDoScaleOnScheduleParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"doScaleOnSchedule",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableInstanceCount) DoScaleToTrackMetric(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps) {
	if err := s.validateDoScaleToTrackMetricParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"doScaleToTrackMetric",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableInstanceCount) ScaleOnInvocations(id *string, props *InvocationsScalingProps) {
	if err := s.validateScaleOnInvocationsParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"scaleOnInvocations",
		[]interface{}{id, props},
	)
}

func (s *jsiiProxy_ScalableInstanceCount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

