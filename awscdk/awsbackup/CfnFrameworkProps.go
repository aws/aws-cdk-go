package awsbackup


// Properties for defining a `CfnFramework`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var controlScope interface{}
//
//   cfnFrameworkProps := &CfnFrameworkProps{
//   	FrameworkControls: []interface{}{
//   		&FrameworkControlProperty{
//   			ControlName: jsii.String("controlName"),
//
//   			// the properties below are optional
//   			ControlInputParameters: []interface{}{
//   				&ControlInputParameterProperty{
//   					ParameterName: jsii.String("parameterName"),
//   					ParameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   			ControlScope: controlScope,
//   		},
//   	},
//
//   	// the properties below are optional
//   	FrameworkDescription: jsii.String("frameworkDescription"),
//   	FrameworkName: jsii.String("frameworkName"),
//   	FrameworkTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFrameworkProps struct {
	// Contains detailed information about all of the controls of a framework.
	//
	// Each framework must contain at least one control.
	FrameworkControls interface{} `field:"required" json:"frameworkControls" yaml:"frameworkControls"`
	// An optional description of the framework with a maximum 1,024 characters.
	FrameworkDescription *string `field:"optional" json:"frameworkDescription" yaml:"frameworkDescription"`
	// The unique name of a framework.
	//
	// This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
	FrameworkName *string `field:"optional" json:"frameworkName" yaml:"frameworkName"`
	// A list of tags with which to tag your framework.
	FrameworkTags interface{} `field:"optional" json:"frameworkTags" yaml:"frameworkTags"`
}

