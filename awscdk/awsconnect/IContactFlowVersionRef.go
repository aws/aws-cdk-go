package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactFlowVersion.
// Experimental.
type IContactFlowVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContactFlowVersionRef
type jsiiProxy_IContactFlowVersionRef struct {
	internal.Type__constructsIConstruct
}

