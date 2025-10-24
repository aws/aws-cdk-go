package awslogs


// Use this processor to convert a value type associated with the specified key to the specified type.
//
// It's a casting processor that changes the types of the specified fields.
// For more information about this processor including examples, see typeConverter in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   typeConverterProperty := &TypeConverterProperty{
//   	Entries: []TypeConverterEntryProperty{
//   		&TypeConverterEntryProperty{
//   			Key: jsii.String("key"),
//   			Type: awscdk.Aws_logs.TypeConverterType_BOOLEAN,
//   		},
//   	},
//   }
//
type TypeConverterProperty struct {
	// An array of TypeConverterEntry objects, where each object contains information about one field to change the type of.
	Entries *[]*TypeConverterEntryProperty `field:"required" json:"entries" yaml:"entries"`
}

