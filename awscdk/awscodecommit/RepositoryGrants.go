package awscodecommit

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodecommit"
)

// Collection of grant methods for a IRepositoryRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var repositoryRef IRepositoryRef
//
//   repositoryGrants := awscdk.Aws_codecommit.RepositoryGrants_FromRepository(repositoryRef)
//
type RepositoryGrants interface {
	Resource() interfacesawscodecommit.IRepositoryRef
	// Grants pull permissions.
	Pull(grantee awsiam.IGrantable) awsiam.Grant
	// Grants pullPush permissions.
	PullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Grants read permissions.
	Read(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for RepositoryGrants
type jsiiProxy_RepositoryGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_RepositoryGrants) Resource() interfacesawscodecommit.IRepositoryRef {
	var returns interfacesawscodecommit.IRepositoryRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for RepositoryGrants.
func RepositoryGrants_FromRepository(resource interfacesawscodecommit.IRepositoryRef) RepositoryGrants {
	_init_.Initialize()

	if err := validateRepositoryGrants_FromRepositoryParameters(resource); err != nil {
		panic(err)
	}
	var returns RepositoryGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codecommit.RepositoryGrants",
		"fromRepository",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryGrants) Pull(grantee awsiam.IGrantable) awsiam.Grant {
	if err := r.validatePullParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"pull",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryGrants) PullPush(grantee awsiam.IGrantable) awsiam.Grant {
	if err := r.validatePullPushParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"pullPush",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryGrants) Read(grantee awsiam.IGrantable) awsiam.Grant {
	if err := r.validateReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"read",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

