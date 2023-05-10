package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// BuildSpec for CodeBuild projects.
//
// Example:
//   // later:
//   var project pipelineProject
//   sourceOutput := codepipeline.NewArtifact()
//   buildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("Build1"),
//   	Input: sourceOutput,
//   	Project: codebuild.NewPipelineProject(this, jsii.String("Project"), &PipelineProjectProps{
//   		BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
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
//   	VariablesNamespace: jsii.String("MyNamespace"),
//   })
//   codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("CodeBuild"),
//   	Project: Project,
//   	Input: sourceOutput,
//   	EnvironmentVariables: map[string]buildEnvironmentVariable{
//   		"MyVar": &buildEnvironmentVariable{
//   			"value": buildAction.variable(jsii.String("MY_VAR")),
//   		},
//   	},
//   })
//
// Experimental.
type BuildSpec interface {
	// Whether the buildspec is directly available or deferred until build-time.
	// Experimental.
	IsImmediate() *bool
	// Render the represented BuildSpec.
	// Experimental.
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


// Experimental.
func NewBuildSpec_Override(b BuildSpec) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codebuild.BuildSpec",
		nil, // no parameters
		b,
	)
}

// Experimental.
func BuildSpec_FromObject(value *map[string]interface{}) BuildSpec {
	_init_.Initialize()

	if err := validateBuildSpec_FromObjectParameters(value); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.BuildSpec",
		"fromObject",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Create a buildspec from an object that will be rendered as YAML in the resulting CloudFormation template.
// Experimental.
func BuildSpec_FromObjectToYaml(value *map[string]interface{}) BuildSpec {
	_init_.Initialize()

	if err := validateBuildSpec_FromObjectToYamlParameters(value); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.BuildSpec",
		"fromObjectToYaml",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Use a file from the source as buildspec.
//
// Use this if you want to use a file different from 'buildspec.yml'`
// Experimental.
func BuildSpec_FromSourceFilename(filename *string) BuildSpec {
	_init_.Initialize()

	if err := validateBuildSpec_FromSourceFilenameParameters(filename); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.BuildSpec",
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

