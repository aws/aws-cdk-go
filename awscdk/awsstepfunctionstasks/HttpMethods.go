package awsstepfunctionstasks


// Method type of a EKS call.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"
//
//
//   myEksCluster := eks.NewCluster(this, jsii.String("my sample cluster"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_32(),
//   	ClusterName: jsii.String("myEksCluster"),
//   	KubectlLayer: kubectlv32.NewKubectlV32Layer(this, jsii.String("kubectl")),
//   })
//
//   tasks.NewEksCall(this, jsii.String("Call a EKS Endpoint"), &EksCallProps{
//   	Cluster: myEksCluster,
//   	HttpMethod: tasks.HttpMethods_GET,
//   	HttpPath: jsii.String("/api/v1/namespaces/default/pods"),
//   })
//
type HttpMethods string

const (
	// Retrieve data from a server at the specified resource.
	HttpMethods_GET HttpMethods = "GET"
	// Send data to the API endpoint to create or update a resource.
	HttpMethods_POST HttpMethods = "POST"
	// Send data to the API endpoint to update or create a resource.
	HttpMethods_PUT HttpMethods = "PUT"
	// Delete the resource at the specified endpoint.
	HttpMethods_DELETE HttpMethods = "DELETE"
	// Apply partial modifications to the resource.
	HttpMethods_PATCH HttpMethods = "PATCH"
	// Retrieve data from a server at the specified resource without the response body.
	HttpMethods_HEAD HttpMethods = "HEAD"
)

