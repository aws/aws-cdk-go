package mixinsawsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCloudFormationProductPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var info interface{}
//
//   cfnCloudFormationProductMixinProps := &CfnCloudFormationProductMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   	Distributor: jsii.String("distributor"),
//   	Name: jsii.String("name"),
//   	Owner: jsii.String("owner"),
//   	ProductType: jsii.String("productType"),
//   	ProvisioningArtifactParameters: []interface{}{
//   		&ProvisioningArtifactPropertiesProperty{
//   			Description: jsii.String("description"),
//   			DisableTemplateValidation: jsii.Boolean(false),
//   			Info: info,
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html
//
type CfnCloudFormationProductMixinProps struct {
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The distributor of the product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-distributor
	//
	Distributor *string `field:"optional" json:"distributor" yaml:"distributor"`
	// The name of the product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The owner of the product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-owner
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The type of product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-producttype
	//
	ProductType *string `field:"optional" json:"productType" yaml:"productType"`
	// The configuration of the provisioning artifact (also known as a version).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-provisioningartifactparameters
	//
	ProvisioningArtifactParameters interface{} `field:"optional" json:"provisioningArtifactParameters" yaml:"provisioningArtifactParameters"`
	// This property is turned off by default.
	//
	// If turned off, you can update provisioning artifacts or product attributes (such as description, distributor, name, owner, and more) and the associated provisioning artifacts will retain the same unique identifier. Provisioning artifacts are matched within the CloudFormationProduct resource, and only those that have been updated will be changed. Provisioning artifacts are matched by a combinaton of provisioning artifact template URL and name.
	//
	// If turned on, provisioning artifacts will be given a new unique identifier when you update the product or provisioning artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-replaceprovisioningartifacts
	//
	ReplaceProvisioningArtifacts interface{} `field:"optional" json:"replaceProvisioningArtifacts" yaml:"replaceProvisioningArtifacts"`
	// A top level `ProductViewDetail` response containing details about the product’s connection.
	//
	// AWS Service Catalog returns this field for the `CreateProduct` , `UpdateProduct` , `DescribeProductAsAdmin` , and `SearchProductAsAdmin` APIs. This response contains the same fields as the `ConnectionParameters` request, with the addition of the `LastSync` response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-sourceconnection
	//
	SourceConnection interface{} `field:"optional" json:"sourceConnection" yaml:"sourceConnection"`
	// The support information about the product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supportdescription
	//
	SupportDescription *string `field:"optional" json:"supportDescription" yaml:"supportDescription"`
	// The contact email for product support.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supportemail
	//
	SupportEmail *string `field:"optional" json:"supportEmail" yaml:"supportEmail"`
	// The contact URL for product support.
	//
	// `^https?:\/\//` / is the pattern used to validate SupportUrl.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supporturl
	//
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// One or more tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

