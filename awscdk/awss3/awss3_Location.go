package awss3


// An interface that represents the location of a specific object in an S3 Bucket.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Athena Start Query"), &athenaStartQueryExecutionProps{
//   	queryString: sfn.jsonPath.format(jsii.String("select contacts where year={};"), sfn.*jsonPath.stringAt(jsii.String("$.year"))),
//   	queryExecutionContext: &queryExecutionContext{
//   		databaseName: jsii.String("interactions"),
//   	},
//   	resultConfiguration: &resultConfiguration{
//   		encryptionConfiguration: &encryptionConfiguration{
//   			encryptionOption: tasks.encryptionOption_S3_MANAGED,
//   		},
//   		outputLocation: &location{
//   			bucketName: jsii.String("mybucket"),
//   			objectKey: jsii.String("myprefix"),
//   		},
//   	},
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   })
//
// Experimental.
type Location struct {
	// The name of the S3 Bucket the object is in.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The path inside the Bucket where the object is located at.
	// Experimental.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// The S3 object version.
	// Experimental.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

