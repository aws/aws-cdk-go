package awsredshift


// The authorized token issuer for the Amazon Redshift IAM Identity Center application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   authorizedTokenIssuerProperty := &AuthorizedTokenIssuerProperty{
//   	AuthorizedAudiencesList: []*string{
//   		jsii.String("authorizedAudiencesList"),
//   	},
//   	TrustedTokenIssuerArn: jsii.String("trustedTokenIssuerArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-authorizedtokenissuer.html
//
type CfnRedshiftIdcApplicationPropsMixin_AuthorizedTokenIssuerProperty struct {
	// The list of audiences for the authorized token issuer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-authorizedtokenissuer.html#cfn-redshift-redshiftidcapplication-authorizedtokenissuer-authorizedaudienceslist
	//
	AuthorizedAudiencesList *[]*string `field:"optional" json:"authorizedAudiencesList" yaml:"authorizedAudiencesList"`
	// The ARN for the authorized token issuer for integrating Amazon Redshift with IDC Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-authorizedtokenissuer.html#cfn-redshift-redshiftidcapplication-authorizedtokenissuer-trustedtokenissuerarn
	//
	TrustedTokenIssuerArn *string `field:"optional" json:"trustedTokenIssuerArn" yaml:"trustedTokenIssuerArn"`
}

