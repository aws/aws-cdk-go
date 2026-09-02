package interfacesawsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceSetting.
// Experimental.
type IServiceSettingRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServiceSetting resource.
	// Experimental.
	ServiceSettingRef() *ServiceSettingReference
}

// The jsii proxy for IServiceSettingRef
type jsiiProxy_IServiceSettingRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IServiceSettingRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IServiceSettingRef) ServiceSettingRef() *ServiceSettingReference {
	var returns *ServiceSettingReference
	_jsii_.Get(
		j,
		"serviceSettingRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceSettingRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceSettingRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

