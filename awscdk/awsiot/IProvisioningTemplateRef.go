package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProvisioningTemplate.
// Experimental.
type IProvisioningTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProvisioningTemplateRef
type jsiiProxy_IProvisioningTemplateRef struct {
	internal.Type__constructsIConstruct
}

