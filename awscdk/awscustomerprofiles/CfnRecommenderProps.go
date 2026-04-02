package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRecommender`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecommenderProps := &CfnRecommenderProps{
//   	DomainName: jsii.String("domainName"),
//   	RecommenderName: jsii.String("recommenderName"),
//   	RecommenderRecipeName: jsii.String("recommenderRecipeName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	RecommenderConfig: &RecommenderConfigProperty{
//   		EventsConfig: &EventsConfigProperty{
//   			EventParametersList: []interface{}{
//   				&EventParametersProperty{
//   					EventType: jsii.String("eventType"),
//
//   					// the properties below are optional
//   					EventValueThreshold: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html
//
type CfnRecommenderProps struct {
	// The name of the domain for which the recommender will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-recommendername
	//
	RecommenderName *string `field:"required" json:"recommenderName" yaml:"recommenderName"`
	// The name of the recommender recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-recommenderrecipename
	//
	RecommenderRecipeName *string `field:"required" json:"recommenderRecipeName" yaml:"recommenderRecipeName"`
	// The description of the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration for the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-recommenderconfig
	//
	RecommenderConfig interface{} `field:"optional" json:"recommenderConfig" yaml:"recommenderConfig"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

