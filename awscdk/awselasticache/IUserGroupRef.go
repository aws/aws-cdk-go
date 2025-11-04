package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserGroup.
// Experimental.
type IUserGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserGroup resource.
	// Experimental.
	UserGroupRef() *UserGroupReference
}

// The jsii proxy for IUserGroupRef
type jsiiProxy_IUserGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IUserGroupRef) UserGroupRef() *UserGroupReference {
	var returns *UserGroupReference
	_jsii_.Get(
		j,
		"userGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

