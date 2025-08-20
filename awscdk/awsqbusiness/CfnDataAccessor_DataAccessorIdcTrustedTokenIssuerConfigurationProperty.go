package awsqbusiness


// Configuration details for IAM Identity Center Trusted Token Issuer (TTI) authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataAccessorIdcTrustedTokenIssuerConfigurationProperty := &DataAccessorIdcTrustedTokenIssuerConfigurationProperty{
//   	IdcTrustedTokenIssuerArn: jsii.String("idcTrustedTokenIssuerArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-dataaccessoridctrustedtokenissuerconfiguration.html
//
type CfnDataAccessor_DataAccessorIdcTrustedTokenIssuerConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the IAM Identity Center Trusted Token Issuer that will be used for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-dataaccessoridctrustedtokenissuerconfiguration.html#cfn-qbusiness-dataaccessor-dataaccessoridctrustedtokenissuerconfiguration-idctrustedtokenissuerarn
	//
	IdcTrustedTokenIssuerArn *string `field:"required" json:"idcTrustedTokenIssuerArn" yaml:"idcTrustedTokenIssuerArn"`
}

