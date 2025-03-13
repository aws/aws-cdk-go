package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnBlueprint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   cfnBlueprintProps := &CfnBlueprintProps{
//   	BlueprintName: jsii.String("blueprintName"),
//   	Schema: schema,
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	KmsEncryptionContext: map[string]*string{
//   		"kmsEncryptionContextKey": jsii.String("kmsEncryptionContext"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html
//
type CfnBlueprintProps struct {
	// Name of the Blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-blueprintname
	//
	BlueprintName *string `field:"required" json:"blueprintName" yaml:"blueprintName"`
	// Schema of the blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-schema
	//
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
	// Modality Type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// KMS encryption context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-kmsencryptioncontext
	//
	KmsEncryptionContext interface{} `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	// KMS key identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// List of Tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

