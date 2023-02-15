package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
)

// An email identity.
type IEmailIdentity interface {
	awscdk.IResource
	// The name of the email identity.
	EmailIdentityName() *string
}

// The jsii proxy for IEmailIdentity
type jsiiProxy_IEmailIdentity struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEmailIdentity) EmailIdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailIdentityName",
		&returns,
	)
	return returns
}

