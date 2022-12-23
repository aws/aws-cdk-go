package awscognito


// Options to customize the behaviour of `signInUrl()`.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("UserPool"), &userPoolProps{
//   })
//   client := userpool.addClient(jsii.String("Client"), &userPoolClientOptions{
//   	// ...
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			implicitCodeGrant: jsii.Boolean(true),
//   		},
//   		callbackUrls: []*string{
//   			jsii.String("https://myapp.com/home"),
//   			jsii.String("https://myapp.com/users"),
//   		},
//   	},
//   })
//   domain := userpool.addDomain(jsii.String("Domain"), &userPoolDomainOptions{
//   })
//   signInUrl := domain.signInUrl(client, &signInUrlOptions{
//   	redirectUri: jsii.String("https://myapp.com/home"),
//   })
//
type SignInUrlOptions struct {
	// Whether to return the FIPS-compliant endpoint.
	Fips *bool `field:"optional" json:"fips" yaml:"fips"`
	// Where to redirect to after sign in.
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
	// The path in the URI where the sign-in page is located.
	SignInPath *string `field:"optional" json:"signInPath" yaml:"signInPath"`
}

