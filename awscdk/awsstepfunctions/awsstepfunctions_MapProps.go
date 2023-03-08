package awsstepfunctions


// Properties for defining a Map state.
//
// Example:
//   map := sfn.NewMap(this, jsii.String("Map State"), &mapProps{
//   	maxConcurrency: jsii.Number(1),
//   	itemsPath: sfn.jsonPath.stringAt(jsii.String("$.inputForMap")),
//   })
//   map.iterator(sfn.NewPass(this, jsii.String("Pass State")))
//
// Experimental.
type MapProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select the array to iterate over.
	// Experimental.
	ItemsPath *string `field:"optional" json:"itemsPath" yaml:"itemsPath"`
	// MaxConcurrency.
	//
	// An upper bound on the number of iterations you want running at once.
	// Experimental.
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// The JSON that you want to override your default iteration input.
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
}

