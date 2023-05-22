package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentPreferenceProperty := &DeploymentPreferenceProperty{
//   	Enabled: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Alarms: []*string{
//   		jsii.String("alarms"),
//   	},
//   	Hooks: &HooksProperty{
//   		PostTraffic: jsii.String("postTraffic"),
//   		PreTraffic: jsii.String("preTraffic"),
//   	},
//   }
//
type CfnFunction_DeploymentPreferenceProperty struct {
	// `CfnFunction.DeploymentPreferenceProperty.Enabled`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// `CfnFunction.DeploymentPreferenceProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnFunction.DeploymentPreferenceProperty.Alarms`.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// `CfnFunction.DeploymentPreferenceProperty.Hooks`.
	Hooks interface{} `field:"optional" json:"hooks" yaml:"hooks"`
}

