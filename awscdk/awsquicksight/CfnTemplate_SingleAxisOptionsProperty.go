package awsquicksight


// The settings of a chart's single axis configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleAxisOptionsProperty := &SingleAxisOptionsProperty{
//   	YAxisOptions: &YAxisOptionsProperty{
//   		YAxis: jsii.String("yAxis"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-singleaxisoptions.html
//
type CfnTemplate_SingleAxisOptionsProperty struct {
	// The Y axis options of a single axis configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-singleaxisoptions.html#cfn-quicksight-template-singleaxisoptions-yaxisoptions
	//
	YAxisOptions interface{} `field:"optional" json:"yAxisOptions" yaml:"yAxisOptions"`
}

