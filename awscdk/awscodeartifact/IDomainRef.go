package awscodeartifact

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeartifact/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Domain.
// Experimental.
type IDomainRef interface {
	constructs.IConstruct
	// A reference to a Domain resource.
	// Experimental.
	DomainRef() *DomainReference
}

// The jsii proxy for IDomainRef
type jsiiProxy_IDomainRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDomainRef) DomainRef() *DomainReference {
	var returns *DomainReference
	_jsii_.Get(
		j,
		"domainRef",
		&returns,
	)
	return returns
}

