// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTag := &cfnTag{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
// Experimental.
type CfnTag struct {
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

