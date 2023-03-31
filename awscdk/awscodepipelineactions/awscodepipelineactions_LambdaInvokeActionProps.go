package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Construction properties of the {@link LambdaInvokeAction Lambda invoke CodePipeline Action}.
//
// Example:
//   var fn function
//
//   sourceOutput := codepipeline.NewArtifact()
//   buildOutput := codepipeline.NewArtifact()
//   lambdaAction := codepipeline_actions.NewLambdaInvokeAction(&lambdaInvokeActionProps{
//   	actionName: jsii.String("Lambda"),
//   	inputs: []artifact{
//   		sourceOutput,
//   		buildOutput,
//   	},
//   	outputs: []*artifact{
//   		codepipeline.NewArtifact(jsii.String("Out1")),
//   		codepipeline.NewArtifact(jsii.String("Out2")),
//   	},
//   	lambda: fn,
//   })
//
// Experimental.
type LambdaInvokeActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The lambda function to invoke.
	// Experimental.
	Lambda awslambda.IFunction `field:"required" json:"lambda" yaml:"lambda"`
	// The optional input Artifacts of the Action.
	//
	// A Lambda Action can have up to 5 inputs.
	// The inputs will appear in the event passed to the Lambda,
	// under the `'CodePipeline.job'.data.inputArtifacts` path.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html#actions-invoke-lambda-function-json-event-example
	//
	// Experimental.
	Inputs *[]awscodepipeline.Artifact `field:"optional" json:"inputs" yaml:"inputs"`
	// The optional names of the output Artifacts of the Action.
	//
	// A Lambda Action can have up to 5 outputs.
	// The outputs will appear in the event passed to the Lambda,
	// under the `'CodePipeline.job'.data.outputArtifacts` path.
	// It is the responsibility of the Lambda to upload ZIP files with the Artifact contents to the provided locations.
	// Experimental.
	Outputs *[]awscodepipeline.Artifact `field:"optional" json:"outputs" yaml:"outputs"`
	// A set of key-value pairs that will be accessible to the invoked Lambda inside the event that the Pipeline will call it with.
	//
	// Only one of `userParameters` or `userParametersString` can be specified.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-invoke-lambda-function.html#actions-invoke-lambda-function-json-event-example
	//
	// Experimental.
	UserParameters *map[string]interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
	// The string representation of the user parameters that will be accessible to the invoked Lambda inside the event that the Pipeline will call it with.
	//
	// Only one of `userParametersString` or `userParameters` can be specified.
	// Experimental.
	UserParametersString *string `field:"optional" json:"userParametersString" yaml:"userParametersString"`
}

