package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainUnit.
// Experimental.
type IDomainUnitRef interface {
	constructs.IConstruct
	// A reference to a DomainUnit resource.
	// Experimental.
	DomainUnitRef() *DomainUnitReference
}

// The jsii proxy for IDomainUnitRef
type jsiiProxy_IDomainUnitRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDomainUnitRef) DomainUnitRef() *DomainUnitReference {
	var returns *DomainUnitReference
	_jsii_.Get(
		j,
		"domainUnitRef",
		&returns,
	)
	return returns
}

