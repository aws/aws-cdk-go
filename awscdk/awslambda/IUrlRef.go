package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Url.
// Experimental.
type IUrlRef interface {
	constructs.IConstruct
	// A reference to a Url resource.
	// Experimental.
	UrlRef() *UrlReference
}

// The jsii proxy for IUrlRef
type jsiiProxy_IUrlRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUrlRef) UrlRef() *UrlReference {
	var returns *UrlReference
	_jsii_.Get(
		j,
		"urlRef",
		&returns,
	)
	return returns
}

