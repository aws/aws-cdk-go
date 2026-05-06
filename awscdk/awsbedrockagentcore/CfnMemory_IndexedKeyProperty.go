package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   indexedKeyProperty := &IndexedKeyProperty{
//   	Key: jsii.String("key"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-indexedkey.html
//
type CfnMemory_IndexedKeyProperty struct {
	// Key name for metadata fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-indexedkey.html#cfn-bedrockagentcore-memory-indexedkey-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// Supported data types for metadata values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-indexedkey.html#cfn-bedrockagentcore-memory-indexedkey-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

