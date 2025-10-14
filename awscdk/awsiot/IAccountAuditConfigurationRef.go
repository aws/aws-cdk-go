package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccountAuditConfiguration.
// Experimental.
type IAccountAuditConfigurationRef interface {
	constructs.IConstruct
	// A reference to a AccountAuditConfiguration resource.
	// Experimental.
	AccountAuditConfigurationRef() *AccountAuditConfigurationReference
}

// The jsii proxy for IAccountAuditConfigurationRef
type jsiiProxy_IAccountAuditConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccountAuditConfigurationRef) AccountAuditConfigurationRef() *AccountAuditConfigurationReference {
	var returns *AccountAuditConfigurationReference
	_jsii_.Get(
		j,
		"accountAuditConfigurationRef",
		&returns,
	)
	return returns
}

