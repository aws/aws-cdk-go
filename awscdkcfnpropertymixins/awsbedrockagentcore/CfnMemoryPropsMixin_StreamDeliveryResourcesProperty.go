package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   streamDeliveryResourcesProperty := &StreamDeliveryResourcesProperty{
//   	Resources: []interface{}{
//   		&StreamDeliveryResourceProperty{
//   			Kinesis: &KinesisResourceProperty{
//   				ContentConfigurations: []interface{}{
//   					&ContentConfigurationProperty{
//   						Level: jsii.String("level"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				DataStreamArn: jsii.String("dataStreamArn"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-streamdeliveryresources.html
//
type CfnMemoryPropsMixin_StreamDeliveryResourcesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-streamdeliveryresources.html#cfn-bedrockagentcore-memory-streamdeliveryresources-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
}

