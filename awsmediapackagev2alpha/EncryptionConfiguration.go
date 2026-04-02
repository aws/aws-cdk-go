package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Base class for encryption configurations.
//
// Use `CmafEncryption.speke()`, `TsEncryption.speke()`, or `IsmEncryption.speke()` to create instances.
// Experimental.
type EncryptionConfiguration interface {
}

// The jsii proxy struct for EncryptionConfiguration
type jsiiProxy_EncryptionConfiguration struct {
	_ byte // padding
}

// Experimental.
func NewEncryptionConfiguration_Override(e EncryptionConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.EncryptionConfiguration",
		nil, // no parameters
		e,
	)
}

