package awslambda


// Authentication type for WebFunction domain configuration.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var codeBucket IBucket
//
//
//   lambda.NewWebFunction(this, jsii.String("PerRegionApi"), &WebFunctionProps{
//   	Code: lambda.WebFunctionCode_FromBucket(codeBucket, jsii.String("app.zip"), jsii.String("version-id")),
//   	Runtime: lambda.WebFunctionRuntime_NODEJS_24_X(),
//   	DefaultEndpointAuth: lambda.WebFunctionAuthType_IAM_AUTH,
//   	DefaultEndpoint: &WebFunctionDefaultEndpointOptions{
//   		EndpointLocation: lambda.WebFunctionEndpointLocation_PerRegion([]*string{
//   			jsii.String("us-east-1"),
//   			jsii.String("us-west-2"),
//   		}),
//   	},
//   })
//
type WebFunctionAuthType string

const (
	// Authentication is managed by the application itself.
	//
	// The Lambda service will not enforce authentication; ensure your handler code validates every request.
	WebFunctionAuthType_APPLICATION_MANAGED WebFunctionAuthType = "APPLICATION_MANAGED"
	// Authentication using AWS IAM.
	//
	// Requests must be signed with SigV4 and carry the `X-Amz-Content-SHA256` header, also for a
	// request without a body. A `MULTI_REGION` endpoint is signed with SigV4a, for the Region set `*`.
	WebFunctionAuthType_IAM_AUTH WebFunctionAuthType = "IAM_AUTH"
)

