package awsec2


// Describes a path component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisComponentProperty := &AnalysisComponentProperty{
//   	Arn: jsii.String("arn"),
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysiscomponent.html
//
type CfnNetworkInsightsAnalysis_AnalysisComponentProperty struct {
	// The Amazon Resource Name (ARN) of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysiscomponent.html#cfn-ec2-networkinsightsanalysis-analysiscomponent-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The ID of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysiscomponent.html#cfn-ec2-networkinsightsanalysis-analysiscomponent-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

