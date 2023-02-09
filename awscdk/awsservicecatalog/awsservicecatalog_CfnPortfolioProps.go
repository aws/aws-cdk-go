package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPortfolio`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPortfolioProps := &cfnPortfolioProps{
//   	displayName: jsii.String("displayName"),
//   	providerName: jsii.String("providerName"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPortfolioProps struct {
	// The name to use for display purposes.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the portfolio provider.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the portfolio.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// One or more tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

