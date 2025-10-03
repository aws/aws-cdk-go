package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserProfile.
// Experimental.
type IUserProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserProfileRef
type jsiiProxy_IUserProfileRef struct {
	internal.Type__constructsIConstruct
}

