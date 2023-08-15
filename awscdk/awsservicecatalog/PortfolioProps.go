package awsservicecatalog


// Properties for a Portfolio.
//
// Example:
//   servicecatalog.NewPortfolio(this, jsii.String("Portfolio"), &PortfolioProps{
//   	DisplayName: jsii.String("MyPortfolio"),
//   	ProviderName: jsii.String("MyTeam"),
//   })
//
type PortfolioProps struct {
	// The name of the portfolio.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The provider name.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Description for portfolio.
	// Default: - No description provided.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The message language.
	//
	// Controls language for
	// status logging and errors.
	// Default: - English.
	//
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// TagOptions associated directly to a portfolio.
	// Default: - No tagOptions provided.
	//
	TagOptions TagOptions `field:"optional" json:"tagOptions" yaml:"tagOptions"`
}

