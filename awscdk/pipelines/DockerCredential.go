package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Represents credentials used to access a Docker registry.
//
// Example:
//   dockerHubSecret := secretsmanager.Secret_FromSecretCompleteArn(this, jsii.String("DHSecret"), jsii.String("arn:aws:..."))
//   customRegSecret := secretsmanager.Secret_FromSecretCompleteArn(this, jsii.String("CRSecret"), jsii.String("arn:aws:..."))
//   repo1 := ecr.Repository_FromRepositoryArn(this, jsii.String("Repo"), jsii.String("arn:aws:ecr:eu-west-1:0123456789012:repository/Repo1"))
//   repo2 := ecr.Repository_FromRepositoryArn(this, jsii.String("Repo"), jsii.String("arn:aws:ecr:eu-west-1:0123456789012:repository/Repo2"))
//
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
//   	DockerCredentials: []dockerCredential{
//   		pipelines.*dockerCredential_DockerHub(dockerHubSecret),
//   		pipelines.*dockerCredential_CustomRegistry(jsii.String("dockerregistry.example.com"), customRegSecret),
//   		pipelines.*dockerCredential_Ecr([]iRepository{
//   			repo1,
//   			repo2,
//   		}),
//   	},
//   	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
//   		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
//   			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		Commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//   })
//
type DockerCredential interface {
	Usages() *[]DockerCredentialUsage
	// Grant read-only access to the registry credentials.
	//
	// This grants read access to any secrets, and pull access to any repositories.
	GrantRead(grantee awsiam.IGrantable, usage DockerCredentialUsage)
}

// The jsii proxy struct for DockerCredential
type jsiiProxy_DockerCredential struct {
	_ byte // padding
}

func (j *jsiiProxy_DockerCredential) Usages() *[]DockerCredentialUsage {
	var returns *[]DockerCredentialUsage
	_jsii_.Get(
		j,
		"usages",
		&returns,
	)
	return returns
}


func NewDockerCredential_Override(d DockerCredential, usages *[]DockerCredentialUsage) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.DockerCredential",
		[]interface{}{usages},
		d,
	)
}

// Creates a DockerCredential for a registry, based on its domain name (e.g., 'www.example.com').
func DockerCredential_CustomRegistry(registryDomain *string, secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) DockerCredential {
	_init_.Initialize()

	if err := validateDockerCredential_CustomRegistryParameters(registryDomain, secret, opts); err != nil {
		panic(err)
	}
	var returns DockerCredential

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.DockerCredential",
		"customRegistry",
		[]interface{}{registryDomain, secret, opts},
		&returns,
	)

	return returns
}

// Creates a DockerCredential for DockerHub.
//
// Convenience method for `customRegistry('https://index.docker.io/v1/', opts)`.
func DockerCredential_DockerHub(secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) DockerCredential {
	_init_.Initialize()

	if err := validateDockerCredential_DockerHubParameters(secret, opts); err != nil {
		panic(err)
	}
	var returns DockerCredential

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.DockerCredential",
		"dockerHub",
		[]interface{}{secret, opts},
		&returns,
	)

	return returns
}

// Creates a DockerCredential for one or more ECR repositories.
//
// NOTE - All ECR repositories in the same account and region share a domain name
// (e.g., 0123456789012.dkr.ecr.eu-west-1.amazonaws.com), and can only have one associated
// set of credentials (and DockerCredential). Attempting to associate one set of credentials
// with one ECR repo and another with another ECR repo in the same account and region will
// result in failures when using these credentials in the pipeline.
func DockerCredential_Ecr(repositories *[]awsecr.IRepository, opts *EcrDockerCredentialOptions) DockerCredential {
	_init_.Initialize()

	if err := validateDockerCredential_EcrParameters(repositories, opts); err != nil {
		panic(err)
	}
	var returns DockerCredential

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.DockerCredential",
		"ecr",
		[]interface{}{repositories, opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerCredential) GrantRead(grantee awsiam.IGrantable, usage DockerCredentialUsage) {
	if err := d.validateGrantReadParameters(grantee, usage); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"grantRead",
		[]interface{}{grantee, usage},
	)
}

