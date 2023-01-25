package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awseks/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents a helm chart within the Kubernetes system.
//
// Applies/deletes the resources using `kubectl` in sync with the resource.
//
// Example:
//   var cluster cluster
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewHelmChart(this, jsii.String("NginxIngress"), &helmChartProps{
//   	cluster: cluster,
//   	chart: jsii.String("nginx-ingress"),
//   	repository: jsii.String("https://helm.nginx.com/stable"),
//   	namespace: jsii.String("kube-system"),
//   })
//
//   // or, option2: use `addHelmChart`
//   cluster.addHelmChart(jsii.String("NginxIngress"), &helmChartOptions{
//   	chart: jsii.String("nginx-ingress"),
//   	repository: jsii.String("https://helm.nginx.com/stable"),
//   	namespace: jsii.String("kube-system"),
//   })
//
// Experimental.
type HelmChart interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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

// The jsii proxy struct for HelmChart
type jsiiProxy_HelmChart struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_HelmChart) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewHelmChart(scope constructs.Construct, id *string, props *HelmChartProps) HelmChart {
	_init_.Initialize()

	if err := validateNewHelmChartParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HelmChart{}

	_jsii_.Create(
		"monocdk.aws_eks.HelmChart",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHelmChart_Override(h HelmChart, scope constructs.Construct, id *string, props *HelmChartProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.HelmChart",
		[]interface{}{scope, id, props},
		h,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func HelmChart_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHelmChart_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.HelmChart",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HelmChart_RESOURCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks.HelmChart",
		"RESOURCE_TYPE",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HelmChart) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmChart) OnSynthesize(session constructs.ISynthesisSession) {
	if err := h.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (h *jsiiProxy_HelmChart) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HelmChart) Prepare() {
	_jsii_.InvokeVoid(
		h,
		"prepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmChart) Synthesize(session awscdk.ISynthesisSession) {
	if err := h.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"synthesize",
		[]interface{}{session},
	)
}

func (h *jsiiProxy_HelmChart) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HelmChart) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

