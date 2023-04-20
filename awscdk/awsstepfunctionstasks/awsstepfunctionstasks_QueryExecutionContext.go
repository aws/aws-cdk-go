package awsstepfunctionstasks


// Database and data catalog context in which the query execution occurs.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Athena Start Query"), &AthenaStartQueryExecutionProps{
//   	QueryString: sfn.JsonPath_Format(jsii.String("select contacts where year={};"), sfn.JsonPath_StringAt(jsii.String("$.year"))),
//   	QueryExecutionContext: &QueryExecutionContext{
//   		DatabaseName: jsii.String("interactions"),
//   	},
//   	ResultConfiguration: &ResultConfiguration{
//   		EncryptionConfiguration: &EncryptionConfiguration{
//   			EncryptionOption: tasks.EncryptionOption_S3_MANAGED,
//   		},
//   		OutputLocation: &Location{
//   			BucketName: jsii.String("mybucket"),
//   			ObjectKey: jsii.String("myprefix"),
//   		},
//   	},
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecutionContext.html
//
// Experimental.
type QueryExecutionContext struct {
	// Name of catalog used in query execution.
	// Experimental.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Name of database used in query execution.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

