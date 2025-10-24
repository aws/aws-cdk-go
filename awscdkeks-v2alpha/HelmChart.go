package awscdkeks-v2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeks-v2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdkeks-v2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a helm chart within the Kubernetes system.
//
// Applies/deletes the resources using `kubectl` in sync with the resource.
//
// Example:
//   var cluster Cluster
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewHelmChart(this, jsii.String("MyOCIChart"), &HelmChartProps{
//   	Cluster: Cluster,
//   	Chart: jsii.String("some-chart"),
//   	Repository: jsii.String("oci://${ACCOUNT_ID}.dkr.ecr.${ACCOUNT_REGION}.amazonaws.com/${REPO_NAME}"),
//   	Namespace: jsii.String("oci"),
//   	Version: jsii.String("0.0.1"),
//   })
//
// Experimental.
type HelmChart interface {
	constructs.Construct
	// Experimental.
	Atomic() *bool
	// Experimental.
	Chart() *string
	// Experimental.
	ChartAsset() awss3assets.Asset
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Repository() *string
	// Experimental.
	Version() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HelmChart
type jsiiProxy_HelmChart struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_HelmChart) Atomic() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"atomic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmChart) Chart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmChart) ChartAsset() awss3assets.Asset {
	var returns awss3assets.Asset
	_jsii_.Get(
		j,
		"chartAsset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmChart) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmChart) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmChart) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
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
		"@aws-cdk/aws-eks-v2-alpha.HelmChart",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHelmChart_Override(h HelmChart, scope constructs.Construct, id *string, props *HelmChartProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.HelmChart",
		[]interface{}{scope, id, props},
		h,
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
func HelmChart_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHelmChart_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.HelmChart",
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
		"@aws-cdk/aws-eks-v2-alpha.HelmChart",
		"RESOURCE_TYPE",
		&returns,
	)
	return returns
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

