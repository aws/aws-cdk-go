package awscognito


// Options while specifying a cognito prefix domain.
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
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain-prefix.html
//
type CognitoDomainOptions struct {
	// The prefix to the Cognito hosted domain name that will be associated with the user pool.
	DomainPrefix *string `field:"required" json:"domainPrefix" yaml:"domainPrefix"`
}

