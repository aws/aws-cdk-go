package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainUnit.
// Experimental.
type IDomainUnitRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDomainUnitRef
type jsiiProxy_IDomainUnitRef struct {
	internal.Type__constructsIConstruct
}

