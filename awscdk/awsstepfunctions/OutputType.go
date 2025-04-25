package awsstepfunctions


// The format of the Output of the child workflow executions.
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
type OutputType string

const (
	// Formats the results as a JSON array.
	OutputType_JSON OutputType = "JSON"
	// Formats the results as JSON Lines.
	OutputType_JSONL OutputType = "JSONL"
)

