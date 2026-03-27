package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   streamDeliveryResourceProperty := &StreamDeliveryResourceProperty{
//   	Kinesis: &KinesisResourceProperty{
//   		ContentConfigurations: []interface{}{
//   			&ContentConfigurationProperty{
//   				Level: jsii.String("level"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		DataStreamArn: jsii.String("dataStreamArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-streamdeliveryresource.html
//
type CfnMemoryPropsMixin_StreamDeliveryResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-streamdeliveryresource.html#cfn-bedrockagentcore-memory-streamdeliveryresource-kinesis
	//
	Kinesis interface{} `field:"optional" json:"kinesis" yaml:"kinesis"`
}

