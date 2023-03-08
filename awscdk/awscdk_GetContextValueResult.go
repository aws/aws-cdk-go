// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   getContextValueResult := &getContextValueResult{
//   	value: value,
//   }
//
// Experimental.
type GetContextValueResult struct {
	// Experimental.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

