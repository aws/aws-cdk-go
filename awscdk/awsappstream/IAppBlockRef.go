package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppBlock.
// Experimental.
type IAppBlockRef interface {
	constructs.IConstruct
	// A reference to a AppBlock resource.
	// Experimental.
	AppBlockRef() *AppBlockReference
}

// The jsii proxy for IAppBlockRef
type jsiiProxy_IAppBlockRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAppBlockRef) AppBlockRef() *AppBlockReference {
	var returns *AppBlockReference
	_jsii_.Get(
		j,
		"appBlockRef",
		&returns,
	)
	return returns
}

