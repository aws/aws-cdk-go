package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainNameAccessAssociation.
// Experimental.
type IDomainNameAccessAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DomainNameAccessAssociation resource.
	// Experimental.
	DomainNameAccessAssociationRef() *DomainNameAccessAssociationReference
}

// The jsii proxy for IDomainNameAccessAssociationRef
type jsiiProxy_IDomainNameAccessAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IDomainNameAccessAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainNameAccessAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

