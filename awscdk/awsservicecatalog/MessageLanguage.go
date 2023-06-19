package awsservicecatalog


// The language code.
//
// Used for error and logging messages for end users.
// The default behavior if not specified is English.
//
// Example:
//   servicecatalog.NewPortfolio(this, jsii.String("Portfolio"), &PortfolioProps{
//   	DisplayName: jsii.String("MyFirstPortfolio"),
//   	ProviderName: jsii.String("SCAdmin"),
//   	Description: jsii.String("Portfolio for a project"),
//   	MessageLanguage: servicecatalog.MessageLanguage_EN,
//   })
//
// Experimental.
type MessageLanguage string

const (
	// English.
	// Experimental.
	MessageLanguage_EN MessageLanguage = "EN"
	// Japanese.
	// Experimental.
	MessageLanguage_JP MessageLanguage = "JP"
	// Chinese.
	// Experimental.
	MessageLanguage_ZH MessageLanguage = "ZH"
)

