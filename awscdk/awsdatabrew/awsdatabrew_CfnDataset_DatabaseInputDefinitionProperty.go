package awsdatabrew


// Connection information for dataset input files stored in a database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseInputDefinitionProperty := &databaseInputDefinitionProperty{
//   	glueConnectionName: jsii.String("glueConnectionName"),
//
//   	// the properties below are optional
//   	databaseTableName: jsii.String("databaseTableName"),
//   	queryString: jsii.String("queryString"),
//   	tempDirectory: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnDataset_DatabaseInputDefinitionProperty struct {
	// The AWS Glue Connection that stores the connection information for the target database.
	GlueConnectionName *string `field:"required" json:"glueConnectionName" yaml:"glueConnectionName"`
	// The table within the target database.
	DatabaseTableName *string `field:"optional" json:"databaseTableName" yaml:"databaseTableName"`
	// Custom SQL to run against the provided AWS Glue connection.
	//
	// This SQL will be used as the input for DataBrew projects and jobs.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// An Amazon location that AWS Glue Data Catalog can use as a temporary directory.
	TempDirectory interface{} `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

