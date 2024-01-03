package awscdk


// Example:
//   regionTable := awscdk.NewCfnMapping(this, jsii.String("RegionTable"), &CfnMappingProps{
//   	Mapping: map[string]map[string]interface{}{
//   		"us-east-1": map[string]interface{}{
//   			"regionName": jsii.String("US East (N. Virginia)"),
//   		},
//   		"us-east-2": map[string]interface{}{
//   			"regionName": jsii.String("US East (Ohio)"),
//   		},
//   	},
//   	Lazy: jsii.Boolean(true),
//   })
//
//   regionTable.FindInMap(jsii.String("us-east-2"), jsii.String("regionName"))
//
type CfnMappingProps struct {
	Lazy *bool `field:"optional" json:"lazy" yaml:"lazy"`
	// Mapping of key to a set of corresponding set of named values.
	//
	// The key identifies a map of name-value pairs and must be unique within the mapping.
	//
	// For example, if you want to set values based on a region, you can create a mapping
	// that uses the region name as a key and contains the values you want to specify for
	// each specific region.
	// Default: - No mapping.
	//
	Mapping *map[string]*map[string]interface{} `field:"optional" json:"mapping" yaml:"mapping"`
}

