package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BasePathMappingV2.
// Experimental.
type IBasePathMappingV2Ref interface {
	constructs.IConstruct
	// A reference to a BasePathMappingV2 resource.
	// Experimental.
	BasePathMappingV2Ref() *BasePathMappingV2Reference
}

// The jsii proxy for IBasePathMappingV2Ref
type jsiiProxy_IBasePathMappingV2Ref struct {
	internal.Type__constructsIConstruct
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

