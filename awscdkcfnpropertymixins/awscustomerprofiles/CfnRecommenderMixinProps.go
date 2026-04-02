package awscustomerprofiles

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
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
//   	RecommenderConfig: &RecommenderConfigProperty{
//   		EventsConfig: &EventsConfigProperty{
//   			EventParametersList: []interface{}{
//   				&EventParametersProperty{
//   					EventType: jsii.String("eventType"),
//   					EventValueThreshold: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	RecommenderName: jsii.String("recommenderName"),
//   	RecommenderRecipeName: jsii.String("recommenderRecipeName"),
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
type CfnRecommenderMixinProps struct {
	// The description of the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the domain for which the recommender will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Configuration for the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-recommenderconfig
	//
	RecommenderConfig interface{} `field:"optional" json:"recommenderConfig" yaml:"recommenderConfig"`
	// The name of the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-recommendername
	//
	RecommenderName *string `field:"optional" json:"recommenderName" yaml:"recommenderName"`
	// The name of the recommender recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-recommenderrecipename
	//
	RecommenderRecipeName *string `field:"optional" json:"recommenderRecipeName" yaml:"recommenderRecipeName"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommender.html#cfn-customerprofiles-recommender-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

