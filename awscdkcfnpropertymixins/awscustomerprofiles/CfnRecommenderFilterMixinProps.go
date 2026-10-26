package awscustomerprofiles


// Properties for CfnRecommenderFilterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRecommenderFilterMixinProps := &CfnRecommenderFilterMixinProps{
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
//   	RecommenderFilterExpression: jsii.String("recommenderFilterExpression"),
//   	RecommenderFilterName: jsii.String("recommenderFilterName"),
//   	RecommenderSchemaName: jsii.String("recommenderSchemaName"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderfilter.html
//
type CfnRecommenderFilterMixinProps struct {
	// The description of the recommender filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderfilter.html#cfn-customerprofiles-recommenderfilter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderfilter.html#cfn-customerprofiles-recommenderfilter-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The filter expression that defines which items to include or exclude from recommendations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderfilter.html#cfn-customerprofiles-recommenderfilter-recommenderfilterexpression
	//
	RecommenderFilterExpression *string `field:"optional" json:"recommenderFilterExpression" yaml:"recommenderFilterExpression"`
	// The name of the recommender filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderfilter.html#cfn-customerprofiles-recommenderfilter-recommenderfiltername
	//
	RecommenderFilterName *string `field:"optional" json:"recommenderFilterName" yaml:"recommenderFilterName"`
	// The name of the recommender schema to use for this recommender filter.
	//
	// If not specified, the default schema is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderfilter.html#cfn-customerprofiles-recommenderfilter-recommenderschemaname
	//
	RecommenderSchemaName *string `field:"optional" json:"recommenderSchemaName" yaml:"recommenderSchemaName"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderfilter.html#cfn-customerprofiles-recommenderfilter-tags
	//
	Tags *[]*CfnRecommenderFilterPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

