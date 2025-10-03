package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserAccessLoggingSettings.
// Experimental.
type IUserAccessLoggingSettingsRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserAccessLoggingSettingsRef
type jsiiProxy_IUserAccessLoggingSettingsRef struct {
	internal.Type__constructsIConstruct
}

