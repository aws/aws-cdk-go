package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// BuildSpec for CodeBuild projects.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // later:
//   var project pipelineProject
//   sourceOutput := codepipeline.NewArtifact()
//   buildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("Build1"),
//   	input: sourceOutput,
//   	project: codebuild.NewPipelineProject(this, jsii.String("Project"), &pipelineProjectProps{
//   		buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   			"version": jsii.String("0.2"),
//   			"env": map[string][]*string{
//   				"exported-variables": []*string{
//   					jsii.String("MY_VAR"),
//   				},
//   			},
//   			"phases": map[string]map[string]*string{
//   				"build": map[string]*string{
//   					"commands": jsii.String("export MY_VAR=\"some value\""),
//   				},
//   			},
//   		}),
//   	}),
//   	variablesNamespace: jsii.String("MyNamespace"),
//   })
//   codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("CodeBuild"),
//   	project: project,
//   	input: sourceOutput,
//   	environmentVariables: map[string]buildEnvironmentVariable{
//   		"MyVar": &buildEnvironmentVariable{
//   			"value": buildAction.variable(jsii.String("MY_VAR")),
//   		},
//   	},
//   })
//
type BuildSpec interface {
	// Whether the buildspec is directly available or deferred until build-time.
	IsImmediate() *bool
	// Render the represented BuildSpec.
	ToBuildSpec() *string
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

func BuildSpec_FromObject(value *map[string]interface{}) BuildSpec {
	_init_.Initialize()

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

	var returns BuildSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		"fromSourceFilename",
		[]interface{}{filename},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildSpec) ToBuildSpec() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toBuildSpec",
		nil, // no parameters
		&returns,
	)

	return returns
}

