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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-executioncontrols.html
//
type CfnRemediationConfiguration_ExecutionControlsProperty struct {
	// A SsmControls object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-executioncontrols.html#cfn-config-remediationconfiguration-executioncontrols-ssmcontrols
	//
	SsmControls interface{} `field:"optional" json:"ssmControls" yaml:"ssmControls"`
}

