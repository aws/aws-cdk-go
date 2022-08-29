package awscloudfront


// The type of events that a CloudFront function can be invoked in response to.
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
type FunctionEventType string

const (
	// The viewer-request specifies the incoming request.
	FunctionEventType_VIEWER_REQUEST FunctionEventType = "VIEWER_REQUEST"
	// The viewer-response specifies the outgoing response.
	FunctionEventType_VIEWER_RESPONSE FunctionEventType = "VIEWER_RESPONSE"
)

