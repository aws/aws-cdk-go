package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataProtectionSettings.
// Experimental.
type IDataProtectionSettingsRef interface {
	constructs.IConstruct
	// A reference to a DataProtectionSettings resource.
	// Experimental.
	DataProtectionSettingsRef() *DataProtectionSettingsReference
}

// The jsii proxy for IDataProtectionSettingsRef
type jsiiProxy_IDataProtectionSettingsRef struct {
	internal.Type__constructsIConstruct
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

