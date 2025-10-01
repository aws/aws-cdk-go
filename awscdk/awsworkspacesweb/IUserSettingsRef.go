package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserSettings.
// Experimental.
type IUserSettingsRef interface {
	constructs.IConstruct
	// A reference to a UserSettings resource.
	// Experimental.
	UserSettingsRef() *UserSettingsReference
}

// The jsii proxy for IUserSettingsRef
type jsiiProxy_IUserSettingsRef struct {
	internal.Type__constructsIConstruct
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

