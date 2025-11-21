package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
)

// A resource that contains data that can be encrypted, using a KMS key.
type IEncryptedResource interface {
	awscdk.IResource
	// Gives permissions to a grantable entity to perform actions on the encryption key.
	GrantOnKey(grantee IGrantable, actions ...*string) *GrantOnKeyResult
}

// The jsii proxy for IEncryptedResource
type jsiiProxy_IEncryptedResource struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEncryptedResource) GrantOnKey(grantee IGrantable, actions ...*string) *GrantOnKeyResult {
	if err := i.validateGrantOnKeyParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns *GrantOnKeyResult

	_jsii_.Invoke(
		i,
		"grantOnKey",
		args,
		&returns,
	)

	return returns
}

