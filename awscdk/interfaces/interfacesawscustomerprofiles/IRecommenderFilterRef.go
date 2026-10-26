package interfacesawscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecommenderFilter.
// Experimental.
type IRecommenderFilterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RecommenderFilter resource.
	// Experimental.
	RecommenderFilterRef() *RecommenderFilterReference
}

// The jsii proxy for IRecommenderFilterRef
type jsiiProxy_IRecommenderFilterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRecommenderFilterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRecommenderFilterRef) RecommenderFilterRef() *RecommenderFilterReference {
	var returns *RecommenderFilterReference
	_jsii_.Get(
		j,
		"recommenderFilterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecommenderFilterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecommenderFilterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

