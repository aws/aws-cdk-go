package awsservicecatalog


// The language code.
//
// Used for error and logging messages for end users.
// The default behavior if not specified is English.
//
// Example:
//   servicecatalog.NewPortfolio(this, jsii.String("Portfolio"), &portfolioProps{
//   	displayName: jsii.String("MyFirstPortfolio"),
//   	providerName: jsii.String("SCAdmin"),
//   	description: jsii.String("Portfolio for a project"),
//   	messageLanguage: servicecatalog.messageLanguage_EN,
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

