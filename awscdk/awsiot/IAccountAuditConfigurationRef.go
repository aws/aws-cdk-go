package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccountAuditConfiguration.
// Experimental.
type IAccountAuditConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccountAuditConfigurationRef
type jsiiProxy_IAccountAuditConfigurationRef struct {
	internal.Type__constructsIConstruct
}

