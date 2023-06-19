package awscdk


// Construction properties of {@link CfnHook}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var properties interface{}
//
//   cfnHookProps := &CfnHookProps{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Properties: map[string]interface{}{
//   		"propertiesKey": properties,
//   	},
//   }
//
// Experimental.
type CfnHookProps struct {
	// The type of the hook (for example, "AWS::CodeDeploy::BlueGreen").
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The properties of the hook.
	// Experimental.
	Properties *map[string]interface{} `field:"optional" json:"properties" yaml:"properties"`
}

