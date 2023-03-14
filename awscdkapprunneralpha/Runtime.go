// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The code runtimes.
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
type Runtime interface {
	// The runtime name.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Other runtimes.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfigurationvalues.html#cfn-apprunner-service-codeconfigurationvalues-runtime for all available runtimes.
//
// Experimental.
func Runtime_Of(name *string) Runtime {
	_init_.Initialize()

	if err := validateRuntime_OfParameters(name); err != nil {
		panic(err)
	}
	var returns Runtime

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func Runtime_CORRETTO_11() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"CORRETTO_11",
		&returns,
	)
	return returns
}

func Runtime_CORRETTO_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"CORRETTO_8",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_6() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"DOTNET_6",
		&returns,
	)
	return returns
}

func Runtime_GO_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"GO_1",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_12() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"NODEJS_12",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_14() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"NODEJS_14",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_16() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"NODEJS_16",
		&returns,
	)
	return returns
}

func Runtime_PHP_81() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"PHP_81",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"PYTHON_3",
		&returns,
	)
	return returns
}

func Runtime_RUBY_31() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		"RUBY_31",
		&returns,
	)
	return returns
}

