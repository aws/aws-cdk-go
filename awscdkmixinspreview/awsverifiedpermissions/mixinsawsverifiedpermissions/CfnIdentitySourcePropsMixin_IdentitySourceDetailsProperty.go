package mixinsawsverifiedpermissions


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnIdentitySourcePropsMixin_IdentitySourceDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourcedetails.html#cfn-verifiedpermissions-identitysource-identitysourcedetails-clientids
	//
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourcedetails.html#cfn-verifiedpermissions-identitysource-identitysourcedetails-discoveryurl
	//
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourcedetails.html#cfn-verifiedpermissions-identitysource-identitysourcedetails-openidissuer
	//
	OpenIdIssuer *string `field:"optional" json:"openIdIssuer" yaml:"openIdIssuer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourcedetails.html#cfn-verifiedpermissions-identitysource-identitysourcedetails-userpoolarn
	//
	UserPoolArn *string `field:"optional" json:"userPoolArn" yaml:"userPoolArn"`
}

