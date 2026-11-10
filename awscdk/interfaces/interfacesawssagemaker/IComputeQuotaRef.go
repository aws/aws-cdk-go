package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputeQuota.
// Experimental.
type IComputeQuotaRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ComputeQuota resource.
	// Experimental.
	ComputeQuotaRef() *ComputeQuotaReference
}

// The jsii proxy for IComputeQuotaRef
type jsiiProxy_IComputeQuotaRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IComputeQuotaRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IComputeQuotaRef) ComputeQuotaRef() *ComputeQuotaReference {
	var returns *ComputeQuotaReference
	_jsii_.Get(
		j,
		"computeQuotaRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeQuotaRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeQuotaRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

