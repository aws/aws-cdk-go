package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FargateProfile.
// Experimental.
type IFargateProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFargateProfileRef
type jsiiProxy_IFargateProfileRef struct {
	internal.Type__constructsIConstruct
}

