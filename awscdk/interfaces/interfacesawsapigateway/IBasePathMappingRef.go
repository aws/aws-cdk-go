package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BasePathMapping.
// Experimental.
type IBasePathMappingRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BasePathMapping resource.
	// Experimental.
	BasePathMappingRef() *BasePathMappingReference
}

// The jsii proxy for IBasePathMappingRef
type jsiiProxy_IBasePathMappingRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IBasePathMappingRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBasePathMappingRef) BasePathMappingRef() *BasePathMappingReference {
	var returns *BasePathMappingReference
	_jsii_.Get(
		j,
		"basePathMappingRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBasePathMappingRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBasePathMappingRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

