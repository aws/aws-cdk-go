// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties of the AppRunner Service.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromGitHub(&githubRepositoryProps{
//   		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		branch: jsii.String("main"),
//   		configurationSource: apprunner.configurationSourceType_API,
//   		codeConfigurationValues: &codeConfigurationValues{
//   			runtime: apprunner.runtime_PYTHON_3(),
//   			port: jsii.String("8000"),
//   			startCommand: jsii.String("python app.py"),
//   			buildCommand: jsii.String("yum install -y pycairo && pip install -r requirements.txt"),
//   		},
//   		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
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

