package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How to respond to AFD values in the input stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   respondToAfd := medialive_alpha.RespondToAfd_NONE()
//
// Experimental.
type RespondToAfd interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for RespondToAfd
type jsiiProxy_RespondToAfd struct {
	_ byte // padding
}

func (j *jsiiProxy_RespondToAfd) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func RespondToAfd_Of(value *string) RespondToAfd {
	_init_.Initialize()

	if err := validateRespondToAfd_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RespondToAfd

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.RespondToAfd",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RespondToAfd_NONE() RespondToAfd {
	_init_.Initialize()
	var returns RespondToAfd
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RespondToAfd",
		"NONE",
		&returns,
	)
	return returns
}

func RespondToAfd_PASSTHROUGH() RespondToAfd {
	_init_.Initialize()
	var returns RespondToAfd
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RespondToAfd",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

func RespondToAfd_RESPOND() RespondToAfd {
	_init_.Initialize()
	var returns RespondToAfd
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RespondToAfd",
		"RESPOND",
		&returns,
	)
	return returns
}

