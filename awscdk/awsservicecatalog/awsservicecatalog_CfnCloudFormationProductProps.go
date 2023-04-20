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
//   cfnCloudFormationProductProps := &CfnCloudFormationProductProps{
//   	Name: jsii.String("name"),
//   	Owner: jsii.String("owner"),
//
//   	// the properties below are optional
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   	Distributor: jsii.String("distributor"),
//   	ProvisioningArtifactParameters: []interface{}{
//   		&ProvisioningArtifactPropertiesProperty{
//   			Info: info,
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			DisableTemplateValidation: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	ReplaceProvisioningArtifacts: jsii.Boolean(false),
//   	SourceConnection: &SourceConnectionProperty{
//   		ConnectionParameters: &ConnectionParametersProperty{
//   			CodeStar: &CodeStarParametersProperty{
//   				ArtifactPath: jsii.String("artifactPath"),
//   				Branch: jsii.String("branch"),
//   				ConnectionArn: jsii.String("connectionArn"),
//   				Repository: jsii.String("repository"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	SupportDescription: jsii.String("supportDescription"),
//   	SupportEmail: jsii.String("supportEmail"),
//   	SupportUrl: jsii.String("supportUrl"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCloudFormationProductProps struct {
	// The name of the product.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The owner of the product.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the product.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The distributor of the product.
	Distributor *string `field:"optional" json:"distributor" yaml:"distributor"`
	// The configuration of the provisioning artifact (also known as a version).
	ProvisioningArtifactParameters interface{} `field:"optional" json:"provisioningArtifactParameters" yaml:"provisioningArtifactParameters"`
	// This property is turned off by default.
	//
	// If turned off, you can update provisioning artifacts or product attributes (such as description, distributor, name, owner, and more) and the associated provisioning artifacts will retain the same unique identifier. Provisioning artifacts are matched within the CloudFormationProduct resource, and only those that have been updated will be changed. Provisioning artifacts are matched by a combinaton of provisioning artifact template URL and name.
	//
	// If turned on, provisioning artifacts will be given a new unique identifier when you update the product or provisioning artifacts.
	ReplaceProvisioningArtifacts interface{} `field:"optional" json:"replaceProvisioningArtifacts" yaml:"replaceProvisioningArtifacts"`
	// A top level `ProductViewDetail` response containing details about the product’s connection.
	//
	// AWS Service Catalog returns this field for the `CreateProduct` , `UpdateProduct` , `DescribeProductAsAdmin` , and `SearchProductAsAdmin` APIs. This response contains the same fields as the `ConnectionParameters` request, with the addition of the `LastSync` response.
	SourceConnection interface{} `field:"optional" json:"sourceConnection" yaml:"sourceConnection"`
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

