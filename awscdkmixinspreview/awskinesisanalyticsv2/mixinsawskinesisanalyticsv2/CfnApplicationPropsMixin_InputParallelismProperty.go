package mixinsawskinesisanalyticsv2


// For a SQL-based Kinesis Data Analytics application, describes the number of in-application streams to create for a given streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputParallelismProperty := &InputParallelismProperty{
//   	Count: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputparallelism.html
//
type CfnApplicationPropsMixin_InputParallelismProperty struct {
	// The number of in-application streams to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputparallelism.html#cfn-kinesisanalyticsv2-application-inputparallelism-count
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

