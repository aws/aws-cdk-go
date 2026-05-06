package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomainObjectType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainObjectTypeProps := &CfnDomainObjectTypeProps{
//   	DomainName: jsii.String("domainName"),
//   	Fields: map[string]interface{}{
//   		"fieldsKey": &DomainObjectTypeFieldProperty{
//   			"source": jsii.String("source"),
//   			"target": jsii.String("target"),
//
//   			// the properties below are optional
//   			"contentType": jsii.String("contentType"),
//   			"featureType": jsii.String("featureType"),
//   		},
//   	},
//   	ObjectTypeName: jsii.String("objectTypeName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html
//
type CfnDomainObjectTypeProps struct {
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// A map of the name and ObjectType field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-fields
	//
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The name of the domain object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-objecttypename
	//
	ObjectTypeName *string `field:"required" json:"objectTypeName" yaml:"objectTypeName"`
	// Description of the domain object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The default encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-encryptionkey
	//
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

