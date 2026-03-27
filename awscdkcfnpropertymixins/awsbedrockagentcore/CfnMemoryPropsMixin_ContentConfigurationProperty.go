package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   contentConfigurationProperty := &ContentConfigurationProperty{
//   	Level: jsii.String("level"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-contentconfiguration.html
//
type CfnMemoryPropsMixin_ContentConfigurationProperty struct {
	// The level of content detail to deliver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-contentconfiguration.html#cfn-bedrockagentcore-memory-contentconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
	// The type of content to deliver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-contentconfiguration.html#cfn-bedrockagentcore-memory-contentconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

