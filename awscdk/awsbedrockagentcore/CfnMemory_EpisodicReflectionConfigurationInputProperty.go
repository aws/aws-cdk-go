package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   episodicReflectionConfigurationInputProperty := &EpisodicReflectionConfigurationInputProperty{
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//
//   	// the properties below are optional
//   	NamespaceTemplates: []*string{
//   		jsii.String("namespaceTemplates"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicreflectionconfigurationinput.html
//
type CfnMemory_EpisodicReflectionConfigurationInputProperty struct {
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicreflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicreflectionconfigurationinput-namespaces
	//
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicreflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicreflectionconfigurationinput-namespacetemplates
	//
	NamespaceTemplates *[]*string `field:"optional" json:"namespaceTemplates" yaml:"namespaceTemplates"`
}

