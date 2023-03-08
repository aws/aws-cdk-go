package awscognito


// Props for UserPoolDomain construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var certificate certificate
//   var userPool userPool
//
//   userPoolDomainProps := &UserPoolDomainProps{
//   	UserPool: userPool,
//
//   	// the properties below are optional
//   	CognitoDomain: &CognitoDomainOptions{
//   		DomainPrefix: jsii.String("domainPrefix"),
//   	},
//   	CustomDomain: &CustomDomainOptions{
//   		Certificate: certificate,
//   		DomainName: jsii.String("domainName"),
//   	},
//   }
//
type UserPoolDomainProps struct {
	// Associate a cognito prefix domain with your user pool Either `customDomain` or `cognitoDomain` must be specified.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain-prefix.html
	//
	CognitoDomain *CognitoDomainOptions `field:"optional" json:"cognitoDomain" yaml:"cognitoDomain"`
	// Associate a custom domain with your user pool Either `customDomain` or `cognitoDomain` must be specified.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html
	//
	CustomDomain *CustomDomainOptions `field:"optional" json:"customDomain" yaml:"customDomain"`
	// The user pool to which this domain should be associated.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
}

