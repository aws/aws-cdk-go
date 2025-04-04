package awsstepfunctions


// Properties for configuring a Distribute Map state that using JSONata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var itemBatcher itemBatcher
//   var itemReader iItemReader
//   var itemSelector interface{}
//   var outputs interface{}
//   var provideItems provideItems
//   var resultWriter resultWriter
//
//   distributedMapJsonataProps := &DistributedMapJsonataProps{
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Comment: jsii.String("comment"),
//   	ItemBatcher: itemBatcher,
//   	ItemReader: itemReader,
//   	Items: provideItems,
//   	ItemSelector: map[string]interface{}{
//   		"itemSelectorKey": itemSelector,
//   	},
//   	Label: jsii.String("label"),
//   	MapExecutionType: awscdk.Aws_stepfunctions.StateMachineType_EXPRESS,
//   	MaxConcurrency: jsii.Number(123),
//   	Outputs: outputs,
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	ResultWriter: resultWriter,
//   	StateName: jsii.String("stateName"),
//   	ToleratedFailureCount: jsii.Number(123),
//   	ToleratedFailurePercentage: jsii.Number(123),
//   }
//
type DistributedMapJsonataProps struct {
	// A comment describing this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The name of the query language used by the state.
	//
	// If the state does not contain a `queryLanguage` field,
	// then it will use the query language specified in the top-level `queryLanguage` field.
	// Default: - JSONPath.
	//
	QueryLanguage QueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// The JSON that you want to override your default iteration input (mutually exclusive  with `parameters`).
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-itemselector.html
	//
	// Default: $.
	//
	ItemSelector *map[string]interface{} `field:"optional" json:"itemSelector" yaml:"itemSelector"`
	// MaxConcurrency.
	//
	// An upper bound on the number of iterations you want running at once.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-asl-use-map-state-inline.html#map-state-inline-additional-fields
	//
	// Default: - full concurrency.
	//
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// Used to specify and transform output from the state.
	//
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
	// The array that the Map state will iterate over.
	// Default: - The state input as is.
	//
	Items ProvideItems `field:"optional" json:"items" yaml:"items"`
	// Specifies to process a group of items in a single child workflow execution.
	// Default: - No itemBatcher.
	//
	ItemBatcher ItemBatcher `field:"optional" json:"itemBatcher" yaml:"itemBatcher"`
	// ItemReader.
	//
	// Configuration for where to read items dataset in S3 to iterate.
	// Default: - No itemReader.
	//
	ItemReader IItemReader `field:"optional" json:"itemReader" yaml:"itemReader"`
	// Label.
	//
	// Unique name for the Distributed Map state added to each Map Run.
	// Default: - No label.
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// MapExecutionType.
	//
	// The execution type of the distributed map state
	//
	// This property overwrites ProcessorConfig.executionType
	// Default: StateMachineType.STANDARD
	//
	MapExecutionType StateMachineType `field:"optional" json:"mapExecutionType" yaml:"mapExecutionType"`
	// Configuration for S3 location in which to save Map Run results.
	// Default: - No resultWriter.
	//
	ResultWriter ResultWriter `field:"optional" json:"resultWriter" yaml:"resultWriter"`
	// ToleratedFailureCount.
	//
	// Number of failed items to tolerate in a Map Run, as static number.
	// Default: - No toleratedFailureCount.
	//
	ToleratedFailureCount *float64 `field:"optional" json:"toleratedFailureCount" yaml:"toleratedFailureCount"`
	// ToleratedFailurePercentage.
	//
	// Percentage of failed items to tolerate in a Map Run, as static number.
	// Default: - No toleratedFailurePercentage.
	//
	ToleratedFailurePercentage *float64 `field:"optional" json:"toleratedFailurePercentage" yaml:"toleratedFailurePercentage"`
}

