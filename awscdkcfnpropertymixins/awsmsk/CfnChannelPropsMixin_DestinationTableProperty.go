package awsmsk


// Destination table configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   destinationTableProperty := &DestinationTableProperty{
//   	DestinationDatabaseName: jsii.String("destinationDatabaseName"),
//   	DestinationTableName: jsii.String("destinationTableName"),
//   	PartitionSpec: &PartitionSpecProperty{
//   		PartitionStrategy: jsii.String("partitionStrategy"),
//   		SourceList: []interface{}{
//   			&PartitionSourceProperty{
//   				SourceName: jsii.String("sourceName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-destinationtable.html
//
type CfnChannelPropsMixin_DestinationTableProperty struct {
	// The destination database name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-destinationtable.html#cfn-msk-channel-destinationtable-destinationdatabasename
	//
	DestinationDatabaseName *string `field:"optional" json:"destinationDatabaseName" yaml:"destinationDatabaseName"`
	// The destination table name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-destinationtable.html#cfn-msk-channel-destinationtable-destinationtablename
	//
	DestinationTableName *string `field:"optional" json:"destinationTableName" yaml:"destinationTableName"`
	// Partition specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-destinationtable.html#cfn-msk-channel-destinationtable-partitionspec
	//
	PartitionSpec interface{} `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
}

