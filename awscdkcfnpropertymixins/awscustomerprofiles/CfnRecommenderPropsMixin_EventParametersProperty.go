package awscustomerprofiles


// Event parameters with type and value threshold.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   eventParametersProperty := &EventParametersProperty{
//   	EventType: jsii.String("eventType"),
//   	EventValueThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-eventparameters.html
//
type CfnRecommenderPropsMixin_EventParametersProperty struct {
	// The type of event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-eventparameters.html#cfn-customerprofiles-recommender-eventparameters-eventtype
	//
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// The threshold of the event type.
	//
	// Only events with a value greater or equal to this threshold will be considered for solution creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-eventparameters.html#cfn-customerprofiles-recommender-eventparameters-eventvaluethreshold
	//
	EventValueThreshold *float64 `field:"optional" json:"eventValueThreshold" yaml:"eventValueThreshold"`
}

