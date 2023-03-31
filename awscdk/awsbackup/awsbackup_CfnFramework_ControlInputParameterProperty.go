package awsbackup


// A list of parameters for a control.
//
// A control can have zero, one, or more than one parameter. An example of a control with two parameters is: "backup plan frequency is at least `daily` and the retention period is at least `1 year` ". The first parameter is `daily` . The second parameter is `1 year` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   controlInputParameterProperty := &controlInputParameterProperty{
//   	parameterName: jsii.String("parameterName"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
type CfnFramework_ControlInputParameterProperty struct {
	// The name of a parameter, for example, `BackupPlanFrequency` .
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The value of parameter, for example, `hourly` .
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

