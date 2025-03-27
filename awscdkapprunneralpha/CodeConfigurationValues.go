package awscdkapprunneralpha


// Describes the basic configuration needed for building and running an AWS App Runner service.
//
// This type doesn't support the full set of possible configuration options. Fur full configuration capabilities,
// use a `apprunner.yaml` file in the source code repository.
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
type CodeConfigurationValues struct {
	// A runtime environment type for building and running an App Runner service.
	//
	// It represents
	// a programming language runtime.
	// Experimental.
	Runtime Runtime `field:"required" json:"runtime" yaml:"runtime"`
	// The command App Runner runs to build your application.
	// Default: - no build command.
	//
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The environment variables that are available to your running App Runner service.
	// Default: - no environment variables.
	//
	// Deprecated: use environmentVariables.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The environment secrets that are available to your running App Runner service.
	// Default: - no environment secrets.
	//
	// Experimental.
	EnvironmentSecrets *map[string]Secret `field:"optional" json:"environmentSecrets" yaml:"environmentSecrets"`
	// The environment variables that are available to your running App Runner service.
	// Default: - no environment variables.
	//
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The port that your application listens to in the container.
	// Default: 8080.
	//
	// Experimental.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The command App Runner runs to start your application.
	// Default: - no start command.
	//
	// Experimental.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

