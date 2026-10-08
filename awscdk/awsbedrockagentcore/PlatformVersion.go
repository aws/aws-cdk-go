package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The platform version of the Agent Runtime.
//
// Example:
//   repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("test-agent-runtime"),
//   })
//
//   agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))
//
//   agentcore.NewRuntime(this, jsii.String("test-runtime"), &RuntimeProps{
//   	RuntimeName: jsii.String("test_runtime"),
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//   	PlatformVersion: agentcore.PlatformVersion_V2(),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-platformversion
//
type PlatformVersion interface {
	// The platform version string value.
	Value() *string
	// Returns the string value.
	ToString() *string
}

// The jsii proxy struct for PlatformVersion
type jsiiProxy_PlatformVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_PlatformVersion) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom platform version not yet defined in this class.
func PlatformVersion_Of(value *string) PlatformVersion {
	_init_.Initialize()

	if err := validatePlatformVersion_OfParameters(value); err != nil {
		panic(err)
	}
	var returns PlatformVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PlatformVersion",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PlatformVersion_V1() PlatformVersion {
	_init_.Initialize()
	var returns PlatformVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.PlatformVersion",
		"V1",
		&returns,
	)
	return returns
}

func PlatformVersion_V2() PlatformVersion {
	_init_.Initialize()
	var returns PlatformVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.PlatformVersion",
		"V2",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PlatformVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

