package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines the hosted configuration content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//
//   configurationContent := appconfig_alpha.ConfigurationContent_FromFile(jsii.String("inputPath"), jsii.String("contentType"))
//
// Deprecated.
type ConfigurationContent interface {
	// The configuration content.
	// Deprecated.
	Content() *string
	// The configuration content type.
	// Deprecated.
	ContentType() *string
}

// The jsii proxy struct for ConfigurationContent
type jsiiProxy_ConfigurationContent struct {
	_ byte // padding
}

func (j *jsiiProxy_ConfigurationContent) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigurationContent) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}


// Deprecated.
func NewConfigurationContent_Override(c ConfigurationContent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationContent",
		nil, // no parameters
		c,
	)
}

// Defines the hosted configuration content from a file.
// Deprecated.
func ConfigurationContent_FromFile(inputPath *string, contentType *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromFileParameters(inputPath); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationContent",
		"fromFile",
		[]interface{}{inputPath, contentType},
		&returns,
	)

	return returns
}

// Defines the hosted configuration content from inline code.
// Deprecated.
func ConfigurationContent_FromInline(content *string, contentType *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromInlineParameters(content); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationContent",
		"fromInline",
		[]interface{}{content, contentType},
		&returns,
	)

	return returns
}

// Defines the hosted configuration content as JSON from inline code.
// Deprecated.
func ConfigurationContent_FromInlineJson(content *string, contentType *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromInlineJsonParameters(content); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationContent",
		"fromInlineJson",
		[]interface{}{content, contentType},
		&returns,
	)

	return returns
}

// Defines the hosted configuration content as text from inline code.
// Deprecated.
func ConfigurationContent_FromInlineText(content *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromInlineTextParameters(content); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationContent",
		"fromInlineText",
		[]interface{}{content},
		&returns,
	)

	return returns
}

// Defines the hosted configuration content as YAML from inline code.
// Deprecated.
func ConfigurationContent_FromInlineYaml(content *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromInlineYamlParameters(content); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationContent",
		"fromInlineYaml",
		[]interface{}{content},
		&returns,
	)

	return returns
}

