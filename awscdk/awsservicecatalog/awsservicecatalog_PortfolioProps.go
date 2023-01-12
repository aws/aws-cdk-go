package awsservicecatalog


// Properties for a Portfolio.
//
// Example:
//   servicecatalog.NewPortfolio(this, jsii.String("Portfolio"), &portfolioProps{
//   	displayName: jsii.String("MyPortfolio"),
//   	providerName: jsii.String("MyTeam"),
//   })
//
type PortfolioProps struct {
	// The name of the portfolio.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The provider name.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Description for portfolio.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The message language.
	//
	// Controls language for
	// status logging and errors.
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// TagOptions associated directly to a portfolio.
	TagOptions TagOptions `field:"optional" json:"tagOptions" yaml:"tagOptions"`
}

