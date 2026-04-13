package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	LogDriver: jsii.String("logDriver"),
//   	Options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	SecretOptions: []interface{}{
//   		&SecretProperty{
//   			Name: jsii.String("name"),
//   			ValueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-logconfiguration.html
//
type CfnDaemonTaskDefinitionPropsMixin_LogConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-logconfiguration.html#cfn-ecs-daemontaskdefinition-logconfiguration-logdriver
	//
	LogDriver *string `field:"optional" json:"logDriver" yaml:"logDriver"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-logconfiguration.html#cfn-ecs-daemontaskdefinition-logconfiguration-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-logconfiguration.html#cfn-ecs-daemontaskdefinition-logconfiguration-secretoptions
	//
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

