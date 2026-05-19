package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol configuration for Agent Runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protocolType := awscdk.Aws_bedrockagentcore.ProtocolType_A2A()
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-protocolconfiguration
//
type ProtocolType interface {
	// The protocol type string value.
	Value() *string
	// Returns the string value.
	ToString() *string
}

// The jsii proxy struct for ProtocolType
type jsiiProxy_ProtocolType struct {
	_ byte // padding
}

func (j *jsiiProxy_ProtocolType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom protocol type not yet defined in this class.
func ProtocolType_Of(value *string) ProtocolType {
	_init_.Initialize()

	if err := validateProtocolType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ProtocolType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.ProtocolType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ProtocolType_A2A() ProtocolType {
	_init_.Initialize()
	var returns ProtocolType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ProtocolType",
		"A2A",
		&returns,
	)
	return returns
}

func ProtocolType_AGUI() ProtocolType {
	_init_.Initialize()
	var returns ProtocolType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ProtocolType",
		"AGUI",
		&returns,
	)
	return returns
}

func ProtocolType_HTTP() ProtocolType {
	_init_.Initialize()
	var returns ProtocolType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ProtocolType",
		"HTTP",
		&returns,
	)
	return returns
}

func ProtocolType_MCP() ProtocolType {
	_init_.Initialize()
	var returns ProtocolType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ProtocolType",
		"MCP",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProtocolType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

