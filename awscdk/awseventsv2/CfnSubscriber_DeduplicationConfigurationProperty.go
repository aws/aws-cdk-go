package awseventsv2


// Deduplication settings applied to the forwarded events on the downstream bus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deduplicationConfigurationProperty := &DeduplicationConfigurationProperty{
//   	DeduplicationType: jsii.String("deduplicationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-deduplicationconfiguration.html
//
type CfnSubscriber_DeduplicationConfigurationProperty struct {
	// How duplicate events are detected: CONTENT_BASED deduplicates by a hash of the event content.
	//
	// To deduplicate by a caller-supplied token instead, omit DeduplicationConfiguration and set SystemMetadata.DeduplicationId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-deduplicationconfiguration.html#cfn-eventsv2-subscriber-deduplicationconfiguration-deduplicationtype
	//
	DeduplicationType *string `field:"required" json:"deduplicationType" yaml:"deduplicationType"`
}

