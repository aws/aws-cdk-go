package previewawscleanroomsmixins


// A reference to a table within Athena.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   athenaTableReferenceProperty := &AthenaTableReferenceProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	OutputLocation: jsii.String("outputLocation"),
//   	Region: jsii.String("region"),
//   	TableName: jsii.String("tableName"),
//   	WorkGroup: jsii.String("workGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html
//
type CfnConfiguredTablePropsMixin_AthenaTableReferenceProperty struct {
	// The database name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html#cfn-cleanrooms-configuredtable-athenatablereference-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The output location for the Athena table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html#cfn-cleanrooms-configuredtable-athenatablereference-outputlocation
	//
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// The AWS Region where the Athena table is located.
	//
	// This parameter is required to uniquely identify and access tables across different Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html#cfn-cleanrooms-configuredtable-athenatablereference-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html#cfn-cleanrooms-configuredtable-athenatablereference-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// The workgroup of the Athena table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-athenatablereference.html#cfn-cleanrooms-configuredtable-athenatablereference-workgroup
	//
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

