package awscdk


// Construction properties of `CfnHook`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
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
type CfnHookProps struct {
	// The type of the hook (for example, "AWS::CodeDeploy::BlueGreen").
	Type *string `field:"required" json:"type" yaml:"type"`
	// The properties of the hook.
	// Default: - no properties.
	//
	Properties *map[string]interface{} `field:"optional" json:"properties" yaml:"properties"`
}

