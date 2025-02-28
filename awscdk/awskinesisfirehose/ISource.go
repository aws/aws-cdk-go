package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// An interface for defining a source that can be used in a Kinesis Data Firehose delivery stream.
type ISource interface {
	// Grant read permissions for this source resource and its contents to an IAM principal (the delivery stream).
	//
	// If an encryption key is used, permission to use the key to decrypt the
	// contents of the stream will also be granted.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy for ISource
type jsiiProxy_ISource struct {
	_ byte // padding
}

func (i *jsiiProxy_ISource) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

