package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The input specification for a channel.
//
// Use the static factory methods to select the input type — mirroring the console's
// "Other" / "CDI" / "Elemental Link" choice.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var cdiInputResolution CdiInputResolution
//   var inputCodec InputCodec
//   var inputMaximumBitrate InputMaximumBitrate
//   var inputResolution InputResolution
//
//   inputSpecification := medialive_alpha.InputSpecification_Cdi(&CdiInputSpecificationProps{
//   	CdiResolution: cdiInputResolution,
//   	Codec: inputCodec,
//   	MaximumBitrate: inputMaximumBitrate,
//   	Resolution: inputResolution,
//   })
//
// Experimental.
type InputSpecification interface {
}

// The jsii proxy struct for InputSpecification
type jsiiProxy_InputSpecification struct {
	_ byte // padding
}

// Experimental.
func NewInputSpecification_Override(i InputSpecification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.InputSpecification",
		nil, // no parameters
		i,
	)
}

// CDI (uncompressed) inputs.
//
// Adds the maximum CDI input resolution.
// Experimental.
func InputSpecification_Cdi(props *CdiInputSpecificationProps) InputSpecification {
	_init_.Initialize()

	if err := validateInputSpecification_CdiParameters(props); err != nil {
		panic(err)
	}
	var returns InputSpecification

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputSpecification",
		"cdi",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Elemental Link inputs.
//
// No additional specification is required.
// Experimental.
func InputSpecification_ElementalLink() InputSpecification {
	_init_.Initialize()

	var returns InputSpecification

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputSpecification",
		"elementalLink",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Standard inputs ("Other" in the console) — the most common case.
// Experimental.
func InputSpecification_Standard(props *StandardInputSpecificationProps) InputSpecification {
	_init_.Initialize()

	if err := validateInputSpecification_StandardParameters(props); err != nil {
		panic(err)
	}
	var returns InputSpecification

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputSpecification",
		"standard",
		[]interface{}{props},
		&returns,
	)

	return returns
}

