// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   cfnJsonProps := &cfnJsonProps{
//   	value: value,
//   }
//
// Experimental.
type CfnJsonProps struct {
	// The value to resolve.
	//
	// Can be any JavaScript object, including tokens and
	// references in keys or values.
	// Experimental.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

