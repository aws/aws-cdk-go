package awsdatabrew


// Connection information for dataset input files stored in a database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseInputDefinitionProperty := &DatabaseInputDefinitionProperty{
//   	GlueConnectionName: jsii.String("glueConnectionName"),
//
//   	// the properties below are optional
//   	DatabaseTableName: jsii.String("databaseTableName"),
//   	QueryString: jsii.String("queryString"),
//   	TempDirectory: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		Key: jsii.String("key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-databaseinputdefinition.html
//
type CfnDataset_DatabaseInputDefinitionProperty struct {
	// The AWS Glue Connection that stores the connection information for the target database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-databaseinputdefinition.html#cfn-databrew-dataset-databaseinputdefinition-glueconnectionname
	//
	GlueConnectionName *string `field:"required" json:"glueConnectionName" yaml:"glueConnectionName"`
	// The table within the target database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-databaseinputdefinition.html#cfn-databrew-dataset-databaseinputdefinition-databasetablename
	//
	DatabaseTableName *string `field:"optional" json:"databaseTableName" yaml:"databaseTableName"`
	// Custom SQL to run against the provided AWS Glue connection.
	//
	// This SQL will be used as the input for DataBrew projects and jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-databaseinputdefinition.html#cfn-databrew-dataset-databaseinputdefinition-querystring
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// An Amazon location that AWS Glue Data Catalog can use as a temporary directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-databaseinputdefinition.html#cfn-databrew-dataset-databaseinputdefinition-tempdirectory
	//
	TempDirectory interface{} `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

