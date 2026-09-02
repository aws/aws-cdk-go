package awsdataexchange


// What occurs to start an action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventProperty := &EventProperty{
//   	RevisionPublished: &RevisionPublishedProperty{
//   		DataSetId: jsii.String("dataSetId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-event.html
//
type CfnEventAction_EventProperty struct {
	// Information about the published revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-event.html#cfn-dataexchange-eventaction-event-revisionpublished
	//
	RevisionPublished interface{} `field:"optional" json:"revisionPublished" yaml:"revisionPublished"`
}

