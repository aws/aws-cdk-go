package awscloudfront


// The type of events that a CloudFront function can be invoked in response to.
//
// Example:
//   var s3Bucket Bucket
//   // Add a cloudfront Function to a Distribution
//   cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
//   	Runtime: cloudfront.FunctionRuntime_JS_2_0(),
//   })
//   cloudfront.NewDistribution(this, jsii.String("distro"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(s3Bucket),
//   		FunctionAssociations: []FunctionAssociation{
//   			&FunctionAssociation{
//   				Function: cfFunction,
//   				EventType: cloudfront.FunctionEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
type FunctionEventType string

const (
	// The viewer-request specifies the incoming request.
	FunctionEventType_VIEWER_REQUEST FunctionEventType = "VIEWER_REQUEST"
	// The viewer-response specifies the outgoing response.
	FunctionEventType_VIEWER_RESPONSE FunctionEventType = "VIEWER_RESPONSE"
)

