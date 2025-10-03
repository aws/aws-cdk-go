package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataProtectionSettings.
// Experimental.
type IDataProtectionSettingsRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataProtectionSettingsRef
type jsiiProxy_IDataProtectionSettingsRef struct {
	internal.Type__constructsIConstruct
}

