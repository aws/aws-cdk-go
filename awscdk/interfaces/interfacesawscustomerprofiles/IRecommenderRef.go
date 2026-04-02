package interfacesawscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Recommender.
// Experimental.
type IRecommenderRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Recommender resource.
	// Experimental.
	RecommenderRef() *RecommenderReference
}

// The jsii proxy for IRecommenderRef
type jsiiProxy_IRecommenderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRecommenderRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRecommenderRef) RecommenderRef() *RecommenderReference {
	var returns *RecommenderReference
	_jsii_.Get(
		j,
		"recommenderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecommenderRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecommenderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

