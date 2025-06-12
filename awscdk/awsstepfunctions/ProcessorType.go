package awsstepfunctions


// Execution type for the Map workflow.
//
// Example:
//   map := sfn.NewMap(this, jsii.String("Map State"), &MapProps{
//   	MaxConcurrency: jsii.Number(1),
//   	ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.inputForMap")),
//   	ItemSelector: map[string]interface{}{
//   		"item": sfn.JsonPath_*StringAt(jsii.String("$.Map.Item.Value")),
//   	},
//   	ResultPath: jsii.String("$.mapOutput"),
//   })
//
//   map.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")), &ProcessorConfig{
//   	Mode: sfn.ProcessorMode_DISTRIBUTED,
//   	ExecutionType: sfn.ProcessorType_STANDARD,
//   })
//
type ProcessorType string

const (
	// Standard execution type.
	ProcessorType_STANDARD ProcessorType = "STANDARD"
	// Express execution type.
	ProcessorType_EXPRESS ProcessorType = "EXPRESS"
)

