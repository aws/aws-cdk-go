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
//   // S3 bucket containing the agent core
//   codeBucket := s3.NewBucket(this, jsii.String("AgentCode"), &BucketProps{
//   	BucketName: jsii.String("my-code-bucket"),
//   	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
//   })
//
//   // the bucket above needs to contain the agent code
//
//   agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromS3(&Location{
//   	BucketName: codeBucket.BucketName,
//   	ObjectKey: jsii.String("deployment_package.zip"),
//   }, agentcore.AgentCoreRuntime_PYTHON_3_12, []*string{
//   	jsii.String("opentelemetry-instrument"),
//   	jsii.String("main.py"),
//   })
//
//   runtimeInstance := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
//   	RuntimeName: jsii.String("myAgent"),
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//   })
//
// Experimental.
type AgentRuntimeArtifact interface {
	// Called when the image is used by a Runtime to handle side effects like permissions.
	// Experimental.
	Bind(scope constructs.Construct, runtime Runtime)
}

// The jsii proxy struct for AgentRuntimeArtifact
type jsiiProxy_AgentRuntimeArtifact struct {
	_ byte // padding
}

// Experimental.
func NewAgentRuntimeArtifact_Override(a AgentRuntimeArtifact) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AgentRuntimeArtifact",
		nil, // no parameters
		a,
	)
}

// Reference an agent runtime artifact that's constructed directly from sources on disk.
// Experimental.
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

// Reference an image in an ECR repository.
// Experimental.
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

// Reference an agent runtime artifact that's constructed directly from an S3 object.
// Experimental.
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

