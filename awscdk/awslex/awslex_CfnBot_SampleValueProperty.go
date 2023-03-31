package awslex


// Defines one of the values for a slot type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleValueProperty := &sampleValueProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnBot_SampleValueProperty struct {
	// The value that can be used for a slot type.
	Value *string `field:"required" json:"value" yaml:"value"`
}

