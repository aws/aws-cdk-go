package previewawsqbusinessmixins


// Contains the authentication configuration details for a data accessor.
//
// This structure defines how the ISV authenticates when accessing data through the data accessor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataAccessorAuthenticationDetailProperty := &DataAccessorAuthenticationDetailProperty{
//   	AuthenticationConfiguration: &DataAccessorAuthenticationConfigurationProperty{
//   		IdcTrustedTokenIssuerConfiguration: &DataAccessorIdcTrustedTokenIssuerConfigurationProperty{
//   			IdcTrustedTokenIssuerArn: jsii.String("idcTrustedTokenIssuerArn"),
//   		},
//   	},
//   	AuthenticationType: jsii.String("authenticationType"),
//   	ExternalIds: []*string{
//   		jsii.String("externalIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-dataaccessorauthenticationdetail.html
//
type CfnDataAccessorPropsMixin_DataAccessorAuthenticationDetailProperty struct {
	// The specific authentication configuration based on the authentication type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-dataaccessorauthenticationdetail.html#cfn-qbusiness-dataaccessor-dataaccessorauthenticationdetail-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// The type of authentication to use for the data accessor.
	//
	// This determines how the ISV authenticates when accessing data. You can use one of two authentication types:
	//
	// - `AWS_IAM_IDC_TTI` - Authentication using IAM Identity Center Trusted Token Issuer (TTI). This authentication type allows the ISV to use a trusted token issuer to generate tokens for accessing the data.
	// - `AWS_IAM_IDC_AUTH_CODE` - Authentication using IAM Identity Center authorization code flow. This authentication type uses the standard OAuth 2.0 authorization code flow for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-dataaccessorauthenticationdetail.html#cfn-qbusiness-dataaccessor-dataaccessorauthenticationdetail-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// A list of external identifiers associated with this authentication configuration.
	//
	// These are used to correlate the data accessor with external systems.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-dataaccessorauthenticationdetail.html#cfn-qbusiness-dataaccessor-dataaccessorauthenticationdetail-externalids
	//
	ExternalIds *[]*string `field:"optional" json:"externalIds" yaml:"externalIds"`
}

