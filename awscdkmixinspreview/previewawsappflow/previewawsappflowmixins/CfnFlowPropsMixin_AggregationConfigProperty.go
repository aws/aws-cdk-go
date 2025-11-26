package previewawsappflowmixins


// The aggregation settings that you can use to customize the output format of your flow data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aggregationConfigProperty := &AggregationConfigProperty{
//   	AggregationType: jsii.String("aggregationType"),
//   	TargetFileSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-aggregationconfig.html
//
type CfnFlowPropsMixin_AggregationConfigProperty struct {
	// Specifies whether Amazon AppFlow aggregates the flow records into a single file, or leave them unaggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-aggregationconfig.html#cfn-appflow-flow-aggregationconfig-aggregationtype
	//
	AggregationType *string `field:"optional" json:"aggregationType" yaml:"aggregationType"`
	// The desired file size, in MB, for each output file that Amazon AppFlow writes to the flow destination.
	//
	// For each file, Amazon AppFlow attempts to achieve the size that you specify. The actual file sizes might differ from this target based on the number and size of the records that each file contains.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-aggregationconfig.html#cfn-appflow-flow-aggregationconfig-targetfilesize
	//
	TargetFileSize *float64 `field:"optional" json:"targetFileSize" yaml:"targetFileSize"`
}

