package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Version.
// Experimental.
type IVersionRef interface {
	constructs.IConstruct
	// A reference to a Version resource.
	// Experimental.
	VersionRef() *VersionReference
}

// The jsii proxy for IVersionRef
type jsiiProxy_IVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVersionRef) VersionRef() *VersionReference {
	var returns *VersionReference
	_jsii_.Get(
		j,
		"versionRef",
		&returns,
	)
	return returns
}

