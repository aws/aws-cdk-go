package interfacesawscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainAccessGrantForOrganization.
// Experimental.
type IDomainAccessGrantForOrganizationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DomainAccessGrantForOrganization resource.
	// Experimental.
	DomainAccessGrantForOrganizationRef() *DomainAccessGrantForOrganizationReference
}

// The jsii proxy for IDomainAccessGrantForOrganizationRef
type jsiiProxy_IDomainAccessGrantForOrganizationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDomainAccessGrantForOrganizationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDomainAccessGrantForOrganizationRef) DomainAccessGrantForOrganizationRef() *DomainAccessGrantForOrganizationReference {
	var returns *DomainAccessGrantForOrganizationReference
	_jsii_.Get(
		j,
		"domainAccessGrantForOrganizationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainAccessGrantForOrganizationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainAccessGrantForOrganizationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

