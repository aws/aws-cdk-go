package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Authorization token to access the global public ECR Gallery via Docker CLI.
//
// Example:
//   user := iam.NewUser(this, jsii.String("User"))
//   ecr.publicGalleryAuthorizationToken.grantRead(user)
//
// See: https://docs.aws.amazon.com/AmazonECR/latest/public/public-registries.html#public-registry-auth
//
// Experimental.
type PublicGalleryAuthorizationToken interface {
}

// The jsii proxy struct for PublicGalleryAuthorizationToken
type jsiiProxy_PublicGalleryAuthorizationToken struct {
	_ byte // padding
}

// Grant access to retrieve an authorization token.
// Experimental.
func PublicGalleryAuthorizationToken_GrantRead(grantee awsiam.IGrantable) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_ecr.PublicGalleryAuthorizationToken",
		"grantRead",
		[]interface{}{grantee},
	)
}

