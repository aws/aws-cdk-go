package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines the hosted configuration content.
//
// Example:
//   var application application
//
//
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
//   	Application: Application,
//   	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
//   	Type: appconfig.ConfigurationType_FEATURE_FLAGS,
//   })
//
type ConfigurationContent interface {
	// The configuration content.
	Content() *string
	// The configuration content type.
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


func NewConfigurationContent_Override(c ConfigurationContent) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.ConfigurationContent",
		nil, // no parameters
		c,
	)
}

// Defines the hosted configuration content from a file.
func ConfigurationContent_FromFile(inputPath *string, contentType *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromFileParameters(inputPath); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.ConfigurationContent",
		"fromFile",
		[]interface{}{inputPath, contentType},
		&returns,
	)

	return returns
}

// Defines the hosted configuration content from inline code.
func ConfigurationContent_FromInline(content *string, contentType *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromInlineParameters(content); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.ConfigurationContent",
		"fromInline",
		[]interface{}{content, contentType},
		&returns,
	)

	return returns
}

// Defines the hosted configuration content as JSON from inline code.
func ConfigurationContent_FromInlineJson(content *string, contentType *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromInlineJsonParameters(content); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.ConfigurationContent",
		"fromInlineJson",
		[]interface{}{content, contentType},
		&returns,
	)

	return returns
}

// Defines the hosted configuration content as text from inline code.
func ConfigurationContent_FromInlineText(content *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromInlineTextParameters(content); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.ConfigurationContent",
		"fromInlineText",
		[]interface{}{content},
		&returns,
	)

	return returns
}

// Defines the hosted configuration content as YAML from inline code.
func ConfigurationContent_FromInlineYaml(content *string) ConfigurationContent {
	_init_.Initialize()

	if err := validateConfigurationContent_FromInlineYamlParameters(content); err != nil {
		panic(err)
	}
	var returns ConfigurationContent

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.ConfigurationContent",
		"fromInlineYaml",
		[]interface{}{content},
		&returns,
	)

	return returns
}

