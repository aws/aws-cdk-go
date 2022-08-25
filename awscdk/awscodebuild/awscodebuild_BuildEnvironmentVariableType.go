package awscodebuild


// Example:
//   import codebuild "github.com/aws/aws-cdk-go/awscdk"
//
//
//   codebuildProject := codebuild.NewProject(this, jsii.String("Project"), &projectProps{
//   	projectName: jsii.String("MyTestProject"),
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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
//   task := tasks.NewCodeBuildStartBuild(this, jsii.String("Task"), &codeBuildStartBuildProps{
//   	project: codebuildProject,
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   	environmentVariablesOverride: map[string]buildEnvironmentVariable{
//   		"ZONE": &buildEnvironmentVariable{
//   			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			"value": sfn.JsonPath.stringAt(jsii.String("$.envVariables.zone")),
//   		},
//   	},
//   })
//
type BuildEnvironmentVariableType string

const (
	// An environment variable in plaintext format.
	BuildEnvironmentVariableType_PLAINTEXT BuildEnvironmentVariableType = "PLAINTEXT"
	// An environment variable stored in Systems Manager Parameter Store.
	BuildEnvironmentVariableType_PARAMETER_STORE BuildEnvironmentVariableType = "PARAMETER_STORE"
	// An environment variable stored in AWS Secrets Manager.
	BuildEnvironmentVariableType_SECRETS_MANAGER BuildEnvironmentVariableType = "SECRETS_MANAGER"
)

