package awseventsv2


// The retry policy for failed deliveries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   retryPolicyProperty := &RetryPolicyProperty{
//   	MaxEventAgeInSeconds: jsii.Number(123),
//   	MaxRetryAttempts: jsii.Number(123),
//   	RetryStrategy: jsii.String("retryStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-retrypolicy.html
//
type CfnSubscriberPropsMixin_RetryPolicyProperty struct {
	// The maximum age of an event in seconds, 60-86400 (24 hours).
	//
	// When an event reaches this age, retries stop; if OnFailureConfiguration is set, the event is delivered to that destination, otherwise it is dropped. The default is 300.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-retrypolicy.html#cfn-eventsv2-subscriber-retrypolicy-maxeventageinseconds
	//
	MaxEventAgeInSeconds *float64 `field:"optional" json:"maxEventAgeInSeconds" yaml:"maxEventAgeInSeconds"`
	// The maximum number of retry attempts, 0-185.
	//
	// When the attempts are exhausted, retries stop; if OnFailureConfiguration is set, the event is delivered to that destination, otherwise it is dropped. The default is 5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-retrypolicy.html#cfn-eventsv2-subscriber-retrypolicy-maxretryattempts
	//
	MaxRetryAttempts *float64 `field:"optional" json:"maxRetryAttempts" yaml:"maxRetryAttempts"`
	// Which errors are retried.
	//
	// ALL retries all errors. The default is ALL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-retrypolicy.html#cfn-eventsv2-subscriber-retrypolicy-retrystrategy
	//
	RetryStrategy *string `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}

