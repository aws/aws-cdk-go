package awseks


// A key-value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelProperty := &labelProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnFargateProfile_LabelProperty struct {
	// Enter a key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Enter a value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

