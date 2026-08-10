package interfacesawswellarchitected

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswellarchitected/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReviewTemplate.
// Experimental.
type IReviewTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReviewTemplate resource.
	// Experimental.
	ReviewTemplateRef() *ReviewTemplateReference
}

// The jsii proxy for IReviewTemplateRef
type jsiiProxy_IReviewTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReviewTemplateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IReviewTemplateRef) ReviewTemplateRef() *ReviewTemplateReference {
	var returns *ReviewTemplateReference
	_jsii_.Get(
		j,
		"reviewTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReviewTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReviewTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

