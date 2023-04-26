package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awseks/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation resource which applies/restores a JSON patch into a Kubernetes resource.
//
// Example:
//   var cluster cluster
//
//   eks.NewKubernetesPatch(this, jsii.String("hello-kub-deployment-label"), &KubernetesPatchProps{
//   	Cluster: Cluster,
//   	ResourceName: jsii.String("deployment/hello-kubernetes"),
//   	ApplyPatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(5),
//   		},
//   	},
//   	RestorePatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(3),
//   		},
//   	},
//   })
//
// See: https://kubernetes.io/docs/tasks/run-application/update-api-object-kubectl-patch/
//
// Experimental.
type KubernetesPatch interface {
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

// The jsii proxy struct for KubernetesPatch
type jsiiProxy_KubernetesPatch struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_KubernetesPatch) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubernetesPatch(scope constructs.Construct, id *string, props *KubernetesPatchProps) KubernetesPatch {
	_init_.Initialize()

	if err := validateNewKubernetesPatchParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesPatch{}

	_jsii_.Create(
		"monocdk.aws_eks.KubernetesPatch",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKubernetesPatch_Override(k KubernetesPatch, scope constructs.Construct, id *string, props *KubernetesPatchProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.KubernetesPatch",
		[]interface{}{scope, id, props},
		k,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func KubernetesPatch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesPatch_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.KubernetesPatch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesPatch) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesPatch) OnSynthesize(session constructs.ISynthesisSession) {
	if err := k.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (k *jsiiProxy_KubernetesPatch) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesPatch) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesPatch) Synthesize(session awscdk.ISynthesisSession) {
	if err := k.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

func (k *jsiiProxy_KubernetesPatch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesPatch) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

