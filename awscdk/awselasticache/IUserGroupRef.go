package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserGroup.
// Experimental.
type IUserGroupRef interface {
	constructs.IConstruct
	// A reference to a UserGroup resource.
	// Experimental.
	UserGroupRef() *UserGroupReference
}

// The jsii proxy for IUserGroupRef
type jsiiProxy_IUserGroupRef struct {
	internal.Type__constructsIConstruct
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

