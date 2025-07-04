package awsquicksight


// A structure that describes the query execution options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryExecutionOptionsProperty := &QueryExecutionOptionsProperty{
//   	QueryExecutionMode: jsii.String("queryExecutionMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-queryexecutionoptions.html
//
type CfnAnalysis_QueryExecutionOptionsProperty struct {
	// A structure that describes the query execution mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-queryexecutionoptions.html#cfn-quicksight-analysis-queryexecutionoptions-queryexecutionmode
	//
	QueryExecutionMode *string `field:"optional" json:"queryExecutionMode" yaml:"queryExecutionMode"`
}

