package awscloudfront


// A CloudFront function that is associated with a cache behavior in a CloudFront distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionAssociationProperty := &functionAssociationProperty{
//   	eventType: jsii.String("eventType"),
//   	functionArn: jsii.String("functionArn"),
//   }
//
type CfnDistribution_FunctionAssociationProperty struct {
	// The event type of the function, either `viewer-request` or `viewer-response` .
	//
	// You cannot use origin-facing event types ( `origin-request` and `origin-response` ) with a CloudFront function.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
}

