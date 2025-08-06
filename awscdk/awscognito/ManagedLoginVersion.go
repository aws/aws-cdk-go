package awscognito


// The branding version of managed login for the domain.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   // Use the new managed login page
//   pool.addDomain(jsii.String("CognitoDomainWithBlandingDesignManagedLogin"), &UserPoolDomainOptions{
//   	CognitoDomain: &CognitoDomainOptions{
//   		DomainPrefix: jsii.String("blanding-design-ui"),
//   	},
//   	ManagedLoginVersion: cognito.ManagedLoginVersion_NEWER_MANAGED_LOGIN,
//   })
//
//   // Use the classic hosted UI
//   pool.addDomain(jsii.String("DomainWithClassicHostedUi"), &UserPoolDomainOptions{
//   	CognitoDomain: &CognitoDomainOptions{
//   		DomainPrefix: jsii.String("classic-hosted-ui"),
//   	},
//   	ManagedLoginVersion: cognito.ManagedLoginVersion_CLASSIC_HOSTED_UI,
//   })
//
type ManagedLoginVersion string

const (
	// The classic hosted UI.
	ManagedLoginVersion_CLASSIC_HOSTED_UI ManagedLoginVersion = "CLASSIC_HOSTED_UI"
	// The newer managed login with the branding designer.
	ManagedLoginVersion_NEWER_MANAGED_LOGIN ManagedLoginVersion = "NEWER_MANAGED_LOGIN"
)

