package awsproton

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsproton/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentTemplate.
// Experimental.
type IEnvironmentTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEnvironmentTemplateRef
type jsiiProxy_IEnvironmentTemplateRef struct {
	internal.Type__constructsIConstruct
}

