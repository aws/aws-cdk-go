package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamDeliveryResourcesProperty := &StreamDeliveryResourcesProperty{
//   	Resources: []interface{}{
//   		&StreamDeliveryResourceProperty{
//   			Kinesis: &KinesisResourceProperty{
//   				ContentConfigurations: []interface{}{
//   					&ContentConfigurationProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						Level: jsii.String("level"),
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
type CfnMemory_StreamDeliveryResourcesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-streamdeliveryresources.html#cfn-bedrockagentcore-memory-streamdeliveryresources-resources
	//
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
}

