package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceProfile.
// Experimental.
type IServiceProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceProfileRef
type jsiiProxy_IServiceProfileRef struct {
	internal.Type__constructsIConstruct
}

