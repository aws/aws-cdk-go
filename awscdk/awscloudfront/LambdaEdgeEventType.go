package awscloudfront


// The type of events that a Lambda@Edge function can be invoked in response to.
//
// Example:
//   var myBucket bucket
//   // A Lambda@Edge function added to default behavior of a Distribution
//   // and triggered on every request
//   myFunc := experimental.NewEdgeFunction(this, jsii.String("MyFunction"), &EdgeFunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(myBucket),
//   		EdgeLambdas: []edgeLambda{
//   			&edgeLambda{
//   				FunctionVersion: myFunc.currentVersion,
//   				EventType: cloudfront.LambdaEdgeEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
type LambdaEdgeEventType string

const (
	// The origin-request specifies the request to the origin location (e.g. S3).
	LambdaEdgeEventType_ORIGIN_REQUEST LambdaEdgeEventType = "ORIGIN_REQUEST"
	// The origin-response specifies the response from the origin location (e.g. S3).
	LambdaEdgeEventType_ORIGIN_RESPONSE LambdaEdgeEventType = "ORIGIN_RESPONSE"
	// The viewer-request specifies the incoming request.
	LambdaEdgeEventType_VIEWER_REQUEST LambdaEdgeEventType = "VIEWER_REQUEST"
	// The viewer-response specifies the outgoing response.
	LambdaEdgeEventType_VIEWER_RESPONSE LambdaEdgeEventType = "VIEWER_RESPONSE"
)

