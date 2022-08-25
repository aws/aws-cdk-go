package awsstepfunctionstasks


// Database and data catalog context in which the query execution occurs.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &athenaStartQueryExecutionProps{
//   	queryString: sfn.jsonPath.stringAt(jsii.String("$.queryString")),
//   	queryExecutionContext: &queryExecutionContext{
//   		databaseName: jsii.String("mydatabase"),
//   	},
//   	resultConfiguration: &resultConfiguration{
//   		encryptionConfiguration: &encryptionConfiguration{
//   			encryptionOption: tasks.encryptionOption_S3_MANAGED,
//   		},
//   		outputLocation: &location{
//   			bucketName: jsii.String("query-results-bucket"),
//   			objectKey: jsii.String("folder"),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecutionContext.html
//
type QueryExecutionContext struct {
	// Name of catalog used in query execution.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Name of database used in query execution.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

