package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchTemplate.
// Experimental.
type ILaunchTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILaunchTemplateRef
type jsiiProxy_ILaunchTemplateRef struct {
	internal.Type__constructsIConstruct
}

