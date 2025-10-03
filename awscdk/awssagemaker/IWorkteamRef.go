package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workteam.
// Experimental.
type IWorkteamRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWorkteamRef
type jsiiProxy_IWorkteamRef struct {
	internal.Type__constructsIConstruct
}

