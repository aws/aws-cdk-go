package awscassandra


// The settings for the CDC stream of a table.
//
// For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/cdc.html) in the *Amazon Keyspaces Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cdcSpecificationProperty := &CdcSpecificationProperty{
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	ViewType: jsii.String("viewType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-cdcspecification.html
//
type CfnTable_CdcSpecificationProperty struct {
	// The status of the CDC stream.
	//
	// You can enable or disable a stream for a table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-cdcspecification.html#cfn-cassandra-table-cdcspecification-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// The view type specifies the changes Amazon Keyspaces records for each changed row in the stream.
	//
	// After you create the stream, you can't make changes to this selection.
	//
	// The options are:
	//
	// - `NEW_AND_OLD_IMAGES` - both versions of the row, before and after the change. This is the default.
	// - `NEW_IMAGE` - the version of the row after the change.
	// - `OLD_IMAGE` - the version of the row before the change.
	// - `KEYS_ONLY` - the partition and clustering keys of the row that was changed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-cdcspecification.html#cfn-cassandra-table-cdcspecification-viewtype
	//
	// Default: - "NEW_AND_OLD_IMAGES".
	//
	ViewType *string `field:"optional" json:"viewType" yaml:"viewType"`
}

