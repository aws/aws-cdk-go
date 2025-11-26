package previewawsgluemixins


// The database and table in the AWS Glue Data Catalog that is used for input or output data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueTablesProperty := &GlueTablesProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	ConnectionName: jsii.String("connectionName"),
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-gluetables.html
//
type CfnMLTransformPropsMixin_GlueTablesProperty struct {
	// A unique identifier for the AWS Glue Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-gluetables.html#cfn-glue-mltransform-gluetables-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the connection to the AWS Glue Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-gluetables.html#cfn-glue-mltransform-gluetables-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A database name in the AWS Glue Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-gluetables.html#cfn-glue-mltransform-gluetables-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// A table name in the AWS Glue Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-gluetables.html#cfn-glue-mltransform-gluetables-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

