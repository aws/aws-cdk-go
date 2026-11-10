package awspersonalize


// The configuration details of the recommender.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   recommenderConfigProperty := &RecommenderConfigProperty{
//   	EnableMetadataWithRecommendations: jsii.Boolean(false),
//   	ItemExplorationConfig: map[string]*string{
//   		"itemExplorationConfigKey": jsii.String("itemExplorationConfig"),
//   	},
//   	MinRecommendationRequestsPerSecond: jsii.Number(123),
//   	TrainingDataConfig: &TrainingDataConfigProperty{
//   		ExcludedDatasetColumns: map[string][]*string{
//   			"excludedDatasetColumnsKey": []*string{
//   				jsii.String("excludedDatasetColumns"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-recommender-recommenderconfig.html
//
type CfnRecommenderPropsMixin_RecommenderConfigProperty struct {
	// Whether metadata with recommendations is enabled for the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-recommender-recommenderconfig.html#cfn-personalize-recommender-recommenderconfig-enablemetadatawithrecommendations
	//
	EnableMetadataWithRecommendations interface{} `field:"optional" json:"enableMetadataWithRecommendations" yaml:"enableMetadataWithRecommendations"`
	// Specifies the exploration configuration hyperparameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-recommender-recommenderconfig.html#cfn-personalize-recommender-recommenderconfig-itemexplorationconfig
	//
	ItemExplorationConfig interface{} `field:"optional" json:"itemExplorationConfig" yaml:"itemExplorationConfig"`
	// Specifies the requested minimum provisioned recommendation requests per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-recommender-recommenderconfig.html#cfn-personalize-recommender-recommenderconfig-minrecommendationrequestspersecond
	//
	MinRecommendationRequestsPerSecond *float64 `field:"optional" json:"minRecommendationRequestsPerSecond" yaml:"minRecommendationRequestsPerSecond"`
	// The training data configuration to use when creating a domain recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-recommender-recommenderconfig.html#cfn-personalize-recommender-recommenderconfig-trainingdataconfig
	//
	TrainingDataConfig interface{} `field:"optional" json:"trainingDataConfig" yaml:"trainingDataConfig"`
}

