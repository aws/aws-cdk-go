package interfacesawsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccountAuditConfiguration.
// Experimental.
type IAccountAuditConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AccountAuditConfiguration resource.
	// Experimental.
	AccountAuditConfigurationRef() *AccountAuditConfigurationReference
}

// The jsii proxy for IAccountAuditConfigurationRef
type jsiiProxy_IAccountAuditConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAccountAuditConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAccountAuditConfigurationRef) AccountAuditConfigurationRef() *AccountAuditConfigurationReference {
	var returns *AccountAuditConfigurationReference
	_jsii_.Get(
		j,
		"accountAuditConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccountAuditConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccountAuditConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

