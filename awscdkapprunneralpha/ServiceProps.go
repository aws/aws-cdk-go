package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties of the AppRunner Service.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromGitHub(&GithubRepositoryProps{
//   		RepositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		Branch: jsii.String("main"),
//   		ConfigurationSource: apprunner.ConfigurationSourceType_API,
//   		CodeConfigurationValues: &CodeConfigurationValues{
//   			Runtime: apprunner.Runtime_PYTHON_3(),
//   			Port: jsii.String("8000"),
//   			StartCommand: jsii.String("python app.py"),
//   			BuildCommand: jsii.String("yum install -y pycairo && pip install -r requirements.txt"),
//   		},
//   		Connection: apprunner.GitHubConnection_FromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type ServiceProps struct {
	// The source of the repository for the service.
	// Experimental.
	Source Source `field:"required" json:"source" yaml:"source"`
	// The IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	//
	// The role must be assumable by the 'build.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.access
	//
	// Experimental.
	AccessRole awsiam.IRole `field:"optional" json:"accessRole" yaml:"accessRole"`
	// Specifies whether to enable continuous integration from the source repository.
	//
	// If true, continuous integration from the source repository is enabled for the App Runner service.
	// Each repository change (including any source code commit or new image version) starts a deployment.
	// By default, App Runner sets to false for a source image that uses an ECR Public repository or an ECR repository that's in an AWS account other than the one that the service is in.
	// App Runner sets to true in all other cases (which currently include a source code repository or a source image using a same-account ECR repository).
	// Experimental.
	AutoDeploymentsEnabled *bool `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// The number of CPU units reserved for each instance of your App Runner service.
	// Experimental.
	Cpu Cpu `field:"optional" json:"cpu" yaml:"cpu"`
	// The IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	//
	// The role must be assumable by the 'tasks.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.instance
	//
	// Experimental.
	InstanceRole awsiam.IRole `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// The amount of memory reserved for each instance of your App Runner service.
	// Experimental.
	Memory Memory `field:"optional" json:"memory" yaml:"memory"`
	// Name of the service.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Settings for an App Runner VPC connector to associate with the service.
	// Experimental.
	VpcConnector IVpcConnector `field:"optional" json:"vpcConnector" yaml:"vpcConnector"`
}

