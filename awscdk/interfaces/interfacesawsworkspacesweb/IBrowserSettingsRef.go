package interfacesawsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BrowserSettings.
// Experimental.
type IBrowserSettingsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BrowserSettings resource.
	// Experimental.
	BrowserSettingsRef() *BrowserSettingsReference
}

// The jsii proxy for IBrowserSettingsRef
type jsiiProxy_IBrowserSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IBrowserSettingsRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBrowserSettingsRef) BrowserSettingsRef() *BrowserSettingsReference {
	var returns *BrowserSettingsReference
	_jsii_.Get(
		j,
		"browserSettingsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserSettingsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserSettingsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

