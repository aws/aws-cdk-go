package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamDeliveryResourceProperty := &StreamDeliveryResourceProperty{
//   	Kinesis: &KinesisResourceProperty{
//   		ContentConfigurations: []interface{}{
//   			&ContentConfigurationProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Level: jsii.String("level"),
//   			},
//   		},
//   		DataStreamArn: jsii.String("dataStreamArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-streamdeliveryresource.html
//
type CfnMemory_StreamDeliveryResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-streamdeliveryresource.html#cfn-bedrockagentcore-memory-streamdeliveryresource-kinesis
	//
	Kinesis interface{} `field:"optional" json:"kinesis" yaml:"kinesis"`
}

