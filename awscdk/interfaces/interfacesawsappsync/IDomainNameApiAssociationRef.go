package interfacesawsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainNameApiAssociation.
// Experimental.
type IDomainNameApiAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DomainNameApiAssociation resource.
	// Experimental.
	DomainNameApiAssociationRef() *DomainNameApiAssociationReference
}

// The jsii proxy for IDomainNameApiAssociationRef
type jsiiProxy_IDomainNameApiAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDomainNameApiAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IDomainNameApiAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainNameApiAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

