package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnObjectType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnObjectTypeProps := &CfnObjectTypeProps{
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	AllowProfileCreation: jsii.Boolean(false),
//   	Description: jsii.String("description"),
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
//   	ObjectTypeName: jsii.String("objectTypeName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateId: jsii.String("templateId"),
//   }
//
type CfnObjectTypeProps struct {
	// The unique name of the domain.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Indicates whether a profile should be created when data is received if one doesn’t exist for an object of this type.
	//
	// The default is `FALSE` . If the AllowProfileCreation flag is set to `FALSE` , then the service tries to fetch a standard profile and associate this object with the profile. If it is set to `TRUE` , and if no match is found, then the service creates a new standard profile.
	AllowProfileCreation interface{} `field:"optional" json:"allowProfileCreation" yaml:"allowProfileCreation"`
	// The description of the profile object type mapping.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The customer-provided key to encrypt the profile object that will be created in this profile object type mapping.
	//
	// If not specified the system will use the encryption key of the domain.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The number of days until the data of this type expires.
	ExpirationDays *float64 `field:"optional" json:"expirationDays" yaml:"expirationDays"`
	// A list of field definitions for the object type mapping.
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// A list of keys that can be used to map data to the profile or search for the profile.
	Keys interface{} `field:"optional" json:"keys" yaml:"keys"`
	// The name of the profile object type.
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A unique identifier for the template mapping.
	//
	// This can be used instead of specifying the Keys and Fields properties directly.
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
}

