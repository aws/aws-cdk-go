package awsstepfunctions


// The transformation to be applied to the Output of the Child Workflow executions.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // create a bucket
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//
//   // create a WriterConfig
//
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("Distributed Map State"), &DistributedMapProps{
//   	ResultWriterV2: sfn.NewResultWriterV2(&ResultWriterV2Props{
//   		Bucket: bucket,
//   		Prefix: jsii.String("my-prefix"),
//   		WriterConfig: map[string]interface{}{
//   			"outputType": sfn.OutputType_JSONL,
//   			"transformation": sfn.Transformation_NONE,
//   		},
//   	}),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")))
//
type Transformation string

const (
	// Returns the output of the child workflow executions unchanged, in addition to the workflow metadata.
	//
	// Default when exporting the child workflow execution results to Amazon S3 and WriterConfig is not specified.
	Transformation_NONE Transformation = "NONE"
	// Returns the output of the child workflow executions.
	//
	// Default when ResultWriter is not specified.
	Transformation_COMPACT Transformation = "COMPACT"
	// Returns the output of the child workflow executions.
	//
	// If a child workflow execution returns an array,this option flattens the array,
	// prior to returning the result to a state output or writing the result to an Amazon S3 object.
	Transformation_FLATTEN Transformation = "FLATTEN"
)

