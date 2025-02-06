package awsstepfunctions


// Options for Choice Transition.
//
// Example:
//   import events "github.com/aws/aws-cdk-go/awscdk"
//   var connection connection
//
//
//   getIssue := tasks.HttpInvoke_Jsonata(this, jsii.String("Get Issue"), &HttpInvokeJsonataProps{
//   	Connection: Connection,
//   	ApiRoot: jsii.String("{% 'https://' & $states.input.hostname %}"),
//   	ApiEndpoint: sfn.TaskInput_FromText(jsii.String("{% 'issues/' & $states.input.issue.id %}")),
//   	Method: sfn.TaskInput_*FromText(jsii.String("GET")),
//   	// Parse the API call result to object and set to the variables
//   	Assign: map[string]interface{}{
//   		"hostname": jsii.String("{% $states.input.hostname %}"),
//   		"issue": jsii.String("{% $parse($states.result.ResponseBody) %}"),
//   	},
//   })
//
//   updateLabels := tasks.HttpInvoke_Jsonata(this, jsii.String("Update Issue Labels"), &HttpInvokeJsonataProps{
//   	Connection: Connection,
//   	ApiRoot: jsii.String("{% 'https://' & $states.input.hostname %}"),
//   	ApiEndpoint: sfn.TaskInput_*FromText(jsii.String("{% 'issues/' & $states.input.issue.id & 'labels' %}")),
//   	Method: sfn.TaskInput_*FromText(jsii.String("POST")),
//   	Body: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"labels": jsii.String("{% [$type, $component] %}"),
//   	}),
//   })
//   notMatchTitleTemplate := sfn.Pass_Jsonata(this, jsii.String("Not Match Title Template"))
//
//   definition := getIssue.Next(sfn.Choice_Jsonata(this, jsii.String("Match Title Template?")).When(sfn.Condition_Jsonata(jsii.String("{% $contains($issue.title, /(feat)|(fix)|(chore)(w*):.*/) %}")), updateLabels, &ChoiceTransitionOptions{
//   	Assign: map[string]interface{}{
//   		"type": jsii.String("{% $match($states.input.title, /(w*)((.*))/).groups[0] %}"),
//   		"component": jsii.String("{% $match($states.input.title, /(w*)((.*))/).groups[1] %}"),
//   	},
//   }).Otherwise(notMatchTitleTemplate))
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   	Comment: jsii.String("automate issue labeling state machine"),
//   })
//
type ChoiceTransitionOptions struct {
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// An optional description for the choice transition.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// This option for JSONata only.
	//
	// When you use JSONPath, then the state ignores this property.
	// Used to specify and transform output from the state.
	// When specified, the value overrides the state output default.
	// The output field accepts any JSON value (object, array, string, number, boolean, null).
	// Any string value, including those inside objects or arrays,
	// will be evaluated as JSONata if surrounded by {% %} characters.
	// Output also accepts a JSONata expression directly.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html
	//
	// Default: - $states.result or $states.errorOutput
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
}

