package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Key provider settings for HLS encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretValue SecretValue
//
//   hlsKeyProviderSettings := medialive_alpha.HlsKeyProviderSettings_StaticKey(&HlsStaticKeyProps{
//   	KeyProviderServerUrl: jsii.String("keyProviderServerUrl"),
//   	StaticKeyValue: secretValue,
//   })
//
// Experimental.
type HlsKeyProviderSettings interface {
}

// The jsii proxy struct for HlsKeyProviderSettings
type jsiiProxy_HlsKeyProviderSettings struct {
	_ byte // padding
}

// Use a static key for HLS encryption.
// Experimental.
func HlsKeyProviderSettings_StaticKey(props *HlsStaticKeyProps) HlsKeyProviderSettings {
	_init_.Initialize()

	if err := validateHlsKeyProviderSettings_StaticKeyParameters(props); err != nil {
		panic(err)
	}
	var returns HlsKeyProviderSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsKeyProviderSettings",
		"staticKey",
		[]interface{}{props},
		&returns,
	)

	return returns
}

