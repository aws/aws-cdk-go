package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolRiskConfigurationAttachment.
// Experimental.
type IUserPoolRiskConfigurationAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPoolRiskConfigurationAttachmentRef
type jsiiProxy_IUserPoolRiskConfigurationAttachmentRef struct {
	internal.Type__constructsIConstruct
}

