package awsverifiedpermissions


// A structure that contains configuration of the identity source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identitySourceDetailsProperty := &IdentitySourceDetailsProperty{
//   	ClientIds: []*string{
//   		jsii.String("clientIds"),
//   	},
//   	DiscoveryUrl: jsii.String("discoveryUrl"),
//   	OpenIdIssuer: jsii.String("openIdIssuer"),
//   	UserPoolArn: jsii.String("userPoolArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourcedetails.html
//
type CfnIdentitySource_IdentitySourceDetailsProperty struct {
	// The application client IDs associated with the specified Amazon Cognito user pool that are enabled for this identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourcedetails.html#cfn-verifiedpermissions-identitysource-identitysourcedetails-clientids
	//
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// The well-known URL that points to this user pool's OIDC discovery endpoint.
	//
	// This is a URL string in the following format. This URL replaces the placeholders for both the AWS Region and the user pool identifier with those appropriate for this user pool.
	//
	// `https://cognito-idp. *<region>* .amazonaws.com/ *<user-pool-id>* /.well-known/openid-configuration`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourcedetails.html#cfn-verifiedpermissions-identitysource-identitysourcedetails-discoveryurl
	//
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
	// A string that identifies the type of OIDC service represented by this identity source.
	//
	// At this time, the only valid value is `cognito` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourcedetails.html#cfn-verifiedpermissions-identitysource-identitysourcedetails-openidissuer
	//
	OpenIdIssuer *string `field:"optional" json:"openIdIssuer" yaml:"openIdIssuer"`
	// The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of the Amazon Cognito user pool whose identities are accessible to this Verified Permissions policy store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourcedetails.html#cfn-verifiedpermissions-identitysource-identitysourcedetails-userpoolarn
	//
	UserPoolArn *string `field:"optional" json:"userPoolArn" yaml:"userPoolArn"`
}

