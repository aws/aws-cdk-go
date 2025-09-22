package awsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SourceLocation.
// Experimental.
type ISourceLocationRef interface {
	constructs.IConstruct
	// A reference to a SourceLocation resource.
	// Experimental.
	SourceLocationRef() *SourceLocationReference
}

// The jsii proxy for ISourceLocationRef
type jsiiProxy_ISourceLocationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISourceLocationRef) SourceLocationRef() *SourceLocationReference {
	var returns *SourceLocationReference
	_jsii_.Get(
		j,
		"sourceLocationRef",
		&returns,
	)
	return returns
}

