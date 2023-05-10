// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var props interface{}
//
//   getContextKeyResult := &GetContextKeyResult{
//   	Key: jsii.String("key"),
//   	Props: map[string]interface{}{
//   		"propsKey": props,
//   	},
//   }
//
// Experimental.
type GetContextKeyResult struct {
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Experimental.
	Props *map[string]interface{} `field:"required" json:"props" yaml:"props"`
}

