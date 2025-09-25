package awslogs


// This object defines one value type that will be converted using the typeConverter processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   typeConverterEntryProperty := &TypeConverterEntryProperty{
//   	Key: jsii.String("key"),
//   	Type: awscdk.Aws_logs.TypeConverterType_BOOLEAN,
//   }
//
type TypeConverterEntryProperty struct {
	// The key with the value that is to be converted to a different type.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The data type to convert the field value to.
	Type TypeConverterType `field:"required" json:"type" yaml:"type"`
}

