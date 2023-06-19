package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awseks"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for calling a EKS endpoint with EksCall.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   myEksCluster := eks.NewCluster(this, jsii.String("my sample cluster"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_18(),
//   	ClusterName: jsii.String("myEksCluster"),
//   })
//
//   tasks.NewEksCall(this, jsii.String("Call a EKS Endpoint"), &EksCallProps{
//   	Cluster: myEksCluster,
//   	HttpMethod: tasks.HttpMethods_GET,
//   	HttpPath: jsii.String("/api/v1/namespaces/default/pods"),
//   })
//
// Experimental.
type EksCallProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The EKS cluster.
	// Experimental.
	Cluster awseks.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// HTTP method ("GET", "POST", "PUT", ...) part of HTTP request.
	// Experimental.
	HttpMethod HttpMethods `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// HTTP path of the Kubernetes REST API operation For example: /api/v1/namespaces/default/pods.
	// Experimental.
	HttpPath *string `field:"required" json:"httpPath" yaml:"httpPath"`
	// Query Parameters part of HTTP request.
	// Experimental.
	QueryParameters *map[string]*[]*string `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// Request body part of HTTP request.
	// Experimental.
	RequestBody awsstepfunctions.TaskInput `field:"optional" json:"requestBody" yaml:"requestBody"`
}

