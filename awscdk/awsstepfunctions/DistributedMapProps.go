package awsstepfunctions


// Properties for configuring a Distribute Map state.
//
// Example:
//   /**
//    * JSON state input:
//    *  {
//    *    "bucketName": "my-bucket",
//    *    "prefix": "item"
//    *  }
//    */
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("DistributedMap"), &DistributedMapProps{
//   	ItemReader: sfn.NewS3ObjectsItemReader(&S3ObjectsItemReaderProps{
//   		BucketNamePath: sfn.JsonPath_StringAt(jsii.String("$.bucketName")),
//   		Prefix: sfn.JsonPath_*StringAt(jsii.String("$.prefix")),
//   	}),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass")))
//
type DistributedMapProps struct {
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
	// See: https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/workflow-variables.html
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
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: $.
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: $.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to select the array to iterate over.
	// Default: $.
	//
	ItemsPath *string `field:"optional" json:"itemsPath" yaml:"itemsPath"`
	// MaxConcurrencyPath.
	//
	// A JsonPath that specifies the maximum concurrency dynamically from the state input.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-asl-use-map-state-inline.html#map-state-inline-additional-fields
	//
	// Default: - full concurrency.
	//
	MaxConcurrencyPath *string `field:"optional" json:"maxConcurrencyPath" yaml:"maxConcurrencyPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: $.
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Default: - None.
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
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
	// ToleratedFailureCountPath.
	//
	// Number of failed items to tolerate in a Map Run, as JsonPath.
	// Default: - No toleratedFailureCountPath.
	//
	ToleratedFailureCountPath *string `field:"optional" json:"toleratedFailureCountPath" yaml:"toleratedFailureCountPath"`
	// ToleratedFailurePercentage.
	//
	// Percentage of failed items to tolerate in a Map Run, as static number.
	// Default: - No toleratedFailurePercentage.
	//
	ToleratedFailurePercentage *float64 `field:"optional" json:"toleratedFailurePercentage" yaml:"toleratedFailurePercentage"`
	// ToleratedFailurePercentagePath.
	//
	// Percentage of failed items to tolerate in a Map Run, as JsonPath.
	// Default: - No toleratedFailurePercentagePath.
	//
	ToleratedFailurePercentagePath *string `field:"optional" json:"toleratedFailurePercentagePath" yaml:"toleratedFailurePercentagePath"`
}

