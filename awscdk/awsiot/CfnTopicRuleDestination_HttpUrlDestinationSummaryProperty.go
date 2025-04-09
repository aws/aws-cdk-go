package awsiot


// HTTP URL destination properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpUrlDestinationSummaryProperty := &HttpUrlDestinationSummaryProperty{
//   	ConfirmationUrl: jsii.String("confirmationUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-httpurldestinationsummary.html
//
type CfnTopicRuleDestination_HttpUrlDestinationSummaryProperty struct {
	// The URL used to confirm the HTTP topic rule destination URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-httpurldestinationsummary.html#cfn-iot-topicruledestination-httpurldestinationsummary-confirmationurl
	//
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
}

