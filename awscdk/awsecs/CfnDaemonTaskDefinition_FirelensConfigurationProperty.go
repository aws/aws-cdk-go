package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firelensConfigurationProperty := &FirelensConfigurationProperty{
//   	Options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-firelensconfiguration.html
//
type CfnDaemonTaskDefinition_FirelensConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-firelensconfiguration.html#cfn-ecs-daemontaskdefinition-firelensconfiguration-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-firelensconfiguration.html#cfn-ecs-daemontaskdefinition-firelensconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

