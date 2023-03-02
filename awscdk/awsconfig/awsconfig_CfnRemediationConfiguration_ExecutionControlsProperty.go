package awsconfig


// An ExecutionControls object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionControlsProperty := &executionControlsProperty{
//   	ssmControls: &ssmControlsProperty{
//   		concurrentExecutionRatePercentage: jsii.Number(123),
//   		errorPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnRemediationConfiguration_ExecutionControlsProperty struct {
	// A SsmControls object.
	SsmControls interface{} `field:"optional" json:"ssmControls" yaml:"ssmControls"`
}

