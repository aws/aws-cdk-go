package interfacesawskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FeaturedResultsSet.
// Experimental.
type IFeaturedResultsSetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FeaturedResultsSet resource.
	// Experimental.
	FeaturedResultsSetRef() *FeaturedResultsSetReference
}

// The jsii proxy for IFeaturedResultsSetRef
type jsiiProxy_IFeaturedResultsSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFeaturedResultsSetRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFeaturedResultsSetRef) FeaturedResultsSetRef() *FeaturedResultsSetReference {
	var returns *FeaturedResultsSetReference
	_jsii_.Get(
		j,
		"featuredResultsSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeaturedResultsSetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeaturedResultsSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

