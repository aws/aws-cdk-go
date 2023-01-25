package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNetworkInsightsAnalysis`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInsightsAnalysisProps := &cfnNetworkInsightsAnalysisProps{
//   	networkInsightsPathId: jsii.String("networkInsightsPathId"),
//
//   	// the properties below are optional
//   	additionalAccounts: []*string{
//   		jsii.String("additionalAccounts"),
//   	},
//   	filterInArns: []*string{
//   		jsii.String("filterInArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnNetworkInsightsAnalysisProps struct {
	// The ID of the path.
	NetworkInsightsPathId *string `field:"required" json:"networkInsightsPathId" yaml:"networkInsightsPathId"`
	// The member accounts that contain resources that the path can traverse.
	AdditionalAccounts *[]*string `field:"optional" json:"additionalAccounts" yaml:"additionalAccounts"`
	// The Amazon Resource Names (ARN) of the resources that the path must traverse.
	FilterInArns *[]*string `field:"optional" json:"filterInArns" yaml:"filterInArns"`
	// The tags to apply.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

