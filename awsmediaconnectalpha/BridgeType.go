package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration of bridge type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   bridgeType := mediaconnect_alpha.BridgeType_Of(jsii.String("value"))
//
// Experimental.
type BridgeType interface {
	// The bridge type string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BridgeType
type jsiiProxy_BridgeType struct {
	_ byte // padding
}

func (j *jsiiProxy_BridgeType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom bridge type value.
// Experimental.
func BridgeType_Of(value *string) BridgeType {
	_init_.Initialize()

	if err := validateBridgeType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns BridgeType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func BridgeType_EGRESS_BRIDGE() BridgeType {
	_init_.Initialize()
	var returns BridgeType
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeType",
		"EGRESS_BRIDGE",
		&returns,
	)
	return returns
}

func BridgeType_INGRESS_BRIDGE() BridgeType {
	_init_.Initialize()
	var returns BridgeType
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeType",
		"INGRESS_BRIDGE",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BridgeType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

