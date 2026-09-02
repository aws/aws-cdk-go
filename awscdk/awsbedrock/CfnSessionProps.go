package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSession`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSessionProps := &CfnSessionProps{
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	SessionMetadata: map[string]*string{
//   		"sessionMetadataKey": jsii.String("sessionMetadata"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-session.html
//
type CfnSessionProps struct {
	// The Amazon Resource Name (ARN) of the KMS key to use to encrypt the session data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-session.html#cfn-bedrock-session-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// A map of key-value pairs containing attributes to be persisted across the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-session.html#cfn-bedrock-session-sessionmetadata
	//
	SessionMetadata interface{} `field:"optional" json:"sessionMetadata" yaml:"sessionMetadata"`
	// A list of tags associated with the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-session.html#cfn-bedrock-session-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

