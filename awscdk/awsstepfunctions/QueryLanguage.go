package awsstepfunctions


// The name of the query language used by the state machine or state.
//
// Example:
//   jsonata := sfn.Pass_Jsonata(this, jsii.String("JSONata"))
//   jsonPath := sfn.Pass_JsonPath(this, jsii.String("JSONPath"))
//   definition := jsonata.Next(jsonPath)
//
//   sfn.NewStateMachine(this, jsii.String("MixedStateMachine"), &StateMachineProps{
//   	// queryLanguage: sfn.QueryLanguage.JSON_PATH, // default
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
//   })
//
//   // This throws an error. If JSONata is specified at the top level, JSONPath cannot be used in the state machine definition.
//   // This throws an error. If JSONata is specified at the top level, JSONPath cannot be used in the state machine definition.
//   sfn.NewStateMachine(this, jsii.String("JSONataOnlyStateMachine"), &StateMachineProps{
//   	QueryLanguage: sfn.QueryLanguage_JSONATA,
//   	DefinitionBody: sfn.DefinitionBody_*FromChainable(definition),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/transforming-data.html
//
// Default: JSON_PATH.
//
type QueryLanguage string

const (
	// Use JSONPath.
	QueryLanguage_JSON_PATH QueryLanguage = "JSON_PATH"
	// Use JSONata.
	QueryLanguage_JSONATA QueryLanguage = "JSONATA"
)

