package awssam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HttpApi.
// Experimental.
type IHttpApiRef interface {
	constructs.IConstruct
	// A reference to a HttpApi resource.
	// Experimental.
	HttpApiRef() *HttpApiReference
}

// The jsii proxy for IHttpApiRef
type jsiiProxy_IHttpApiRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHttpApiRef) HttpApiRef() *HttpApiReference {
	var returns *HttpApiReference
	_jsii_.Get(
		j,
		"httpApiRef",
		&returns,
	)
	return returns
}

