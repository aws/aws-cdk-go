package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The transport protocol for a multicast source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   multicastProtocol := medialive_alpha.MulticastProtocol_Of(jsii.String("value"))
//
// Experimental.
type MulticastProtocol interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MulticastProtocol
type jsiiProxy_MulticastProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_MulticastProtocol) Value() *string {
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
func MulticastProtocol_Of(value *string) MulticastProtocol {
	_init_.Initialize()

	if err := validateMulticastProtocol_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MulticastProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MulticastProtocol",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MulticastProtocol_RTP() MulticastProtocol {
	_init_.Initialize()
	var returns MulticastProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MulticastProtocol",
		"RTP",
		&returns,
	)
	return returns
}

func MulticastProtocol_UDP() MulticastProtocol {
	_init_.Initialize()
	var returns MulticastProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MulticastProtocol",
		"UDP",
		&returns,
	)
	return returns
}

