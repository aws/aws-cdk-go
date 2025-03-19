package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Authorization token to access the global public ECR Gallery via Docker CLI.
//
// Example:
//   user := iam.NewUser(this, jsii.String("User"))
//   ecr.PublicGalleryAuthorizationToken_GrantRead(user)
//
// See: https://docs.aws.amazon.com/AmazonECR/latest/public/public-registries.html#public-registry-auth
//
type PublicGalleryAuthorizationToken interface {
}

// The jsii proxy struct for PublicGalleryAuthorizationToken
type jsiiProxy_PublicGalleryAuthorizationToken struct {
	_ byte // padding
}

// Grant access to retrieve an authorization token.
func PublicGalleryAuthorizationToken_GrantRead(grantee awsiam.IGrantable) {
	_init_.Initialize()

	if err := validatePublicGalleryAuthorizationToken_GrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_ecr.PublicGalleryAuthorizationToken",
		"grantRead",
		[]interface{}{grantee},
	)
}

