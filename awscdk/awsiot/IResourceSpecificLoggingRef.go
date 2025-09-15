package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceSpecificLogging.
// Experimental.
type IResourceSpecificLoggingRef interface {
	constructs.IConstruct
	// A reference to a ResourceSpecificLogging resource.
	// Experimental.
	ResourceSpecificLoggingRef() *ResourceSpecificLoggingReference
}

// The jsii proxy for IResourceSpecificLoggingRef
type jsiiProxy_IResourceSpecificLoggingRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceSpecificLoggingRef) ResourceSpecificLoggingRef() *ResourceSpecificLoggingReference {
	var returns *ResourceSpecificLoggingReference
	_jsii_.Get(
		j,
		"resourceSpecificLoggingRef",
		&returns,
	)
	return returns
}

