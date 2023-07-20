package awsglue


// A list of AWS Glue table definitions used by the transform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputRecordTablesProperty := &InputRecordTablesProperty{
//   	GlueTables: []interface{}{
//   		&GlueTablesProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			CatalogId: jsii.String("catalogId"),
//   			ConnectionName: jsii.String("connectionName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-inputrecordtables.html
//
type CfnMLTransform_InputRecordTablesProperty struct {
	// The database and table in the AWS Glue Data Catalog that is used for input or output data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-inputrecordtables.html#cfn-glue-mltransform-inputrecordtables-gluetables
	//
	GlueTables interface{} `field:"optional" json:"glueTables" yaml:"glueTables"`
}

