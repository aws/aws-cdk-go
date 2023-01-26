// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var props interface{}
//
//   getContextKeyResult := &getContextKeyResult{
//   	key: jsii.String("key"),
//   	props: map[string]interface{}{
//   		"propsKey": props,
//   	},
//   }
//
type GetContextKeyResult struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Props *map[string]interface{} `field:"required" json:"props" yaml:"props"`
}

