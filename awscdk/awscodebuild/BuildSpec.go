package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// BuildSpec for CodeBuild projects.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   codebuildProject := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	ProjectName: jsii.String("MyTestProject"),
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("echo \"Hello, CodeBuild!\""),
//   				},
//   			},
//   		},
//   	}),
//   })
//
//   task := tasks.NewCodeBuildStartBuild(this, jsii.String("Task"), &CodeBuildStartBuildProps{
//   	Project: codebuildProject,
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   	EnvironmentVariablesOverride: map[string]buildEnvironmentVariable{
//   		"ZONE": &buildEnvironmentVariable{
//   			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			"value": sfn.JsonPath_stringAt(jsii.String("$.envVariables.zone")),
//   		},
//   	},
//   })
//
type BuildSpec interface {
	// Whether the buildspec is directly available or deferred until build-time.
	IsImmediate() *bool
	// Render the represented BuildSpec.
	ToBuildSpec(scope constructs.Construct) *string
}

// The jsii proxy struct for BuildSpec
type jsiiProxy_BuildSpec struct {
	_ byte // padding
}

func (j *jsiiProxy_BuildSpec) IsImmediate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isImmediate",
		&returns,
	)
	return returns
}


func NewBuildSpec_Override(b BuildSpec) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		nil, // no parameters
		b,
	)
}

// Use the contents of a local file as the build spec string.
//
// Use this if you have a local .yml or .json file that you want to use as the buildspec
func BuildSpec_FromAsset(path *string) BuildSpec {
	_init_.Initialize()

	if err := validateBuildSpec_FromAssetParameters(path); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		"fromAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func BuildSpec_FromObject(value *map[string]interface{}) BuildSpec {
	_init_.Initialize()

	if err := validateBuildSpec_FromObjectParameters(value); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		"fromObject",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Create a buildspec from an object that will be rendered as YAML in the resulting CloudFormation template.
func BuildSpec_FromObjectToYaml(value *map[string]interface{}) BuildSpec {
	_init_.Initialize()

	if err := validateBuildSpec_FromObjectToYamlParameters(value); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		"fromObjectToYaml",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Use a file from the source as buildspec.
//
// Use this if you want to use a file different from 'buildspec.yml'`
func BuildSpec_FromSourceFilename(filename *string) BuildSpec {
	_init_.Initialize()

	if err := validateBuildSpec_FromSourceFilenameParameters(filename); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		"fromSourceFilename",
		[]interface{}{filename},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildSpec) ToBuildSpec(scope constructs.Construct) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toBuildSpec",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

