package awspersonalize

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRecommenderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRecommenderMixinProps := &CfnRecommenderMixinProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	Name: jsii.String("name"),
//   	RecipeArn: jsii.String("recipeArn"),
//   	RecommenderConfig: &RecommenderConfigProperty{
//   		EnableMetadataWithRecommendations: jsii.Boolean(false),
//   		ItemExplorationConfig: map[string]*string{
//   			"itemExplorationConfigKey": jsii.String("itemExplorationConfig"),
//   		},
//   		MinRecommendationRequestsPerSecond: jsii.Number(123),
//   		TrainingDataConfig: &TrainingDataConfigProperty{
//   			ExcludedDatasetColumns: map[string][]*string{
//   				"excludedDatasetColumnsKey": []*string{
//   					jsii.String("excludedDatasetColumns"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-recommender.html
//
type CfnRecommenderMixinProps struct {
	// The ARN of the Domain dataset group that the recommender is created for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-recommender.html#cfn-personalize-recommender-datasetgrouparn
	//
	DatasetGroupArn *string `field:"optional" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The name of the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-recommender.html#cfn-personalize-recommender-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the recipe (Domain dataset group use case) that the recommender uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-recommender.html#cfn-personalize-recommender-recipearn
	//
	RecipeArn *string `field:"optional" json:"recipeArn" yaml:"recipeArn"`
	// The configuration details of the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-recommender.html#cfn-personalize-recommender-recommenderconfig
	//
	RecommenderConfig interface{} `field:"optional" json:"recommenderConfig" yaml:"recommenderConfig"`
	// Tags to associate with the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-recommender.html#cfn-personalize-recommender-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

