package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BasePathMappingV2.
// Experimental.
type IBasePathMappingV2Ref interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BasePathMappingV2 resource.
	// Experimental.
	BasePathMappingV2Ref() *BasePathMappingV2Reference
}

// The jsii proxy for IBasePathMappingV2Ref
type jsiiProxy_IBasePathMappingV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IBasePathMappingV2Ref) BasePathMappingV2Ref() *BasePathMappingV2Reference {
	var returns *BasePathMappingV2Reference
	_jsii_.Get(
		j,
		"basePathMappingV2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBasePathMappingV2Ref) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBasePathMappingV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

