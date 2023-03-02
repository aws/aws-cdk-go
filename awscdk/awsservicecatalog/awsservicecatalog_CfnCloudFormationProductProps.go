package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCloudFormationProduct`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var info interface{}
//
//   cfnCloudFormationProductProps := &cfnCloudFormationProductProps{
//   	name: jsii.String("name"),
//   	owner: jsii.String("owner"),
//   	provisioningArtifactParameters: []interface{}{
//   		&provisioningArtifactPropertiesProperty{
//   			info: info,
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			disableTemplateValidation: jsii.Boolean(false),
//   			name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   	distributor: jsii.String("distributor"),
//   	replaceProvisioningArtifacts: jsii.Boolean(false),
//   	supportDescription: jsii.String("supportDescription"),
//   	supportEmail: jsii.String("supportEmail"),
//   	supportUrl: jsii.String("supportUrl"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCloudFormationProductProps struct {
	// The name of the product.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The owner of the product.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The configuration of the provisioning artifact (also known as a version).
	ProvisioningArtifactParameters interface{} `field:"required" json:"provisioningArtifactParameters" yaml:"provisioningArtifactParameters"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the product.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The distributor of the product.
	Distributor *string `field:"optional" json:"distributor" yaml:"distributor"`
	// This property is turned off by default.
	//
	// If turned off, you can update provisioning artifacts or product attributes (such as description, distributor, name, owner, and more) and the associated provisioning artifacts will retain the same unique identifier. Provisioning artifacts are matched within the CloudFormationProduct resource, and only those that have been updated will be changed. Provisioning artifacts are matched by a combinaton of provisioning artifact template URL and name.
	//
	// If turned on, provisioning artifacts will be given a new unique identifier when you update the product or provisioning artifacts.
	ReplaceProvisioningArtifacts interface{} `field:"optional" json:"replaceProvisioningArtifacts" yaml:"replaceProvisioningArtifacts"`
	// The support information about the product.
	SupportDescription *string `field:"optional" json:"supportDescription" yaml:"supportDescription"`
	// The contact email for product support.
	SupportEmail *string `field:"optional" json:"supportEmail" yaml:"supportEmail"`
	// The contact URL for product support.
	//
	// `^https?:\/\//` / is the pattern used to validate SupportUrl.
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// One or more tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

