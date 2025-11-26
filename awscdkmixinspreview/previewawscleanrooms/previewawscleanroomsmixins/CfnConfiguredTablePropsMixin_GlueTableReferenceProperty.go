package previewawscleanroomsmixins


// A reference to a table within an AWS Glue data catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueTableReferenceProperty := &GlueTableReferenceProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	Region: jsii.String("region"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-gluetablereference.html
//
type CfnConfiguredTablePropsMixin_GlueTableReferenceProperty struct {
	// The name of the database the AWS Glue table belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-gluetablereference.html#cfn-cleanrooms-configuredtable-gluetablereference-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The AWS Region where the AWS Glue table is located.
	//
	// This parameter is required to uniquely identify and access tables across different Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-gluetablereference.html#cfn-cleanrooms-configuredtable-gluetablereference-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The name of the AWS Glue table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-gluetablereference.html#cfn-cleanrooms-configuredtable-gluetablereference-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

