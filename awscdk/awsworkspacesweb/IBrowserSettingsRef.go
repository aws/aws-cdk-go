package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BrowserSettings.
// Experimental.
type IBrowserSettingsRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBrowserSettingsRef
type jsiiProxy_IBrowserSettingsRef struct {
	internal.Type__constructsIConstruct
}

