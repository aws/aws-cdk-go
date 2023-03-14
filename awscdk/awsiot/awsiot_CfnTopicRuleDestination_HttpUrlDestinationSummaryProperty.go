package awsiot


// HTTP URL destination properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpUrlDestinationSummaryProperty := &httpUrlDestinationSummaryProperty{
//   	confirmationUrl: jsii.String("confirmationUrl"),
//   }
//
type CfnTopicRuleDestination_HttpUrlDestinationSummaryProperty struct {
	// The URL used to confirm the HTTP topic rule destination URL.
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
}

