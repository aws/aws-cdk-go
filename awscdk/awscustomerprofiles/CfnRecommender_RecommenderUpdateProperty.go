package awscustomerprofiles


// Information about the latest recommender update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recommenderUpdateProperty := &RecommenderUpdateProperty{
//   	CreationDateTime: jsii.String("creationDateTime"),
//   	FailureReason: jsii.String("failureReason"),
//   	LastUpdatedDateTime: jsii.String("lastUpdatedDateTime"),
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
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-recommenderupdate.html
//
type CfnRecommender_RecommenderUpdateProperty struct {
	// The timestamp of when the update was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-recommenderupdate.html#cfn-customerprofiles-recommender-recommenderupdate-creationdatetime
	//
	CreationDateTime *string `field:"optional" json:"creationDateTime" yaml:"creationDateTime"`
	// The reason for update failure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-recommenderupdate.html#cfn-customerprofiles-recommender-recommenderupdate-failurereason
	//
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// The timestamp of when the update was last modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-recommenderupdate.html#cfn-customerprofiles-recommender-recommenderupdate-lastupdateddatetime
	//
	LastUpdatedDateTime *string `field:"optional" json:"lastUpdatedDateTime" yaml:"lastUpdatedDateTime"`
	// Configuration for the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-recommenderupdate.html#cfn-customerprofiles-recommender-recommenderupdate-recommenderconfig
	//
	RecommenderConfig interface{} `field:"optional" json:"recommenderConfig" yaml:"recommenderConfig"`
	// The status of the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-recommenderupdate.html#cfn-customerprofiles-recommender-recommenderupdate-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

