package awscognito


// Options to customize the behaviour of `signInUrl()`.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("UserPool"), &UserPoolProps{
//   })
//   client := userpool.addClient(jsii.String("Client"), &UserPoolClientOptions{
//   	// ...
//   	OAuth: &OAuthSettings{
//   		Flows: &OAuthFlows{
//   			ImplicitCodeGrant: jsii.Boolean(true),
//   		},
//   		CallbackUrls: []*string{
//   			jsii.String("https://myapp.com/home"),
//   			jsii.String("https://myapp.com/users"),
//   		},
//   	},
//   })
//   domain := userpool.addDomain(jsii.String("Domain"), &UserPoolDomainOptions{
//   })
//   signInUrl := domain.SignInUrl(client, &SignInUrlOptions{
//   	RedirectUri: jsii.String("https://myapp.com/home"),
//   })
//
type SignInUrlOptions struct {
	// Whether to return the FIPS-compliant endpoint.
	// Default: return the standard URL.
	//
	Fips *bool `field:"optional" json:"fips" yaml:"fips"`
	// Where to redirect to after sign in.
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
	// The path in the URI where the sign-in page is located.
	// Default: '/login'.
	//
	SignInPath *string `field:"optional" json:"signInPath" yaml:"signInPath"`
}

