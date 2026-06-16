package awsopensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCollectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCollectionMixinProps := &CfnCollectionMixinProps{
//   	CollectionGroupName: jsii.String("collectionGroupName"),
//   	DeletionProtection: jsii.String("deletionProtection"),
//   	Description: jsii.String("description"),
//   	EncryptionConfig: &EncryptionConfigProperty{
//   		AwsOwnedKey: jsii.Boolean(false),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	Name: jsii.String("name"),
//   	StandbyReplicas: jsii.String("standbyReplicas"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	VectorOptions: &VectorOptionsProperty{
//   		ServerlessVectorAcceleration: jsii.String("serverlessVectorAcceleration"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html
//
type CfnCollectionMixinProps struct {
	// The name of the collection group to associate with the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html#cfn-opensearchserverless-collection-collectiongroupname
	//
	CollectionGroupName *string `field:"optional" json:"collectionGroupName" yaml:"collectionGroupName"`
	// The deletion protection state of the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html#cfn-opensearchserverless-collection-deletionprotection
	//
	DeletionProtection *string `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// A description of the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html#cfn-opensearchserverless-collection-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Encryption settings for the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html#cfn-opensearchserverless-collection-encryptionconfig
	//
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The name of the collection.
	//
	// Collection names must meet the following criteria:
	//
	// - Starts with a lowercase letter
	// - Unique to your account and AWS Region
	// - Contains between 3 and 28 characters
	// - Contains only lowercase letters a-z, the numbers 0-9, and the hyphen (-).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html#cfn-opensearchserverless-collection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates whether to use standby replicas for the collection.
	//
	// You can't update this property after the collection is already created. If you attempt to modify this property, the collection continues to use the original value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html#cfn-opensearchserverless-collection-standbyreplicas
	//
	StandbyReplicas *string `field:"optional" json:"standbyReplicas" yaml:"standbyReplicas"`
	// An arbitrary set of tags (key–value pairs) to associate with the collection.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html#cfn-opensearchserverless-collection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of collection.
	//
	// Possible values are `SEARCH` , `TIMESERIES` , and `VECTORSEARCH` . For more information, see [Choosing a collection type](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-overview.html#serverless-usecase) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html#cfn-opensearchserverless-collection-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Vector search configuration options for the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collection.html#cfn-opensearchserverless-collection-vectoroptions
	//
	VectorOptions interface{} `field:"optional" json:"vectorOptions" yaml:"vectorOptions"`
}

