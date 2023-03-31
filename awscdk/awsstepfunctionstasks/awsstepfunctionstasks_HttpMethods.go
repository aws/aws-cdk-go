package awsstepfunctionstasks


// Method type of a EKS call.
//
// Example:
//   import eks "github.com/aws/aws-cdk-go/awscdk"
//
//
//   myEksCluster := eks.NewCluster(this, jsii.String("my sample cluster"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_18(),
//   	clusterName: jsii.String("myEksCluster"),
//   })
//
//   tasks.NewEksCall(this, jsii.String("Call a EKS Endpoint"), &eksCallProps{
//   	cluster: myEksCluster,
//   	httpMethod: tasks.httpMethods_GET,
//   	httpPath: jsii.String("/api/v1/namespaces/default/pods"),
//   })
//
// Experimental.
type HttpMethods string

const (
	// Retrieve data from a server at the specified resource.
	// Experimental.
	HttpMethods_GET HttpMethods = "GET"
	// Send data to the API endpoint to create or update a resource.
	// Experimental.
	HttpMethods_POST HttpMethods = "POST"
	// Send data to the API endpoint to update or create a resource.
	// Experimental.
	HttpMethods_PUT HttpMethods = "PUT"
	// Delete the resource at the specified endpoint.
	// Experimental.
	HttpMethods_DELETE HttpMethods = "DELETE"
	// Apply partial modifications to the resource.
	// Experimental.
	HttpMethods_PATCH HttpMethods = "PATCH"
	// Retrieve data from a server at the specified resource without the response body.
	// Experimental.
	HttpMethods_HEAD HttpMethods = "HEAD"
)

