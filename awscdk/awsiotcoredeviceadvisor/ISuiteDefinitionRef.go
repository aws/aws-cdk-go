package awsiotcoredeviceadvisor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotcoredeviceadvisor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SuiteDefinition.
// Experimental.
type ISuiteDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISuiteDefinitionRef
type jsiiProxy_ISuiteDefinitionRef struct {
	internal.Type__constructsIConstruct
}

