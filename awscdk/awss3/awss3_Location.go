package awss3


// An interface that represents the location of a specific object in an S3 Bucket.
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

