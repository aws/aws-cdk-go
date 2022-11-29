package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndexProps := &cfnIndexProps{
//   	edition: jsii.String("edition"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	capacityUnits: &capacityUnitsConfigurationProperty{
//   		queryCapacityUnits: jsii.Number(123),
//   		storageCapacityUnits: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	documentMetadataConfigurations: []interface{}{
//   		&documentMetadataConfigurationProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			relevance: &relevanceProperty{
//   				duration: jsii.String("duration"),
//   				freshness: jsii.Boolean(false),
//   				importance: jsii.Number(123),
//   				rankOrder: jsii.String("rankOrder"),
//   				valueImportanceItems: []interface{}{
//   					&valueImportanceItemProperty{
//   						key: jsii.String("key"),
//   						value: jsii.Number(123),
//   					},
//   				},
//   			},
//   			search: &searchProperty{
//   				displayable: jsii.Boolean(false),
//   				facetable: jsii.Boolean(false),
//   				searchable: jsii.Boolean(false),
//   				sortable: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	serverSideEncryptionConfiguration: &serverSideEncryptionConfigurationProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userContextPolicy: jsii.String("userContextPolicy"),
//   	userTokenConfigurations: []interface{}{
//   		&userTokenConfigurationProperty{
//   			jsonTokenTypeConfiguration: &jsonTokenTypeConfigurationProperty{
//   				groupAttributeField: jsii.String("groupAttributeField"),
//   				userNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   			jwtTokenTypeConfiguration: &jwtTokenTypeConfigurationProperty{
//   				keyLocation: jsii.String("keyLocation"),
//
//   				// the properties below are optional
//   				claimRegex: jsii.String("claimRegex"),
//   				groupAttributeField: jsii.String("groupAttributeField"),
//   				issuer: jsii.String("issuer"),
//   				secretManagerArn: jsii.String("secretManagerArn"),
//   				url: jsii.String("url"),
//   				userNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   		},
//   	},
//   }
//
type CfnIndexProps struct {
	// Indicates whether the index is a enterprise edition index or a developer edition index.
	//
	// Valid values are `DEVELOPER_EDITION` and `ENTERPRISE_EDITION` .
	Edition *string `field:"required" json:"edition" yaml:"edition"`
	// The identifier of the index.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics.
	//
	// This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Specifies capacity units configured for your index.
	//
	// You can add and remove capacity units to tune an index to your requirements. You can set capacity units only for Enterprise edition indexes.
	CapacityUnits interface{} `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// A description of the index.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the properties of an index field.
	//
	// You can add either a custom or a built-in field. You can add and remove built-in fields at any time. When a built-in field is removed it's configuration reverts to the default for the field. Custom fields can't be removed from an index after they are added.
	DocumentMetadataConfigurations interface{} `field:"optional" json:"documentMetadataConfigurations" yaml:"documentMetadataConfigurations"`
	// The identifier of the AWS KMS customer managed key (CMK) to use to encrypt data indexed by Amazon Kendra.
	//
	// Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The user context policy.
	//
	// ATTRIBUTE_FILTER
	//
	// - All indexed content is searchable and displayable for all users. If there is an access control list, it is ignored. You can filter on user and group attributes.
	//
	// USER_TOKEN
	//
	// - Enables SSO and token-based user access control. All documents with no access control and all documents accessible to the user will be searchable and displayable.
	UserContextPolicy *string `field:"optional" json:"userContextPolicy" yaml:"userContextPolicy"`
	// Defines the type of user token used for the index.
	UserTokenConfigurations interface{} `field:"optional" json:"userTokenConfigurations" yaml:"userTokenConfigurations"`
}

