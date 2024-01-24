package awsstepfunctions


// Specifies the configuration for the processor Map state.
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
type ProcessorConfig struct {
	// Specifies the execution type for the Map workflow.
	//
	// You must provide this field if you specified `DISTRIBUTED` for the `mode` sub-field.
	// Default: - no execution type.
	//
	ExecutionType ProcessorType `field:"optional" json:"executionType" yaml:"executionType"`
	// Specifies the execution mode for the Map workflow.
	// Default: - ProcessorMode.INLINE
	//
	Mode ProcessorMode `field:"optional" json:"mode" yaml:"mode"`
}

