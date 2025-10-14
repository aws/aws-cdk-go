package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserAccessLoggingSettings.
// Experimental.
type IUserAccessLoggingSettingsRef interface {
	constructs.IConstruct
	// A reference to a UserAccessLoggingSettings resource.
	// Experimental.
	UserAccessLoggingSettingsRef() *UserAccessLoggingSettingsReference
}

// The jsii proxy for IUserAccessLoggingSettingsRef
type jsiiProxy_IUserAccessLoggingSettingsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserAccessLoggingSettingsRef) UserAccessLoggingSettingsRef() *UserAccessLoggingSettingsReference {
	var returns *UserAccessLoggingSettingsReference
	_jsii_.Get(
		j,
		"userAccessLoggingSettingsRef",
		&returns,
	)
	return returns
}

