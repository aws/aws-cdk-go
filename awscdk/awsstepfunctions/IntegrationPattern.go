package awsstepfunctions


// AWS Step Functions integrates with services directly in the Amazon States Language.
//
// You can control these AWS services using service integration patterns:.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   project := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	ProjectName: jsii.String("MyTestProject"),
//   	BuildSpec: codebuild.BuildSpec_FromObjectToYaml(map[string]interface{}{
//   		"version": jsii.Number(0.2),
//   		"batch": map[string][]map[string]*string{
//   			"build-list": []map[string]*string{
//   				map[string]*string{
//   					"identifier": jsii.String("id"),
//   					"buildspec": jsii.String("version: 0.2\nphases:\n  build:\n    commands:\n      - echo \"Hello, from small!\""),
//   				},
//   			},
//   		},
//   	}),
//   })
//   project.EnableBatchBuilds()
//
//   task := tasks.NewCodeBuildStartBuildBatch(this, jsii.String("buildBatchTask"), &CodeBuildStartBuildBatchProps{
//   	Project: Project,
//   	IntegrationPattern: sfn.IntegrationPattern_REQUEST_RESPONSE,
//   	EnvironmentVariablesOverride: map[string]buildEnvironmentVariable{
//   		"test": &buildEnvironmentVariable{
//   			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			"value": jsii.String("testValue"),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html
//
type IntegrationPattern string

const (
	// Step Functions will wait for an HTTP response and then progress to the next state.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-default
	//
	IntegrationPattern_REQUEST_RESPONSE IntegrationPattern = "REQUEST_RESPONSE"
	// Step Functions can wait for a request to complete before progressing to the next state.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync
	//
	IntegrationPattern_RUN_JOB IntegrationPattern = "RUN_JOB"
	// Callback tasks provide a way to pause a workflow until a task token is returned.
	//
	// You must set a task token when using the callback pattern.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern_WAIT_FOR_TASK_TOKEN IntegrationPattern = "WAIT_FOR_TASK_TOKEN"
)

