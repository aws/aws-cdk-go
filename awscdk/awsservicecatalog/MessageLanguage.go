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
type MessageLanguage string

const (
	// English.
	MessageLanguage_EN MessageLanguage = "EN"
	// Japanese.
	MessageLanguage_JP MessageLanguage = "JP"
	// Chinese.
	MessageLanguage_ZH MessageLanguage = "ZH"
)

