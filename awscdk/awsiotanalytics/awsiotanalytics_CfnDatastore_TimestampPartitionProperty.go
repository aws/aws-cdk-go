package awsiotanalytics


// A partition dimension defined by a timestamp attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestampPartitionProperty := &timestampPartitionProperty{
//   	attributeName: jsii.String("attributeName"),
//
//   	// the properties below are optional
//   	timestampFormat: jsii.String("timestampFormat"),
//   }
//
type CfnDatastore_TimestampPartitionProperty struct {
	// The attribute name of the partition defined by a timestamp.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// The timestamp format of a partition defined by a timestamp.
	//
	// The default format is seconds since epoch (January 1, 1970 at midnight UTC time).
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

