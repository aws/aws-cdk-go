package awscleanrooms


// A reference to a table within Athena.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   athenaTableReferenceProperty := &AthenaTableReferenceProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   	WorkGroup: jsii.String("workGroup"),
//
//   	// the properties below are optional
//   	OutputLocation: jsii.String("outputLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html
//
type CfnConfiguredTable_AthenaTableReferenceProperty struct {
	// The database name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html#cfn-cleanrooms-configuredtable-athenatablereference-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html#cfn-cleanrooms-configuredtable-athenatablereference-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The workgroup of the Athena table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html#cfn-cleanrooms-configuredtable-athenatablereference-workgroup
	//
	WorkGroup *string `field:"required" json:"workGroup" yaml:"workGroup"`
	// The output location for the Athena table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html#cfn-cleanrooms-configuredtable-athenatablereference-outputlocation
	//
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

