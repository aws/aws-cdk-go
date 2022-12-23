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
//   cfnFrameworkProps := &cfnFrameworkProps{
//   	frameworkControls: []interface{}{
//   		&frameworkControlProperty{
//   			controlName: jsii.String("controlName"),
//
//   			// the properties below are optional
//   			controlInputParameters: []interface{}{
//   				&controlInputParameterProperty{
//   					parameterName: jsii.String("parameterName"),
//   					parameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   			controlScope: controlScope,
//   		},
//   	},
//
//   	// the properties below are optional
//   	frameworkDescription: jsii.String("frameworkDescription"),
//   	frameworkName: jsii.String("frameworkName"),
//   	frameworkTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

