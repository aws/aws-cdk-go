package awspaymentcryptography

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspaymentcryptography/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Alias.
// Experimental.
type IAliasRef interface {
	constructs.IConstruct
	// A reference to a Alias resource.
	// Experimental.
	AliasRef() *AliasReference
}

// The jsii proxy for IAliasRef
type jsiiProxy_IAliasRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAliasRef) AliasRef() *AliasReference {
	var returns *AliasReference
	_jsii_.Get(
		j,
		"aliasRef",
		&returns,
	)
	return returns
}

