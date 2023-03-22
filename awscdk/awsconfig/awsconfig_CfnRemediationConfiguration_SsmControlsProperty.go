package awsconfig


// AWS Systems Manager (SSM) specific remediation controls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmControlsProperty := &ssmControlsProperty{
//   	concurrentExecutionRatePercentage: jsii.Number(123),
//   	errorPercentage: jsii.Number(123),
//   }
//
type CfnRemediationConfiguration_SsmControlsProperty struct {
	// The maximum percentage of remediation actions allowed to run in parallel on the non-compliant resources for that specific rule.
	//
	// You can specify a percentage, such as 10%. The default value is 10.
	ConcurrentExecutionRatePercentage *float64 `field:"optional" json:"concurrentExecutionRatePercentage" yaml:"concurrentExecutionRatePercentage"`
	// The percentage of errors that are allowed before SSM stops running automations on non-compliant resources for that specific rule.
	//
	// You can specify a percentage of errors, for example 10%. If you do not specifiy a percentage, the default is 50%. For example, if you set the ErrorPercentage to 40% for 10 non-compliant resources, then SSM stops running the automations when the fifth error is received.
	ErrorPercentage *float64 `field:"optional" json:"errorPercentage" yaml:"errorPercentage"`
}

