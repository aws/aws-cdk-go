package awscloudfront


// Properties for creating a CloudFront Function.
//
// Example:
//   var s3Bucket bucket
//   // Add a cloudfront Function to a Distribution
//   cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
//   })
//   cloudfront.NewDistribution(this, jsii.String("distro"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(s3Bucket),
//   		FunctionAssociations: []functionAssociation{
//   			&functionAssociation{
//   				Function: cfFunction,
//   				EventType: cloudfront.FunctionEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type FunctionProps struct {
	// The source code of the function.
	// Experimental.
	Code FunctionCode `field:"required" json:"code" yaml:"code"`
	// A comment to describe the function.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A name to identify the function.
	// Experimental.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
}

