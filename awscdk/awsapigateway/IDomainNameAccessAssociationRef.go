package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainNameAccessAssociation.
// Experimental.
type IDomainNameAccessAssociationRef interface {
	constructs.IConstruct
	// A reference to a DomainNameAccessAssociation resource.
	// Experimental.
	DomainNameAccessAssociationRef() *DomainNameAccessAssociationReference
}

// The jsii proxy for IDomainNameAccessAssociationRef
type jsiiProxy_IDomainNameAccessAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDomainNameAccessAssociationRef) DomainNameAccessAssociationRef() *DomainNameAccessAssociationReference {
	var returns *DomainNameAccessAssociationReference
	_jsii_.Get(
		j,
		"domainNameAccessAssociationRef",
		&returns,
	)
	return returns
}

