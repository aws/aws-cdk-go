package previewawsbedrockmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnBlueprintPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schema interface{}
//
//   cfnBlueprintMixinProps := &CfnBlueprintMixinProps{
//   	BlueprintName: jsii.String("blueprintName"),
//   	KmsEncryptionContext: map[string]*string{
//   		"kmsEncryptionContextKey": jsii.String("kmsEncryptionContext"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Schema: schema,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html
//
type CfnBlueprintMixinProps struct {
	// The blueprint's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-blueprintname
	//
	BlueprintName *string `field:"optional" json:"blueprintName" yaml:"blueprintName"`
	// Name-value pairs to include as an encryption context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-kmsencryptioncontext
	//
	KmsEncryptionContext interface{} `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	// The AWS  key to use for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The blueprint's schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// List of Tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The blueprint's type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-blueprint.html#cfn-bedrock-blueprint-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

