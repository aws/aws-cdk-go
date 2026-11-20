package interfacesawsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DefaultHostManagementConfiguration.
// Experimental.
type IDefaultHostManagementConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DefaultHostManagementConfiguration resource.
	// Experimental.
	DefaultHostManagementConfigurationRef() *DefaultHostManagementConfigurationReference
}

// The jsii proxy for IDefaultHostManagementConfigurationRef
type jsiiProxy_IDefaultHostManagementConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDefaultHostManagementConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDefaultHostManagementConfigurationRef) DefaultHostManagementConfigurationRef() *DefaultHostManagementConfigurationReference {
	var returns *DefaultHostManagementConfigurationReference
	_jsii_.Get(
		j,
		"defaultHostManagementConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDefaultHostManagementConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDefaultHostManagementConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

