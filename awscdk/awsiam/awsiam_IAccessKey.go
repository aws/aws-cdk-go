package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam/internal"
)

// Represents an IAM Access Key.
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html
//
// Experimental.
type IAccessKey interface {
	awscdk.IResource
	// The Access Key ID.
	// Experimental.
	AccessKeyId() *string
	// The Secret Access Key.
	// Experimental.
	SecretAccessKey() awscdk.SecretValue
}

// The jsii proxy for IAccessKey
type jsiiProxy_IAccessKey struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAccessKey) AccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessKey) SecretAccessKey() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretAccessKey",
		&returns,
	)
	return returns
}

