package awsmsk


// Destination table configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationTableProperty := &DestinationTableProperty{
//   	DestinationDatabaseName: jsii.String("destinationDatabaseName"),
//   	DestinationTableName: jsii.String("destinationTableName"),
//
//   	// the properties below are optional
//   	PartitionSpec: &PartitionSpecProperty{
//   		PartitionStrategy: jsii.String("partitionStrategy"),
//
//   		// the properties below are optional
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
type CfnChannel_DestinationTableProperty struct {
	// The destination database name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-destinationtable.html#cfn-msk-channel-destinationtable-destinationdatabasename
	//
	DestinationDatabaseName *string `field:"required" json:"destinationDatabaseName" yaml:"destinationDatabaseName"`
	// The destination table name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-destinationtable.html#cfn-msk-channel-destinationtable-destinationtablename
	//
	DestinationTableName *string `field:"required" json:"destinationTableName" yaml:"destinationTableName"`
	// Partition specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-destinationtable.html#cfn-msk-channel-destinationtable-partitionspec
	//
	PartitionSpec interface{} `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
}

