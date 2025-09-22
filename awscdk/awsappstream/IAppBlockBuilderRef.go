package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppBlockBuilder.
// Experimental.
type IAppBlockBuilderRef interface {
	constructs.IConstruct
	// A reference to a AppBlockBuilder resource.
	// Experimental.
	AppBlockBuilderRef() *AppBlockBuilderReference
}

// The jsii proxy for IAppBlockBuilderRef
type jsiiProxy_IAppBlockBuilderRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAppBlockBuilderRef) AppBlockBuilderRef() *AppBlockBuilderReference {
	var returns *AppBlockBuilderReference
	_jsii_.Get(
		j,
		"appBlockBuilderRef",
		&returns,
	)
	return returns
}

