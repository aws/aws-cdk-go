// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mapping interface{}
//
//   cfnMappingProps := &cfnMappingProps{
//   	lazy: jsii.Boolean(false),
//   	mapping: map[string]map[string]interface{}{
//   		"mappingKey": map[string]interface{}{
//   			"mappingKey": mapping,
//   		},
//   	},
//   }
//
// Experimental.
type CfnMappingProps struct {
	// Experimental.
	Lazy *bool `field:"optional" json:"lazy" yaml:"lazy"`
	// Mapping of key to a set of corresponding set of named values.
	//
	// The key identifies a map of name-value pairs and must be unique within the mapping.
	//
	// For example, if you want to set values based on a region, you can create a mapping
	// that uses the region name as a key and contains the values you want to specify for
	// each specific region.
	// Experimental.
	Mapping *map[string]*map[string]interface{} `field:"optional" json:"mapping" yaml:"mapping"`
}

