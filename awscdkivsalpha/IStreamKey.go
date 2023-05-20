package awscdkivsalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkivsalpha/v2/internal"
)

// Represents an IVS Stream Key.
// Experimental.
type IStreamKey interface {
	awscdk.IResource
	// The stream-key ARN.
	//
	// For example: arn:aws:ivs:us-west-2:123456789012:stream-key/g1H2I3j4k5L6.
	// Experimental.
	StreamKeyArn() *string
}

// The jsii proxy for IStreamKey
type jsiiProxy_IStreamKey struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IStreamKey) StreamKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamKeyArn",
		&returns,
	)
	return returns
}

