package awsemr


// `KeyValue` is a subproperty of the `HadoopJarStepConfig` property type.
//
// `KeyValue` is used to pass parameters to a step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValueProperty := &keyValueProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnStep_KeyValueProperty struct {
	// The unique identifier of a key-value pair.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value part of the identified key.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

