package interfacesawsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QuotaShare.
// Experimental.
type IQuotaShareRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a QuotaShare resource.
	// Experimental.
	QuotaShareRef() *QuotaShareReference
}

// The jsii proxy for IQuotaShareRef
type jsiiProxy_IQuotaShareRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IQuotaShareRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IQuotaShareRef) QuotaShareRef() *QuotaShareReference {
	var returns *QuotaShareReference
	_jsii_.Get(
		j,
		"quotaShareRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQuotaShareRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQuotaShareRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

