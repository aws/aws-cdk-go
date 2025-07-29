package awskinesisfirehose


// Describes the configuration of a destination in Apache Iceberg Tables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationTableConfigurationProperty := &DestinationTableConfigurationProperty{
//   	DestinationDatabaseName: jsii.String("destinationDatabaseName"),
//   	DestinationTableName: jsii.String("destinationTableName"),
//
//   	// the properties below are optional
//   	PartitionSpec: &PartitionSpecProperty{
//   		Identity: []interface{}{
//   			&PartitionFieldProperty{
//   				SourceName: jsii.String("sourceName"),
//   			},
//   		},
//   	},
//   	S3ErrorOutputPrefix: jsii.String("s3ErrorOutputPrefix"),
//   	UniqueKeys: []*string{
//   		jsii.String("uniqueKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html
//
type CfnDeliveryStream_DestinationTableConfigurationProperty struct {
	// The name of the Apache Iceberg database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html#cfn-kinesisfirehose-deliverystream-destinationtableconfiguration-destinationdatabasename
	//
	DestinationDatabaseName *string `field:"required" json:"destinationDatabaseName" yaml:"destinationDatabaseName"`
	// Specifies the name of the Apache Iceberg Table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html#cfn-kinesisfirehose-deliverystream-destinationtableconfiguration-destinationtablename
	//
	DestinationTableName *string `field:"required" json:"destinationTableName" yaml:"destinationTableName"`
	// The partition spec configuration for a table that is used by automatic table creation.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html#cfn-kinesisfirehose-deliverystream-destinationtableconfiguration-partitionspec
	//
	PartitionSpec interface{} `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// The table specific S3 error output prefix.
	//
	// All the errors that occurred while delivering to this table will be prefixed with this value in S3 destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html#cfn-kinesisfirehose-deliverystream-destinationtableconfiguration-s3erroroutputprefix
	//
	S3ErrorOutputPrefix *string `field:"optional" json:"s3ErrorOutputPrefix" yaml:"s3ErrorOutputPrefix"`
	// A list of unique keys for a given Apache Iceberg table.
	//
	// Firehose will use these for running Create, Update, or Delete operations on the given Iceberg table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.html#cfn-kinesisfirehose-deliverystream-destinationtableconfiguration-uniquekeys
	//
	UniqueKeys *[]*string `field:"optional" json:"uniqueKeys" yaml:"uniqueKeys"`
}

