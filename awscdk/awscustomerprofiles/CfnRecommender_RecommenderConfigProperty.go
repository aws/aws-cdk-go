package awscustomerprofiles


// Configuration for the recommender.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recommenderConfigProperty := &RecommenderConfigProperty{
//   	EventsConfig: &EventsConfigProperty{
//   		EventParametersList: []interface{}{
//   			&EventParametersProperty{
//   				EventType: jsii.String("eventType"),
//
//   				// the properties below are optional
//   				EventValueThreshold: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-recommenderconfig.html
//
type CfnRecommender_RecommenderConfigProperty struct {
	// Configuration for events used in the recommender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-recommenderconfig.html#cfn-customerprofiles-recommender-recommenderconfig-eventsconfig
	//
	EventsConfig interface{} `field:"optional" json:"eventsConfig" yaml:"eventsConfig"`
}

