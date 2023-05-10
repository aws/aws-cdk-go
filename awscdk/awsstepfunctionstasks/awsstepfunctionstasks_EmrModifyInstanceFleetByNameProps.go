package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for EmrModifyInstanceFleetByName.
//
// Example:
//   tasks.NewEmrModifyInstanceFleetByName(this, jsii.String("Task"), &EmrModifyInstanceFleetByNameProps{
//   	ClusterId: jsii.String("ClusterId"),
//   	InstanceFleetName: jsii.String("InstanceFleetName"),
//   	TargetOnDemandCapacity: jsii.Number(2),
//   	TargetSpotCapacity: jsii.Number(0),
//   })
//
// Experimental.
type EmrModifyInstanceFleetByNameProps struct {
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
	// The ClusterId to update.
	// Experimental.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The InstanceFleetName to update.
	// Experimental.
	InstanceFleetName *string `field:"required" json:"instanceFleetName" yaml:"instanceFleetName"`
	// The target capacity of On-Demand units for the instance fleet.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetModifyConfig.html
	//
	// Experimental.
	TargetOnDemandCapacity *float64 `field:"required" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetModifyConfig.html
	//
	// Experimental.
	TargetSpotCapacity *float64 `field:"required" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

