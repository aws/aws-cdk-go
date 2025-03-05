package awstimestream


// An attribute used in partitioning data in a table.
//
// A dimension key partitions data using the values of the dimension specified by the dimension-name as partition key, while a measure key partitions data using measure names (values of the 'measure_name' column).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partitionKeyProperty := &PartitionKeyProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	EnforcementInRecord: jsii.String("enforcementInRecord"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-partitionkey.html
//
type CfnTable_PartitionKeyProperty struct {
	// The type of the partition key.
	//
	// Options are DIMENSION (dimension key) and MEASURE (measure key).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-partitionkey.html#cfn-timestream-table-partitionkey-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The level of enforcement for the specification of a dimension key in ingested records.
	//
	// Options are REQUIRED (dimension key must be specified) and OPTIONAL (dimension key does not have to be specified).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-partitionkey.html#cfn-timestream-table-partitionkey-enforcementinrecord
	//
	EnforcementInRecord *string `field:"optional" json:"enforcementInRecord" yaml:"enforcementInRecord"`
	// The name of the attribute used for a dimension key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-partitionkey.html#cfn-timestream-table-partitionkey-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

