package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   additionalDetailProperty := &additionalDetailProperty{
//   	additionalDetailType: jsii.String("additionalDetailType"),
//   	component: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   }
//
type CfnNetworkInsightsAnalysis_AdditionalDetailProperty struct {
	// `CfnNetworkInsightsAnalysis.AdditionalDetailProperty.AdditionalDetailType`.
	AdditionalDetailType *string `field:"optional" json:"additionalDetailType" yaml:"additionalDetailType"`
	// `CfnNetworkInsightsAnalysis.AdditionalDetailProperty.Component`.
	Component interface{} `field:"optional" json:"component" yaml:"component"`
}

