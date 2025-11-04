package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataProtectionSettings.
// Experimental.
type IDataProtectionSettingsRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DataProtectionSettings resource.
	// Experimental.
	DataProtectionSettingsRef() *DataProtectionSettingsReference
}

// The jsii proxy for IDataProtectionSettingsRef
type jsiiProxy_IDataProtectionSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IDataProtectionSettingsRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

