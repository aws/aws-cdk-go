package awspersonalize


// The configuration details of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignConfigProperty := &CampaignConfigProperty{
//   	EnableMetadataWithRecommendations: jsii.Boolean(false),
//   	ItemExplorationConfig: map[string]*string{
//   		"itemExplorationConfigKey": jsii.String("itemExplorationConfig"),
//   	},
//   	RankingInfluence: map[string]*f64{
//   		"rankingInfluenceKey": jsii.Number(123),
//   	},
//   	SyncWithLatestSolutionVersion: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-campaign-campaignconfig.html
//
type CfnCampaign_CampaignConfigProperty struct {
	// Whether metadata with recommendations is enabled for the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-campaign-campaignconfig.html#cfn-personalize-campaign-campaignconfig-enablemetadatawithrecommendations
	//
	EnableMetadataWithRecommendations interface{} `field:"optional" json:"enableMetadataWithRecommendations" yaml:"enableMetadataWithRecommendations"`
	// Specifies the exploration configuration hyperparameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-campaign-campaignconfig.html#cfn-personalize-campaign-campaignconfig-itemexplorationconfig
	//
	ItemExplorationConfig interface{} `field:"optional" json:"itemExplorationConfig" yaml:"itemExplorationConfig"`
	// A map of ranking influence values for POPULARITY and FRESHNESS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-campaign-campaignconfig.html#cfn-personalize-campaign-campaignconfig-rankinginfluence
	//
	RankingInfluence interface{} `field:"optional" json:"rankingInfluence" yaml:"rankingInfluence"`
	// Whether the campaign automatically updates to use the latest solution version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-campaign-campaignconfig.html#cfn-personalize-campaign-campaignconfig-syncwithlatestsolutionversion
	//
	SyncWithLatestSolutionVersion interface{} `field:"optional" json:"syncWithLatestSolutionVersion" yaml:"syncWithLatestSolutionVersion"`
}

