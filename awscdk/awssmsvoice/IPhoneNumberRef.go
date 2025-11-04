package awssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PhoneNumber.
// Experimental.
type IPhoneNumberRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PhoneNumber resource.
	// Experimental.
	PhoneNumberRef() *PhoneNumberReference
}

// The jsii proxy for IPhoneNumberRef
type jsiiProxy_IPhoneNumberRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPhoneNumberRef) PhoneNumberRef() *PhoneNumberReference {
	var returns *PhoneNumberReference
	_jsii_.Get(
		j,
		"phoneNumberRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPhoneNumberRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPhoneNumberRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

