package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactFlowModule.
// Experimental.
type IContactFlowModuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContactFlowModuleRef
type jsiiProxy_IContactFlowModuleRef struct {
	internal.Type__constructsIConstruct
}

