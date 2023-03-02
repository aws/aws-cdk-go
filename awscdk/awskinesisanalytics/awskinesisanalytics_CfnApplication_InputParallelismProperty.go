package awskinesisanalytics


// Describes the number of in-application streams to create for a given streaming source.
//
// For information about parallelism, see [Configuring Application Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputParallelismProperty := &inputParallelismProperty{
//   	count: jsii.Number(123),
//   }
//
type CfnApplication_InputParallelismProperty struct {
	// Number of in-application streams to create.
	//
	// For more information, see [Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html) .
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

