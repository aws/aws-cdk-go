package awsec2


// Describes a path component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisComponentProperty := &analysisComponentProperty{
//   	arn: jsii.String("arn"),
//   	id: jsii.String("id"),
//   }
//
type CfnNetworkInsightsAnalysis_AnalysisComponentProperty struct {
	// The Amazon Resource Name (ARN) of the component.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The ID of the component.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

