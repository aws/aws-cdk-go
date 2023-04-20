package awsiotanalytics


// A single dimension to partition a data store.
//
// The dimension must be an `AttributePartition` or a `TimestampPartition` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datastorePartitionProperty := &DatastorePartitionProperty{
//   	Partition: &PartitionProperty{
//   		AttributeName: jsii.String("attributeName"),
//   	},
//   	TimestampPartition: &TimestampPartitionProperty{
//   		AttributeName: jsii.String("attributeName"),
//
//   		// the properties below are optional
//   		TimestampFormat: jsii.String("timestampFormat"),
//   	},
//   }
//
type CfnDatastore_DatastorePartitionProperty struct {
	// A partition dimension defined by an attribute.
	Partition interface{} `field:"optional" json:"partition" yaml:"partition"`
	// A partition dimension defined by a timestamp attribute.
	TimestampPartition interface{} `field:"optional" json:"timestampPartition" yaml:"timestampPartition"`
}

