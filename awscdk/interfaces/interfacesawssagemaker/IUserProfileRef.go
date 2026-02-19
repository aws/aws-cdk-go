package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserProfile.
// Experimental.
type IUserProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UserProfile resource.
	// Experimental.
	UserProfileRef() *UserProfileReference
}

// The jsii proxy for IUserProfileRef
type jsiiProxy_IUserProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IUserProfileRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IUserProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

