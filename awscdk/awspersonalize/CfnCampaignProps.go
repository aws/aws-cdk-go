package awspersonalize

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCampaign`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCampaignProps := &CfnCampaignProps{
//   	Name: jsii.String("name"),
//   	SolutionVersionArn: jsii.String("solutionVersionArn"),
//
//   	// the properties below are optional
//   	CampaignConfig: &CampaignConfigProperty{
//   		EnableMetadataWithRecommendations: jsii.Boolean(false),
//   		ItemExplorationConfig: map[string]*string{
//   			"itemExplorationConfigKey": jsii.String("itemExplorationConfig"),
//   		},
//   		RankingInfluence: map[string]*f64{
//   			"rankingInfluenceKey": jsii.Number(123),
//   		},
//   		SyncWithLatestSolutionVersion: jsii.Boolean(false),
//   	},
//   	MinProvisionedTps: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-campaign.html
//
type CfnCampaignProps struct {
	// The name of the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-campaign.html#cfn-personalize-campaign-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the solution version to deploy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-campaign.html#cfn-personalize-campaign-solutionversionarn
	//
	SolutionVersionArn *string `field:"required" json:"solutionVersionArn" yaml:"solutionVersionArn"`
	// The configuration details of a campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-campaign.html#cfn-personalize-campaign-campaignconfig
	//
	CampaignConfig interface{} `field:"optional" json:"campaignConfig" yaml:"campaignConfig"`
	// Specifies the requested minimum provisioned transactions per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-campaign.html#cfn-personalize-campaign-minprovisionedtps
	//
	MinProvisionedTps *float64 `field:"optional" json:"minProvisionedTps" yaml:"minProvisionedTps"`
	// Tags to associate with the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-campaign.html#cfn-personalize-campaign-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

