package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainNameApiAssociation.
// Experimental.
type IDomainNameApiAssociationRef interface {
	constructs.IConstruct
	// A reference to a DomainNameApiAssociation resource.
	// Experimental.
	DomainNameApiAssociationRef() *DomainNameApiAssociationReference
}

// The jsii proxy for IDomainNameApiAssociationRef
type jsiiProxy_IDomainNameApiAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDomainNameApiAssociationRef) DomainNameApiAssociationRef() *DomainNameApiAssociationReference {
	var returns *DomainNameApiAssociationReference
	_jsii_.Get(
		j,
		"domainNameApiAssociationRef",
		&returns,
	)
	return returns
}

