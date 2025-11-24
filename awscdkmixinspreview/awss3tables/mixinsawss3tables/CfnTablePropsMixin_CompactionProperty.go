package mixinsawss3tables


// Contains details about the compaction settings for an Iceberg table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   compactionProperty := &CompactionProperty{
//   	Status: jsii.String("status"),
//   	TargetFileSizeMb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-compaction.html
//
type CfnTablePropsMixin_CompactionProperty struct {
	// The status of the maintenance configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-compaction.html#cfn-s3tables-table-compaction-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The target file size for the table in MB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-compaction.html#cfn-s3tables-table-compaction-targetfilesizemb
	//
	TargetFileSizeMb *float64 `field:"optional" json:"targetFileSizeMb" yaml:"targetFileSizeMb"`
}

