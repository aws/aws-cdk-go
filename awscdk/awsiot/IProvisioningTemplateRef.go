package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProvisioningTemplate.
// Experimental.
type IProvisioningTemplateRef interface {
	constructs.IConstruct
	// A reference to a ProvisioningTemplate resource.
	// Experimental.
	ProvisioningTemplateRef() *ProvisioningTemplateReference
}

// The jsii proxy for IProvisioningTemplateRef
type jsiiProxy_IProvisioningTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProvisioningTemplateRef) ProvisioningTemplateRef() *ProvisioningTemplateReference {
	var returns *ProvisioningTemplateReference
	_jsii_.Get(
		j,
		"provisioningTemplateRef",
		&returns,
	)
	return returns
}

