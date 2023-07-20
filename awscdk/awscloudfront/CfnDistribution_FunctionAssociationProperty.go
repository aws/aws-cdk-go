package awscloudfront


// A CloudFront function that is associated with a cache behavior in a CloudFront distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionAssociationProperty := &FunctionAssociationProperty{
//   	EventType: jsii.String("eventType"),
//   	FunctionArn: jsii.String("functionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-functionassociation.html
//
type CfnDistribution_FunctionAssociationProperty struct {
	// The event type of the function, either `viewer-request` or `viewer-response` .
	//
	// You cannot use origin-facing event types ( `origin-request` and `origin-response` ) with a CloudFront function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-functionassociation.html#cfn-cloudfront-distribution-functionassociation-eventtype
	//
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// The Amazon Resource Name (ARN) of the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-functionassociation.html#cfn-cloudfront-distribution-functionassociation-functionarn
	//
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
}

