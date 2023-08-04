package awscognito


// Options to create a UserPoolDomain.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   pool.addDomain(jsii.String("CognitoDomain"), &UserPoolDomainOptions{
//   	CognitoDomain: &CognitoDomainOptions{
//   		DomainPrefix: jsii.String("my-awesome-app"),
//   	},
//   })
//
//   certificateArn := "arn:aws:acm:us-east-1:123456789012:certificate/11-3336f1-44483d-adc7-9cd375c5169d"
//
//   domainCert := certificatemanager.Certificate_FromCertificateArn(this, jsii.String("domainCert"), certificateArn)
//   pool.addDomain(jsii.String("CustomDomain"), &UserPoolDomainOptions{
//   	CustomDomain: &CustomDomainOptions{
//   		DomainName: jsii.String("user.myapp.com"),
//   		Certificate: domainCert,
//   	},
//   })
//
type UserPoolDomainOptions struct {
	// Associate a cognito prefix domain with your user pool Either `customDomain` or `cognitoDomain` must be specified.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain-prefix.html
	//
	// Default: - not set if `customDomain` is specified, otherwise, throws an error.
	//
	CognitoDomain *CognitoDomainOptions `field:"optional" json:"cognitoDomain" yaml:"cognitoDomain"`
	// Associate a custom domain with your user pool Either `customDomain` or `cognitoDomain` must be specified.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html
	//
	// Default: - not set if `cognitoDomain` is specified, otherwise, throws an error.
	//
	CustomDomain *CustomDomainOptions `field:"optional" json:"customDomain" yaml:"customDomain"`
}

