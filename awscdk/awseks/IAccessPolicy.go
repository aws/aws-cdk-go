package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an access policy that defines the permissions and scope for a user or role to access an Amazon EKS cluster.
type IAccessPolicy interface {
	// The scope of the access policy, which determines the level of access granted.
	AccessScope() *AccessScope
	// The access policy itself, which defines the specific permissions.
	Policy() *string
}

// The jsii proxy for IAccessPolicy
type jsiiProxy_IAccessPolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_IAccessPolicy) AccessScope() *AccessScope {
	var returns *AccessScope
	_jsii_.Get(
		j,
		"accessScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPolicy) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

