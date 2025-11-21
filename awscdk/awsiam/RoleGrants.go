package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

// Collection of grant methods for a IRoleRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var roleRef IRoleRef
//
//   roleGrants := awscdk.Aws_iam.RoleGrants_FromRole(roleRef)
//
type RoleGrants interface {
	// Grant permissions to the given principal to assume this role.
	AssumeRole(identity IPrincipal) Grant
	// Grant permissions to the given principal to pass this role.
	PassRole(identity IPrincipal) Grant
}

// The jsii proxy struct for RoleGrants
type jsiiProxy_RoleGrants struct {
	_ byte // padding
}

// Creates grants for IRoleRef.
func RoleGrants_FromRole(role interfacesawsiam.IRoleRef) RoleGrants {
	_init_.Initialize()

	if err := validateRoleGrants_FromRoleParameters(role); err != nil {
		panic(err)
	}
	var returns RoleGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.RoleGrants",
		"fromRole",
		[]interface{}{role},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleGrants) AssumeRole(identity IPrincipal) Grant {
	if err := r.validateAssumeRoleParameters(identity); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.Invoke(
		r,
		"assumeRole",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RoleGrants) PassRole(identity IPrincipal) Grant {
	if err := r.validatePassRoleParameters(identity); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.Invoke(
		r,
		"passRole",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

