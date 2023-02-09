package awscognitoidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscognitoidentitypool/internal"
)

// Represents an Identity Pool Role Attachment.
// Experimental.
type IIdentityPoolRoleAttachment interface {
	awscdk.IResource
	// Id of the Attachments Underlying Identity Pool.
	// Experimental.
	IdentityPoolId() *string
}

// The jsii proxy for IIdentityPoolRoleAttachment
type jsiiProxy_IIdentityPoolRoleAttachment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IIdentityPoolRoleAttachment) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

