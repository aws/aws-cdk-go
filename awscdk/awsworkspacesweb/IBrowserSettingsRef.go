package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BrowserSettings.
// Experimental.
type IBrowserSettingsRef interface {
	constructs.IConstruct
	// A reference to a BrowserSettings resource.
	// Experimental.
	BrowserSettingsRef() *BrowserSettingsReference
}

// The jsii proxy for IBrowserSettingsRef
type jsiiProxy_IBrowserSettingsRef struct {
	internal.Type__constructsIConstruct
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

