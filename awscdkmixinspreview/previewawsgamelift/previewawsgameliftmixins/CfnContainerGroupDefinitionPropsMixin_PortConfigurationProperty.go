package previewawsgameliftmixins


// Defines the ports on a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portConfigurationProperty := &PortConfigurationProperty{
//   	ContainerPortRanges: []interface{}{
//   		&ContainerPortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			Protocol: jsii.String("protocol"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-portconfiguration.html
//
type CfnContainerGroupDefinitionPropsMixin_PortConfigurationProperty struct {
	// Specifies one or more ranges of ports on a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-portconfiguration.html#cfn-gamelift-containergroupdefinition-portconfiguration-containerportranges
	//
	ContainerPortRanges interface{} `field:"optional" json:"containerPortRanges" yaml:"containerPortRanges"`
}

