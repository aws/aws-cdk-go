package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDomainObjectTypePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDomainObjectTypeMixinProps := &CfnDomainObjectTypeMixinProps{
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	Fields: map[string]interface{}{
//   		"fieldsKey": &DomainObjectTypeFieldProperty{
//   			"contentType": jsii.String("contentType"),
//   			"featureType": jsii.String("featureType"),
//   			"source": jsii.String("source"),
//   			"target": jsii.String("target"),
//   		},
//   	},
//   	ObjectTypeName: jsii.String("objectTypeName"),
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
type CfnDomainObjectTypeMixinProps struct {
	// Description of the domain object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The default encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-encryptionkey
	//
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// A map of the name and ObjectType field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The name of the domain object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-objecttypename
	//
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domainobjecttype.html#cfn-customerprofiles-domainobjecttype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

