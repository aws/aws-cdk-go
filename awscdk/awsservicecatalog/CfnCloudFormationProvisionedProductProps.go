package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCloudFormationProvisionedProduct`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudFormationProvisionedProductProps := &CfnCloudFormationProvisionedProductProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	PathId: jsii.String("pathId"),
//   	PathName: jsii.String("pathName"),
//   	ProductId: jsii.String("productId"),
//   	ProductName: jsii.String("productName"),
//   	ProvisionedProductName: jsii.String("provisionedProductName"),
//   	ProvisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	ProvisioningArtifactName: jsii.String("provisioningArtifactName"),
//   	ProvisioningParameters: []interface{}{
//   		&ProvisioningParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProvisioningPreferences: &ProvisioningPreferencesProperty{
//   		StackSetAccounts: []*string{
//   			jsii.String("stackSetAccounts"),
//   		},
//   		StackSetFailureToleranceCount: jsii.Number(123),
//   		StackSetFailureTolerancePercentage: jsii.Number(123),
//   		StackSetMaxConcurrencyCount: jsii.Number(123),
//   		StackSetMaxConcurrencyPercentage: jsii.Number(123),
//   		StackSetOperationType: jsii.String("stackSetOperationType"),
//   		StackSetRegions: []*string{
//   			jsii.String("stackSetRegions"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html
//
type CfnCloudFormationProvisionedProductProps struct {
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// Passed to AWS CloudFormation .
	//
	// The SNS topic ARNs to which to publish stack-related events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-notificationarns
	//
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// The path identifier of the product.
	//
	// This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use [ListLaunchPaths](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ListLaunchPaths.html) .
	//
	// > You must provide the name or ID, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-pathid
	//
	PathId *string `field:"optional" json:"pathId" yaml:"pathId"`
	// The name of the path.
	//
	// This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use [ListLaunchPaths](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ListLaunchPaths.html) .
	//
	// > You must provide the name or ID, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-pathname
	//
	PathName *string `field:"optional" json:"pathName" yaml:"pathName"`
	// The product identifier.
	//
	// > You must specify either the ID or the name of the product, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-productid
	//
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// The name of the Service Catalog product.
	//
	// Each time a stack is created or updated, if `ProductName` is provided it will successfully resolve to `ProductId` as long as only one product exists in the account or Region with that `ProductName` .
	//
	// > You must specify either the name or the ID of the product, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-productname
	//
	ProductName *string `field:"optional" json:"productName" yaml:"productName"`
	// A user-friendly name for the provisioned product.
	//
	// This value must be unique for the AWS account and cannot be updated after the product is provisioned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisionedproductname
	//
	ProvisionedProductName *string `field:"optional" json:"provisionedProductName" yaml:"provisionedProductName"`
	// The identifier of the provisioning artifact (also known as a version).
	//
	// > You must specify either the ID or the name of the provisioning artifact, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningartifactid
	//
	ProvisioningArtifactId *string `field:"optional" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// The name of the provisioning artifact (also known as a version) for the product.
	//
	// This name must be unique for the product.
	//
	// > You must specify either the name or the ID of the provisioning artifact, but not both. You must also specify either the name or the ID of the product, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningartifactname
	//
	ProvisioningArtifactName *string `field:"optional" json:"provisioningArtifactName" yaml:"provisioningArtifactName"`
	// Parameters specified by the administrator that are required for provisioning the product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningparameters
	//
	ProvisioningParameters interface{} `field:"optional" json:"provisioningParameters" yaml:"provisioningParameters"`
	// StackSet preferences that are required for provisioning the product or updating a provisioned product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences
	//
	ProvisioningPreferences interface{} `field:"optional" json:"provisioningPreferences" yaml:"provisioningPreferences"`
	// One or more tags.
	//
	// > Requires the provisioned product to have an [ResourceUpdateConstraint](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-resourceupdateconstraint.html) resource with `TagUpdatesOnProvisionedProduct` set to `ALLOWED` to allow tag updates. If `RESOURCE_UPDATE` constraint is not present, tags updates are ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

