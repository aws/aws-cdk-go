package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndexProps := &CfnIndexProps{
//   	Edition: jsii.String("edition"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	CapacityUnits: &CapacityUnitsConfigurationProperty{
//   		QueryCapacityUnits: jsii.Number(123),
//   		StorageCapacityUnits: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	DocumentMetadataConfigurations: []interface{}{
//   		&DocumentMetadataConfigurationProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Relevance: &RelevanceProperty{
//   				Duration: jsii.String("duration"),
//   				Freshness: jsii.Boolean(false),
//   				Importance: jsii.Number(123),
//   				RankOrder: jsii.String("rankOrder"),
//   				ValueImportanceItems: []interface{}{
//   					&ValueImportanceItemProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Search: &SearchProperty{
//   				Displayable: jsii.Boolean(false),
//   				Facetable: jsii.Boolean(false),
//   				Searchable: jsii.Boolean(false),
//   				Sortable: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserContextPolicy: jsii.String("userContextPolicy"),
//   	UserTokenConfigurations: []interface{}{
//   		&UserTokenConfigurationProperty{
//   			JsonTokenTypeConfiguration: &JsonTokenTypeConfigurationProperty{
//   				GroupAttributeField: jsii.String("groupAttributeField"),
//   				UserNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   			JwtTokenTypeConfiguration: &JwtTokenTypeConfigurationProperty{
//   				KeyLocation: jsii.String("keyLocation"),
//
//   				// the properties below are optional
//   				ClaimRegex: jsii.String("claimRegex"),
//   				GroupAttributeField: jsii.String("groupAttributeField"),
//   				Issuer: jsii.String("issuer"),
//   				SecretManagerArn: jsii.String("secretManagerArn"),
//   				Url: jsii.String("url"),
//   				UserNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html
//
type CfnIndexProps struct {
	// Indicates whether the index is a Enterprise Edition index or a Developer Edition index.
	//
	// Valid values are `DEVELOPER_EDITION` and `ENTERPRISE_EDITION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-edition
	//
	Edition *string `field:"required" json:"edition" yaml:"edition"`
	// The name of the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics.
	//
	// This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-capacityunits
	//
	CapacityUnits interface{} `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// A description for the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the properties of an index field.
	//
	// You can add either a custom or a built-in field. You can add and remove built-in fields at any time. When a built-in field is removed it's configuration reverts to the default for the field. Custom fields can't be removed from an index after they are added.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-documentmetadataconfigurations
	//
	DocumentMetadataConfigurations interface{} `field:"optional" json:"documentMetadataConfigurations" yaml:"documentMetadataConfigurations"`
	// The identifier of the AWS KMS customer managed key (CMK) to use to encrypt data indexed by Amazon Kendra.
	//
	// Amazon Kendra doesn't support asymmetric CMKs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-serversideencryptionconfiguration
	//
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The user context policy.
	//
	// ATTRIBUTE_FILTER
	//
	// - All indexed content is searchable and displayable for all users. If you want to filter search results on user context, you can use the attribute filters of `_user_id` and `_group_ids` or you can provide user and group information in `UserContext` .
	//
	// USER_TOKEN
	//
	// - Enables token-based user access control to filter search results on user context. All documents with no access control and all documents accessible to the user will be searchable and displayable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-usercontextpolicy
	//
	UserContextPolicy *string `field:"optional" json:"userContextPolicy" yaml:"userContextPolicy"`
	// Defines the type of user token used for the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-usertokenconfigurations
	//
	UserTokenConfigurations interface{} `field:"optional" json:"userTokenConfigurations" yaml:"userTokenConfigurations"`
}

