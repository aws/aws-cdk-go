package mixinsawsiotanalytics


// A single dimension to partition a data store.
//
// The dimension must be an `AttributePartition` or a `TimestampPartition` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   partitionProperty := &PartitionProperty{
//   	AttributeName: jsii.String("attributeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-partition.html
//
type CfnDatastorePropsMixin_PartitionProperty struct {
	// The name of the attribute that defines a partition dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-partition.html#cfn-iotanalytics-datastore-partition-attributename
	//
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
}

