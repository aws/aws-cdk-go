// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Properties of a Partition Index.
//
// Example:
//   var myTable table
//
//   myTable.addPartitionIndex(&partitionIndex{
//   	indexName: jsii.String("my-index"),
//   	keyNames: []*string{
//   		jsii.String("year"),
//   	},
//   })
//
// Experimental.
type PartitionIndex struct {
	// The partition key names that comprise the partition index.
	//
	// The names must correspond to a name in the
	// table's partition keys.
	// Experimental.
	KeyNames *[]*string `field:"required" json:"keyNames" yaml:"keyNames"`
	// The name of the partition index.
	// Experimental.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
}

