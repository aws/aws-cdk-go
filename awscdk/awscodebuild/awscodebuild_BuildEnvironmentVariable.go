package awscodebuild


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
type BuildEnvironmentVariable struct {
	// The value of the environment variable.
	//
	// For plain-text variables (the default), this is the literal value of variable.
	// For SSM parameter variables, pass the name of the parameter here (`parameterName` property of `IParameter`).
	// For SecretsManager variables secrets, pass either the secret name (`secretName` property of `ISecret`)
	// or the secret ARN (`secretArn` property of `ISecret`) here,
	// along with optional SecretsManager qualifiers separated by ':', like the JSON key, or the version or stage
	// (see https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec.env.secrets-manager for details).
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// The type of environment variable.
	Type BuildEnvironmentVariableType `field:"optional" json:"type" yaml:"type"`
}

