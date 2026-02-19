package interfacesawsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataProtectionSettings.
// Experimental.
type IDataProtectionSettingsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataProtectionSettings resource.
	// Experimental.
	DataProtectionSettingsRef() *DataProtectionSettingsReference
}

// The jsii proxy for IDataProtectionSettingsRef
type jsiiProxy_IDataProtectionSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDataProtectionSettingsRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDataProtectionSettingsRef) DataProtectionSettingsRef() *DataProtectionSettingsReference {
	var returns *DataProtectionSettingsReference
	_jsii_.Get(
		j,
		"dataProtectionSettingsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataProtectionSettingsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataProtectionSettingsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

