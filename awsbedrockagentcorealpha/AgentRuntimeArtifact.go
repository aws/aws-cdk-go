package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Abstract base class for agent runtime artifacts.
//
// Provides methods to reference container images from ECR repositories or local assets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkMode NetworkMode
//   var platform Platform
//
//   agentRuntimeArtifact := bedrock_agentcore_alpha.AgentRuntimeArtifact_FromAsset(jsii.String("directory"), &DockerImageAssetOptions{
//   	AssetName: jsii.String("assetName"),
//   	BuildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	BuildContexts: map[string]*string{
//   		"buildContextsKey": jsii.String("buildContexts"),
//   	},
//   	BuildSecrets: map[string]*string{
//   		"buildSecretsKey": jsii.String("buildSecrets"),
//   	},
//   	BuildSsh: jsii.String("buildSsh"),
//   	CacheDisabled: jsii.Boolean(false),
//   	CacheFrom: []DockerCacheOption{
//   		&DockerCacheOption{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Params: map[string]*string{
//   				"paramsKey": jsii.String("params"),
//   			},
//   		},
//   	},
//   	CacheTo: &DockerCacheOption{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Params: map[string]*string{
//   			"paramsKey": jsii.String("params"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	ExtraHash: jsii.String("extraHash"),
//   	File: jsii.String("file"),
//   	FollowSymlinks: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   	Invalidation: &DockerImageAssetInvalidationOptions{
//   		BuildArgs: jsii.Boolean(false),
//   		BuildContexts: jsii.Boolean(false),
//   		BuildSecrets: jsii.Boolean(false),
//   		BuildSsh: jsii.Boolean(false),
//   		ExtraHash: jsii.Boolean(false),
//   		File: jsii.Boolean(false),
//   		NetworkMode: jsii.Boolean(false),
//   		Outputs: jsii.Boolean(false),
//   		Platform: jsii.Boolean(false),
//   		RepositoryName: jsii.Boolean(false),
//   		Target: jsii.Boolean(false),
//   	},
//   	NetworkMode: networkMode,
//   	Outputs: []*string{
//   		jsii.String("outputs"),
//   	},
//   	Platform: platform,
//   	Target: jsii.String("target"),
//   })
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type AgentRuntimeArtifact interface {
	// Called when the image is used by a Runtime to handle side effects like permissions.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Bind(scope constructs.Construct, runtime Runtime)
}

// The jsii proxy struct for AgentRuntimeArtifact
type jsiiProxy_AgentRuntimeArtifact struct {
	_ byte // padding
}

// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func NewAgentRuntimeArtifact_Override(a AgentRuntimeArtifact) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentRuntimeArtifact",
		nil, // no parameters
		a,
	)
}

// Reference an agent runtime artifact that's constructed directly from sources on disk.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func AgentRuntimeArtifact_FromAsset(directory *string, options *awsecrassets.DockerImageAssetOptions) AgentRuntimeArtifact {
	_init_.Initialize()

	if err := validateAgentRuntimeArtifact_FromAssetParameters(directory, options); err != nil {
		panic(err)
	}
	var returns AgentRuntimeArtifact

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentRuntimeArtifact",
		"fromAsset",
		[]interface{}{directory, options},
		&returns,
	)

	return returns
}

// Reference an agent runtime artifact that's constructed from local code assets uploaded to a CDK-managed S3 bucket.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func AgentRuntimeArtifact_FromCodeAsset(options *CodeAssetOptions) AgentRuntimeArtifact {
	_init_.Initialize()

	if err := validateAgentRuntimeArtifact_FromCodeAssetParameters(options); err != nil {
		panic(err)
	}
	var returns AgentRuntimeArtifact

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentRuntimeArtifact",
		"fromCodeAsset",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func AgentRuntimeArtifact_FromEcrRepository(repository awsecr.IRepository, tag *string) AgentRuntimeArtifact {
	_init_.Initialize()

	if err := validateAgentRuntimeArtifact_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns AgentRuntimeArtifact

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentRuntimeArtifact",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image using an ECR container URI.
//
// Use this when referencing ECR images from CloudFormation parameters or cross-stack references.
//
// **Note:** No IAM permissions are automatically granted. You must ensure the runtime has
// ECR pull permissions for the repository.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func AgentRuntimeArtifact_FromImageUri(containerUri *string) AgentRuntimeArtifact {
	_init_.Initialize()

	if err := validateAgentRuntimeArtifact_FromImageUriParameters(containerUri); err != nil {
		panic(err)
	}
	var returns AgentRuntimeArtifact

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentRuntimeArtifact",
		"fromImageUri",
		[]interface{}{containerUri},
		&returns,
	)

	return returns
}

// Reference an agent runtime artifact that's constructed directly from an S3 object.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func AgentRuntimeArtifact_FromS3(s3Location *awss3.Location, runtime AgentCoreRuntime, entrypoint *[]*string) AgentRuntimeArtifact {
	_init_.Initialize()

	if err := validateAgentRuntimeArtifact_FromS3Parameters(s3Location, runtime, entrypoint); err != nil {
		panic(err)
	}
	var returns AgentRuntimeArtifact

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentRuntimeArtifact",
		"fromS3",
		[]interface{}{s3Location, runtime, entrypoint},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentRuntimeArtifact) Bind(scope constructs.Construct, runtime Runtime) {
	if err := a.validateBindParameters(scope, runtime); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{scope, runtime},
	)
}

