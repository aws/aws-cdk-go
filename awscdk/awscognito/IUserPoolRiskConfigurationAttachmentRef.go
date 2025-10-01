package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolRiskConfigurationAttachment.
// Experimental.
type IUserPoolRiskConfigurationAttachmentRef interface {
	constructs.IConstruct
	// A reference to a UserPoolRiskConfigurationAttachment resource.
	// Experimental.
	UserPoolRiskConfigurationAttachmentRef() *UserPoolRiskConfigurationAttachmentReference
}

// The jsii proxy for IUserPoolRiskConfigurationAttachmentRef
type jsiiProxy_IUserPoolRiskConfigurationAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserPoolRiskConfigurationAttachmentRef) UserPoolRiskConfigurationAttachmentRef() *UserPoolRiskConfigurationAttachmentReference {
	var returns *UserPoolRiskConfigurationAttachmentReference
	_jsii_.Get(
		j,
		"userPoolRiskConfigurationAttachmentRef",
		&returns,
	)
	return returns
}

