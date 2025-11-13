package interfacesawsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationConfigRule.
// Experimental.
type IOrganizationConfigRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OrganizationConfigRule resource.
	// Experimental.
	OrganizationConfigRuleRef() *OrganizationConfigRuleReference
}

// The jsii proxy for IOrganizationConfigRuleRef
type jsiiProxy_IOrganizationConfigRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IOrganizationConfigRuleRef) OrganizationConfigRuleRef() *OrganizationConfigRuleReference {
	var returns *OrganizationConfigRuleReference
	_jsii_.Get(
		j,
		"organizationConfigRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizationConfigRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizationConfigRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

