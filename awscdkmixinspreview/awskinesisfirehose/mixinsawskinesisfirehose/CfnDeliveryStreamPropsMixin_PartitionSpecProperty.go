package mixinsawskinesisfirehose


// Represents how to produce partition data for a table.
//
// Partition data is produced by transforming columns in a table. Each column transform is represented by a named `PartitionField` .
//
// Here is an example of the schema in JSON.
//
// `"partitionSpec": { "identity": [ {"sourceName": "column1"}, {"sourceName": "column2"}, {"sourceName": "column3"} ] }`
//
// Amazon Data Firehose is in preview release and is subject to change.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   partitionSpecProperty := &PartitionSpecProperty{
//   	Identity: []interface{}{
//   		&PartitionFieldProperty{
//   			SourceName: jsii.String("sourceName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-partitionspec.html
//
type CfnDeliveryStreamPropsMixin_PartitionSpecProperty struct {
	// List of identity [transforms](https://docs.aws.amazon.com/https://iceberg.apache.org/spec/#partition-transforms) that performs an identity transformation. The transform takes the source value, and does not modify it. Result type is the source type.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-partitionspec.html#cfn-kinesisfirehose-deliverystream-partitionspec-identity
	//
	Identity interface{} `field:"optional" json:"identity" yaml:"identity"`
}

