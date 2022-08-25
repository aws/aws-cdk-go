package awscloudfront


// Properties for creating a CloudFront Function.
//
// Example:
//   var s3Bucket bucket
//   // Add a cloudfront Function to a Distribution
//   cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &functionProps{
//   	code: cloudfront.functionCode.fromInline(jsii.String("function handler(event) { return event.request }")),
//   })
//   cloudfront.NewDistribution(this, jsii.String("distro"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewS3Origin(s3Bucket),
//   		functionAssociations: []functionAssociation{
//   			&functionAssociation{
//   				function: cfFunction,
//   				eventType: cloudfront.functionEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
type FunctionProps struct {
	// The source code of the function.
	Code FunctionCode `field:"required" json:"code" yaml:"code"`
	// A comment to describe the function.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A name to identify the function.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
}

