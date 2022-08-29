// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Construction properties of {@link CfnHook}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var properties interface{}
//
//   cfnHookProps := &cfnHookProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	properties: map[string]interface{}{
//   		"propertiesKey": properties,
//   	},
//   }
//
type CfnHookProps struct {
	// The type of the hook (for example, "AWS::CodeDeploy::BlueGreen").
	Type *string `field:"required" json:"type" yaml:"type"`
	// The properties of the hook.
	Properties *map[string]interface{} `field:"optional" json:"properties" yaml:"properties"`
}

