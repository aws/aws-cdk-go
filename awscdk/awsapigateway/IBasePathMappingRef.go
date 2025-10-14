package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BasePathMapping.
// Experimental.
type IBasePathMappingRef interface {
	constructs.IConstruct
	// A reference to a BasePathMapping resource.
	// Experimental.
	BasePathMappingRef() *BasePathMappingReference
}

// The jsii proxy for IBasePathMappingRef
type jsiiProxy_IBasePathMappingRef struct {
	internal.Type__constructsIConstruct
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

