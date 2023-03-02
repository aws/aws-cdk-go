package awsiotanalytics


// Information about the partition dimensions in a data store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datastorePartitionsProperty := &datastorePartitionsProperty{
//   	partitions: []interface{}{
//   		&datastorePartitionProperty{
//   			partition: &partitionProperty{
//   				attributeName: jsii.String("attributeName"),
//   			},
//   			timestampPartition: &timestampPartitionProperty{
//   				attributeName: jsii.String("attributeName"),
//
//   				// the properties below are optional
//   				timestampFormat: jsii.String("timestampFormat"),
//   			},
//   		},
//   	},
//   }
//
type CfnDatastore_DatastorePartitionsProperty struct {
	// A list of partition dimensions in a data store.
	Partitions interface{} `field:"optional" json:"partitions" yaml:"partitions"`
}

