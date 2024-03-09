package awsstepfunctions


// Properties for configuring a Distribute Map state.
//
// Example:
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("Distributed Map State"), &DistributedMapProps{
//   	MaxConcurrency: jsii.Number(1),
//   	ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.inputForMap")),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")))
//
type DistributedMapProps struct {
	// An optional description for this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: $.
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// The JSON that you want to override your default iteration input (mutually exclusive  with `parameters`).
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-itemselector.html
	//
	// Default: $.
	//
	ItemSelector *map[string]interface{} `field:"optional" json:"itemSelector" yaml:"itemSelector"`
	// JSONPath expression to select the array to iterate over.
	// Default: $.
	//
	ItemsPath *string `field:"optional" json:"itemsPath" yaml:"itemsPath"`
	// MaxConcurrency.
	//
	// An upper bound on the number of iterations you want running at once.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-asl-use-map-state-inline.html#map-state-inline-additional-fields
	//
	// Default: - full concurrency.
	//
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// MaxConcurrencyPath.
	//
	// A JsonPath that specifies the maximum concurrency dynamically from the state input.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-asl-use-map-state-inline.html#map-state-inline-additional-fields
	//
	// Default: - full concurrency.
	//
	MaxConcurrencyPath *string `field:"optional" json:"maxConcurrencyPath" yaml:"maxConcurrencyPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: $.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
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
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
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
	// The execution type of the distributed map state.
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

