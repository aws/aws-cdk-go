package awsbackup


// Contains detailed information about all of the controls of a framework.
//
// Each framework must contain at least one control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var controlScope interface{}
//
//   frameworkControlProperty := &frameworkControlProperty{
//   	controlName: jsii.String("controlName"),
//
//   	// the properties below are optional
//   	controlInputParameters: []interface{}{
//   		&controlInputParameterProperty{
//   			parameterName: jsii.String("parameterName"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	controlScope: controlScope,
//   }
//
type CfnFramework_FrameworkControlProperty struct {
	// The name of a control.
	//
	// This name is between 1 and 256 characters.
	ControlName *string `field:"required" json:"controlName" yaml:"controlName"`
	// A list of `ParameterName` and `ParameterValue` pairs.
	ControlInputParameters interface{} `field:"optional" json:"controlInputParameters" yaml:"controlInputParameters"`
	// The scope of a control.
	//
	// The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans. For more information, see [`ControlScope` .](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ControlScope.html)
	ControlScope interface{} `field:"optional" json:"controlScope" yaml:"controlScope"`
}

