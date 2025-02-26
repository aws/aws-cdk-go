package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Authorization token to access private ECR repositories in the current environment via Docker CLI.
//
// Example:
//   user := iam.NewUser(this, jsii.String("User"))
//   ecr.AuthorizationToken_GrantRead(user)
//
// See: https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry_auth.html
//
type AuthorizationToken interface {
}

// The jsii proxy struct for AuthorizationToken
type jsiiProxy_AuthorizationToken struct {
	_ byte // padding
}

// Grant access to retrieve an authorization token.
func AuthorizationToken_GrantRead(grantee awsiam.IGrantable) {
	_init_.Initialize()

	if err := validateAuthorizationToken_GrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_ecr.AuthorizationToken",
		"grantRead",
		[]interface{}{grantee},
	)
}

