package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserProfile.
// Experimental.
type IUserProfileRef interface {
	constructs.IConstruct
	// A reference to a UserProfile resource.
	// Experimental.
	UserProfileRef() *UserProfileReference
}

// The jsii proxy for IUserProfileRef
type jsiiProxy_IUserProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserProfileRef) UserProfileRef() *UserProfileReference {
	var returns *UserProfileReference
	_jsii_.Get(
		j,
		"userProfileRef",
		&returns,
	)
	return returns
}

