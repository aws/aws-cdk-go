package awscloudfront


// A complex type that contains a Lambda@Edge function association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaFunctionAssociationProperty := &lambdaFunctionAssociationProperty{
//   	eventType: jsii.String("eventType"),
//   	includeBody: jsii.Boolean(false),
//   	lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   }
//
type CfnDistribution_LambdaFunctionAssociationProperty struct {
	// Specifies the event type that triggers a Lambda@Edge function invocation. You can specify the following values:.
	//
	// - `viewer-request` : The function executes when CloudFront receives a request from a viewer and before it checks to see whether the requested object is in the edge cache.
	// - `origin-request` : The function executes only when CloudFront sends a request to your origin. When the requested object is in the edge cache, the function doesn't execute.
	// - `origin-response` : The function executes after CloudFront receives a response from the origin and before it caches the object in the response. When the requested object is in the edge cache, the function doesn't execute.
	// - `viewer-response` : The function executes before CloudFront returns the requested object to the viewer. The function executes regardless of whether the object was already in the edge cache.
	//
	// If the origin returns an HTTP status code other than HTTP 200 (OK), the function doesn't execute.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// A flag that allows a Lambda@Edge function to have read access to the body content.
	//
	// For more information, see [Accessing the Request Body by Choosing the Include Body Option](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the Amazon CloudFront Developer Guide.
	IncludeBody interface{} `field:"optional" json:"includeBody" yaml:"includeBody"`
	// The ARN of the Lambda@Edge function.
	//
	// You must specify the ARN of a function version; you can't specify an alias or $LATEST.
	LambdaFunctionArn *string `field:"optional" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}

