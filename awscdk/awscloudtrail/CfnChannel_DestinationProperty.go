package awscloudtrail


// Contains information about the destination receiving events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &DestinationProperty{
//   	Location: jsii.String("location"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-channel-destination.html
//
type CfnChannel_DestinationProperty struct {
	// For channels used for a CloudTrail Lake integration, the location is the ARN of an event data store that receives events from a channel.
	//
	// For service-linked channels, the location is the name of the AWS service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-channel-destination.html#cfn-cloudtrail-channel-destination-location
	//
	Location *string `field:"required" json:"location" yaml:"location"`
	// The type of destination for events arriving from a channel.
	//
	// For channels used for a CloudTrail Lake integration, the value is `EVENT_DATA_STORE` . For service-linked channels, the value is `AWS_SERVICE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-channel-destination.html#cfn-cloudtrail-channel-destination-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

