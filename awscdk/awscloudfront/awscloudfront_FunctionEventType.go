package awscloudfront


// The type of events that a CloudFront function can be invoked in response to.
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
type FunctionEventType string

const (
	// The viewer-request specifies the incoming request.
	// Experimental.
	FunctionEventType_VIEWER_REQUEST FunctionEventType = "VIEWER_REQUEST"
	// The viewer-response specifies the outgoing response.
	// Experimental.
	FunctionEventType_VIEWER_RESPONSE FunctionEventType = "VIEWER_RESPONSE"
)

