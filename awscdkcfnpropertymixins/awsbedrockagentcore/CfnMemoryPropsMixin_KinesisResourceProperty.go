package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   kinesisResourceProperty := &KinesisResourceProperty{
//   	ContentConfigurations: []interface{}{
//   		&ContentConfigurationProperty{
//   			Level: jsii.String("level"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	DataStreamArn: jsii.String("dataStreamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-kinesisresource.html
//
type CfnMemoryPropsMixin_KinesisResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-kinesisresource.html#cfn-bedrockagentcore-memory-kinesisresource-contentconfigurations
	//
	ContentConfigurations interface{} `field:"optional" json:"contentConfigurations" yaml:"contentConfigurations"`
	// ARN format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-kinesisresource.html#cfn-bedrockagentcore-memory-kinesisresource-datastreamarn
	//
	DataStreamArn *string `field:"optional" json:"dataStreamArn" yaml:"dataStreamArn"`
}

