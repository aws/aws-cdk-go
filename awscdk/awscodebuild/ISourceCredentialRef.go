package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SourceCredential.
// Experimental.
type ISourceCredentialRef interface {
	constructs.IConstruct
	// A reference to a SourceCredential resource.
	// Experimental.
	SourceCredentialRef() *SourceCredentialReference
}

// The jsii proxy for ISourceCredentialRef
type jsiiProxy_ISourceCredentialRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISourceCredentialRef) SourceCredentialRef() *SourceCredentialReference {
	var returns *SourceCredentialReference
	_jsii_.Get(
		j,
		"sourceCredentialRef",
		&returns,
	)
	return returns
}

