package mixinsawsverifiedpermissions


// A structure that contains configuration information used when creating or updating an identity source that represents a connection to an Amazon Cognito user pool used as an identity provider for Verified Permissions .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cognitoUserPoolConfigurationProperty := &CognitoUserPoolConfigurationProperty{
//   	ClientIds: []*string{
//   		jsii.String("clientIds"),
//   	},
//   	GroupConfiguration: &CognitoGroupConfigurationProperty{
//   		GroupEntityType: jsii.String("groupEntityType"),
//   	},
//   	UserPoolArn: jsii.String("userPoolArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-cognitouserpoolconfiguration.html
//
type CfnIdentitySourcePropsMixin_CognitoUserPoolConfigurationProperty struct {
	// The unique application client IDs that are associated with the specified Amazon Cognito user pool.
	//
	// Example: `"ClientIds": ["&ExampleCogClientId;"]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-cognitouserpoolconfiguration.html#cfn-verifiedpermissions-identitysource-cognitouserpoolconfiguration-clientids
	//
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// The type of entity that a policy store maps to groups from an Amazon Cognito user pool identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-cognitouserpoolconfiguration.html#cfn-verifiedpermissions-identitysource-cognitouserpoolconfiguration-groupconfiguration
	//
	GroupConfiguration interface{} `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
	// The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of the Amazon Cognito user pool that contains the identities to be authorized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-cognitouserpoolconfiguration.html#cfn-verifiedpermissions-identitysource-cognitouserpoolconfiguration-userpoolarn
	//
	UserPoolArn *string `field:"optional" json:"userPoolArn" yaml:"userPoolArn"`
}

