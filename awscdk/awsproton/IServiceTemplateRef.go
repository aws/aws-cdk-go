package awsproton

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsproton/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceTemplate.
// Experimental.
type IServiceTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceTemplateRef
type jsiiProxy_IServiceTemplateRef struct {
	internal.Type__constructsIConstruct
}

