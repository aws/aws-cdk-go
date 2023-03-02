package awskinesisanalyticsv2


// For a SQL-based Kinesis Data Analytics application, describes the number of in-application streams to create for a given streaming source.
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
	// The number of in-application streams to create.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

