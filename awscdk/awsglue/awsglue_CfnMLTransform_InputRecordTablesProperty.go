package awsglue


// A list of AWS Glue table definitions used by the transform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputRecordTablesProperty := &inputRecordTablesProperty{
//   	glueTables: []interface{}{
//   		&glueTablesProperty{
//   			databaseName: jsii.String("databaseName"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			catalogId: jsii.String("catalogId"),
//   			connectionName: jsii.String("connectionName"),
//   		},
//   	},
//   }
//
type CfnMLTransform_InputRecordTablesProperty struct {
	// The database and table in the AWS Glue Data Catalog that is used for input or output data.
	GlueTables interface{} `field:"optional" json:"glueTables" yaml:"glueTables"`
}

