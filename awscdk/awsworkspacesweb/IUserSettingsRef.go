package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserSettings.
// Experimental.
type IUserSettingsRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserSettingsRef
type jsiiProxy_IUserSettingsRef struct {
	internal.Type__constructsIConstruct
}

