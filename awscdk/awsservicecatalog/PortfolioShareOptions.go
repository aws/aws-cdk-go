package awsservicecatalog


// Options for portfolio share.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portfolioShareOptions := &PortfolioShareOptions{
//   	MessageLanguage: awscdk.Aws_servicecatalog.MessageLanguage_EN,
//   	ShareTagOptions: jsii.Boolean(false),
//   }
//
// Experimental.
type PortfolioShareOptions struct {
	// The message language of the share.
	//
	// Controls status and error message language for share.
	// Experimental.
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// Whether to share tagOptions as a part of the portfolio share.
	// Experimental.
	ShareTagOptions *bool `field:"optional" json:"shareTagOptions" yaml:"shareTagOptions"`
}

