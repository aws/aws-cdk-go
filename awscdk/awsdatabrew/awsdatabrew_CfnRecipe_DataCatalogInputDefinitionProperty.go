package awsdatabrew


// Represents how metadata stored in the AWS Glue Data Catalog is defined in a DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogInputDefinitionProperty := &dataCatalogInputDefinitionProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	tableName: jsii.String("tableName"),
//   	tempDirectory: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnRecipe_DataCatalogInputDefinitionProperty struct {
	// The unique identifier of the AWS account that holds the Data Catalog that stores the data.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of a database in the Data Catalog.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of a database table in the Data Catalog.
	//
	// This table corresponds to a DataBrew dataset.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Represents an Amazon location where DataBrew can store intermediate results.
	TempDirectory interface{} `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

