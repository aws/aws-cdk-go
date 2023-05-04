package awsconfig


// An ExecutionControls object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionControlsProperty := &ExecutionControlsProperty{
//   	SsmControls: &SsmControlsProperty{
//   		ConcurrentExecutionRatePercentage: jsii.Number(123),
//   		ErrorPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnRemediationConfiguration_ExecutionControlsProperty struct {
	// A SsmControls object.
	SsmControls interface{} `field:"optional" json:"ssmControls" yaml:"ssmControls"`
}

