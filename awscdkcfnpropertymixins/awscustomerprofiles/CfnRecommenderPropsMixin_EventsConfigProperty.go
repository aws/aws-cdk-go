package awscustomerprofiles


// Configuration for events used in the recommender.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   eventsConfigProperty := &EventsConfigProperty{
//   	EventParametersList: []interface{}{
//   		&EventParametersProperty{
//   			EventType: jsii.String("eventType"),
//   			EventValueThreshold: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-eventsconfig.html
//
type CfnRecommenderPropsMixin_EventsConfigProperty struct {
	// List of event parameters with their value thresholds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-eventsconfig.html#cfn-customerprofiles-recommender-eventsconfig-eventparameterslist
	//
	EventParametersList interface{} `field:"optional" json:"eventParametersList" yaml:"eventParametersList"`
}

