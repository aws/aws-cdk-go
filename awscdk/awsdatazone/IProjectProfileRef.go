package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProjectProfile.
// Experimental.
type IProjectProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProjectProfileRef
type jsiiProxy_IProjectProfileRef struct {
	internal.Type__constructsIConstruct
}

