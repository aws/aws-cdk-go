package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserAccessLoggingSettings.
// Experimental.
type IUserAccessLoggingSettingsRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserAccessLoggingSettings resource.
	// Experimental.
	UserAccessLoggingSettingsRef() *UserAccessLoggingSettingsReference
}

// The jsii proxy for IUserAccessLoggingSettingsRef
type jsiiProxy_IUserAccessLoggingSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IUserAccessLoggingSettingsRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserAccessLoggingSettingsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

