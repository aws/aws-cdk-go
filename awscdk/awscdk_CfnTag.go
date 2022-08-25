// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTag := &cfnTag{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnTag struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Value *string `field:"required" json:"value" yaml:"value"`
}

