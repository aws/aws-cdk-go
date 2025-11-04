package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BrowserSettings.
// Experimental.
type IBrowserSettingsRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a BrowserSettings resource.
	// Experimental.
	BrowserSettingsRef() *BrowserSettingsReference
}

// The jsii proxy for IBrowserSettingsRef
type jsiiProxy_IBrowserSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IBrowserSettingsRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

