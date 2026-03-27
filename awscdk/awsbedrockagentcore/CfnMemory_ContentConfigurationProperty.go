package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentConfigurationProperty := &ContentConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Level: jsii.String("level"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-contentconfiguration.html
//
type CfnMemory_ContentConfigurationProperty struct {
	// The type of content to deliver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-contentconfiguration.html#cfn-bedrockagentcore-memory-contentconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The level of content detail to deliver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-contentconfiguration.html#cfn-bedrockagentcore-memory-contentconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
}

