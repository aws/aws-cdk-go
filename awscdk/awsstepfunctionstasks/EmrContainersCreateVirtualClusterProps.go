package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties to define a EMR Containers CreateVirtualCluster Task on an EKS cluster.
//
// Example:
//   tasks.NewEmrContainersCreateVirtualCluster(this, jsii.String("Create a Virtual Cluster"), &EmrContainersCreateVirtualClusterProps{
//   	EksCluster: tasks.EksClusterInput_FromTaskInput(sfn.TaskInput_FromText(jsii.String("clusterId"))),
//   	EksNamespace: jsii.String("specified-namespace"),
//   })
//
// Experimental.
type EmrContainersCreateVirtualClusterProps struct {
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
	// EKS Cluster or task input that contains the name of the cluster.
	// Experimental.
	EksCluster EksClusterInput `field:"required" json:"eksCluster" yaml:"eksCluster"`
	// The namespace of an EKS cluster.
	// Experimental.
	EksNamespace *string `field:"optional" json:"eksNamespace" yaml:"eksNamespace"`
	// The tags assigned to the virtual cluster.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Name of the virtual cluster that will be created.
	// Experimental.
	VirtualClusterName *string `field:"optional" json:"virtualClusterName" yaml:"virtualClusterName"`
}

