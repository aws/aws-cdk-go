package awsimagebuilderalpha


// Properties for an EC2 Image Builder AWS Marketplace component.
//
// Example:
//   marketplaceComponent := imagebuilder.AwsMarketplaceComponent_FromAwsMarketplaceComponentAttributes(this, jsii.String("MarketplaceComponent"), &AwsMarketplaceComponentAttributes{
//   	ComponentName: jsii.String("my-marketplace-component"),
//   	MarketplaceProductId: jsii.String("prod-1234567890abcdef0"),
//   })
//
// Experimental.
type AwsMarketplaceComponentAttributes struct {
	// The name of the AWS Marketplace component.
	//
	// This name should exclude the marketplace product ID from it.
	// Experimental.
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// The marketplace product ID associated with the component.
	// Experimental.
	MarketplaceProductId *string `field:"required" json:"marketplaceProductId" yaml:"marketplaceProductId"`
	// The version of the AWS Marketplace component.
	// Default: - the latest version of the component, x.x.x
	//
	// Experimental.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
}

