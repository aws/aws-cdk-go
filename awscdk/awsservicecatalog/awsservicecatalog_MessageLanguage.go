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

