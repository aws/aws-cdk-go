package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Key types used across AWS Elemental MediaConnect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   keyType := mediaconnect_alpha.KeyType_Of(jsii.String("value"))
//
// Experimental.
type KeyType interface {
	// The key type string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyType
type jsiiProxy_KeyType struct {
	_ byte // padding
}

func (j *jsiiProxy_KeyType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom key type value.
// Experimental.
func KeyType_Of(value *string) KeyType {
	_init_.Initialize()

	if err := validateKeyType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns KeyType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.KeyType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func KeyType_SRT_PASSWORD() KeyType {
	_init_.Initialize()
	var returns KeyType
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.KeyType",
		"SRT_PASSWORD",
		&returns,
	)
	return returns
}

func KeyType_STATIC_KEY() KeyType {
	_init_.Initialize()
	var returns KeyType
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.KeyType",
		"STATIC_KEY",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KeyType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

