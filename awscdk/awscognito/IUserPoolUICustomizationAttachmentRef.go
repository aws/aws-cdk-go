package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolUICustomizationAttachment.
// Experimental.
type IUserPoolUICustomizationAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPoolUICustomizationAttachmentRef
type jsiiProxy_IUserPoolUICustomizationAttachmentRef struct {
	internal.Type__constructsIConstruct
}

