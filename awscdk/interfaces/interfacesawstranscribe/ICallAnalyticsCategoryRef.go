package interfacesawstranscribe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstranscribe/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CallAnalyticsCategory.
// Experimental.
type ICallAnalyticsCategoryRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CallAnalyticsCategory resource.
	// Experimental.
	CallAnalyticsCategoryRef() *CallAnalyticsCategoryReference
}

// The jsii proxy for ICallAnalyticsCategoryRef
type jsiiProxy_ICallAnalyticsCategoryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICallAnalyticsCategoryRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICallAnalyticsCategoryRef) CallAnalyticsCategoryRef() *CallAnalyticsCategoryReference {
	var returns *CallAnalyticsCategoryReference
	_jsii_.Get(
		j,
		"callAnalyticsCategoryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICallAnalyticsCategoryRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICallAnalyticsCategoryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

