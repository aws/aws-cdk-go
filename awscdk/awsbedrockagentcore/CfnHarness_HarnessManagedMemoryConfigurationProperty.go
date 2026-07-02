package awsbedrockagentcore


// Configuration for managed memory.
//
// The harness creates and manages a memory resource in the customer's account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessManagedMemoryConfigurationProperty := &HarnessManagedMemoryConfigurationProperty{
//   	Arn: jsii.String("arn"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	EventExpiryDuration: jsii.Number(123),
//   	Strategies: []*string{
//   		jsii.String("strategies"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmanagedmemoryconfiguration.html
//
type CfnHarness_HarnessManagedMemoryConfigurationProperty struct {
	// The ARN of the managed memory resource.
	//
	// Read-only, populated by the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmanagedmemoryconfiguration.html#cfn-bedrockagentcore-harness-harnessmanagedmemoryconfiguration-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Customer-managed KMS key ARN.
	//
	// Defaults to AWS-owned key. Not updatable after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmanagedmemoryconfiguration.html#cfn-bedrockagentcore-harness-harnessmanagedmemoryconfiguration-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Event retention in days.
	//
	// Defaults to 30.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmanagedmemoryconfiguration.html#cfn-bedrockagentcore-harness-harnessmanagedmemoryconfiguration-eventexpiryduration
	//
	EventExpiryDuration *float64 `field:"optional" json:"eventExpiryDuration" yaml:"eventExpiryDuration"`
	// Strategy types to enable.
	//
	// Defaults to [SEMANTIC, SUMMARIZATION].
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmanagedmemoryconfiguration.html#cfn-bedrockagentcore-harness-harnessmanagedmemoryconfiguration-strategies
	//
	Strategies *[]*string `field:"optional" json:"strategies" yaml:"strategies"`
}

