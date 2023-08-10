package awscdkappconfigalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Implemented by allowed extension event destinations.
// Experimental.
type IEventDestination interface {
	// The URI of the extension event destination.
	// Experimental.
	ExtensionUri() *string
	// The IAM policy document to invoke the event destination.
	// Experimental.
	PolicyDocument() awsiam.PolicyDocument
	// The type of the extension event destination.
	// Experimental.
	Type() SourceType
}

// The jsii proxy for IEventDestination
type jsiiProxy_IEventDestination struct {
	_ byte // padding
}

func (j *jsiiProxy_IEventDestination) ExtensionUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventDestination) PolicyDocument() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventDestination) Type() SourceType {
	var returns SourceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

