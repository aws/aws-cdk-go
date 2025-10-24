package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnObjectType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnObjectTypeProps := &CfnObjectTypeProps{
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
//   	ObjectTypeName: jsii.String("objectTypeName"),
//
//   	// the properties below are optional
//   	AllowProfileCreation: jsii.Boolean(false),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	ExpirationDays: jsii.Number(123),
//   	Fields: []interface{}{
//   		&FieldMapProperty{
//   			Name: jsii.String("name"),
//   			ObjectTypeField: &ObjectTypeFieldProperty{
//   				ContentType: jsii.String("contentType"),
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//   			},
//   		},
//   	},
//   	Keys: []interface{}{
//   		&KeyMapProperty{
//   			Name: jsii.String("name"),
//   			ObjectTypeKeyList: []interface{}{
//   				&ObjectTypeKeyProperty{
//   					FieldNames: []*string{
//   						jsii.String("fieldNames"),
//   					},
//   					StandardIdentifiers: []*string{
//   						jsii.String("standardIdentifiers"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	MaxProfileObjectCount: jsii.Number(123),
//   	SourceLastUpdatedTimestampFormat: jsii.String("sourceLastUpdatedTimestampFormat"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateId: jsii.String("templateId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html
//
type CfnObjectTypeProps struct {
	// The description of the profile object type mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of the profile object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-objecttypename
	//
	ObjectTypeName *string `field:"required" json:"objectTypeName" yaml:"objectTypeName"`
	// Indicates whether a profile should be created when data is received if one doesn’t exist for an object of this type.
	//
	// The default is `FALSE` . If the AllowProfileCreation flag is set to `FALSE` , then the service tries to fetch a standard profile and associate this object with the profile. If it is set to `TRUE` , and if no match is found, then the service creates a new standard profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-allowprofilecreation
	//
	AllowProfileCreation interface{} `field:"optional" json:"allowProfileCreation" yaml:"allowProfileCreation"`
	// The customer-provided key to encrypt the profile object that will be created in this profile object type mapping.
	//
	// If not specified the system will use the encryption key of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-encryptionkey
	//
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The number of days until the data of this type expires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-expirationdays
	//
	ExpirationDays *float64 `field:"optional" json:"expirationDays" yaml:"expirationDays"`
	// A list of field definitions for the object type mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// A list of keys that can be used to map data to the profile or search for the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-keys
	//
	Keys interface{} `field:"optional" json:"keys" yaml:"keys"`
	// The amount of profile object max count assigned to the object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-maxprofileobjectcount
	//
	MaxProfileObjectCount *float64 `field:"optional" json:"maxProfileObjectCount" yaml:"maxProfileObjectCount"`
	// The format of your sourceLastUpdatedTimestamp that was previously set up.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-sourcelastupdatedtimestampformat
	//
	SourceLastUpdatedTimestampFormat *string `field:"optional" json:"sourceLastUpdatedTimestampFormat" yaml:"sourceLastUpdatedTimestampFormat"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A unique identifier for the template mapping.
	//
	// This can be used instead of specifying the Keys and Fields properties directly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html#cfn-customerprofiles-objecttype-templateid
	//
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
}

