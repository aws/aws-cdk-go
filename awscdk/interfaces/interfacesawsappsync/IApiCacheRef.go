package interfacesawsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiCache.
// Experimental.
type IApiCacheRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApiCache resource.
	// Experimental.
	ApiCacheRef() *ApiCacheReference
}

// The jsii proxy for IApiCacheRef
type jsiiProxy_IApiCacheRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IApiCacheRef) ApiCacheRef() *ApiCacheReference {
	var returns *ApiCacheReference
	_jsii_.Get(
		j,
		"apiCacheRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiCacheRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiCacheRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

