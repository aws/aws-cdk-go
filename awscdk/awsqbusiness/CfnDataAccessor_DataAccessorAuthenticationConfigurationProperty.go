package awsqbusiness


// A union type that contains the specific authentication configuration based on the authentication type selected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataAccessorAuthenticationConfigurationProperty := &DataAccessorAuthenticationConfigurationProperty{
//   	IdcTrustedTokenIssuerConfiguration: &DataAccessorIdcTrustedTokenIssuerConfigurationProperty{
//   		IdcTrustedTokenIssuerArn: jsii.String("idcTrustedTokenIssuerArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-dataaccessorauthenticationconfiguration.html
//
type CfnDataAccessor_DataAccessorAuthenticationConfigurationProperty struct {
	// Configuration for IAM Identity Center Trusted Token Issuer (TTI) authentication used when the authentication type is `AWS_IAM_IDC_TTI` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-dataaccessorauthenticationconfiguration.html#cfn-qbusiness-dataaccessor-dataaccessorauthenticationconfiguration-idctrustedtokenissuerconfiguration
	//
	IdcTrustedTokenIssuerConfiguration interface{} `field:"required" json:"idcTrustedTokenIssuerConfiguration" yaml:"idcTrustedTokenIssuerConfiguration"`
}

