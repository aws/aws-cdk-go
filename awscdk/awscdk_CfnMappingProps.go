// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   regionTable := awscdk.NewCfnMapping(this, jsii.String("RegionTable"), &cfnMappingProps{
//   	mapping: map[string]map[string]interface{}{
//   		"us-east-1": map[string]interface{}{
//   			"regionName": jsii.String("US East (N. Virginia)"),
//   		},
//   		"us-east-2": map[string]interface{}{
//   			"regionName": jsii.String("US East (Ohio)"),
//   		},
//   	},
//   })
//
//   regionTable.findInMap(awscdk.Aws_REGION(), jsii.String("regionName"))
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
	Mapping *map[string]*map[string]interface{} `field:"optional" json:"mapping" yaml:"mapping"`
}

