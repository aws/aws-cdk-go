package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisResourceProperty := &KinesisResourceProperty{
//   	ContentConfigurations: []interface{}{
//   		&ContentConfigurationProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Level: jsii.String("level"),
//   		},
//   	},
//   	DataStreamArn: jsii.String("dataStreamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-kinesisresource.html
//
type CfnMemory_KinesisResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-kinesisresource.html#cfn-bedrockagentcore-memory-kinesisresource-contentconfigurations
	//
	ContentConfigurations interface{} `field:"required" json:"contentConfigurations" yaml:"contentConfigurations"`
	// ARN format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-kinesisresource.html#cfn-bedrockagentcore-memory-kinesisresource-datastreamarn
	//
	DataStreamArn *string `field:"required" json:"dataStreamArn" yaml:"dataStreamArn"`
}

