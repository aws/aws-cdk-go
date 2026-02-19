package interfacesawsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserSettings.
// Experimental.
type IUserSettingsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UserSettings resource.
	// Experimental.
	UserSettingsRef() *UserSettingsReference
}

// The jsii proxy for IUserSettingsRef
type jsiiProxy_IUserSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IUserSettingsRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IUserSettingsRef) UserSettingsRef() *UserSettingsReference {
	var returns *UserSettingsReference
	_jsii_.Get(
		j,
		"userSettingsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserSettingsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserSettingsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

