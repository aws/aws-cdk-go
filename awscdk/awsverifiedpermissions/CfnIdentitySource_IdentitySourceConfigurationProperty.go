package awsverifiedpermissions


// A structure that contains configuration information used when creating or updating a new identity source.
//
// > At this time, the only valid member of this structure is a Amazon Cognito user pool configuration.
// >
// > You must specify a `userPoolArn` , and optionally, a `ClientId` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identitySourceConfigurationProperty := &IdentitySourceConfigurationProperty{
//   	CognitoUserPoolConfiguration: &CognitoUserPoolConfigurationProperty{
//   		UserPoolArn: jsii.String("userPoolArn"),
//
//   		// the properties below are optional
//   		ClientIds: []*string{
//   			jsii.String("clientIds"),
//   		},
//   		GroupConfiguration: &CognitoGroupConfigurationProperty{
//   			GroupEntityType: jsii.String("groupEntityType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourceconfiguration.html
//
type CfnIdentitySource_IdentitySourceConfigurationProperty struct {
	// A structure that contains configuration information used when creating or updating an identity source that represents a connection to an Amazon Cognito user pool used as an identity provider for Verified Permissions .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-identitysourceconfiguration.html#cfn-verifiedpermissions-identitysource-identitysourceconfiguration-cognitouserpoolconfiguration
	//
	CognitoUserPoolConfiguration interface{} `field:"required" json:"cognitoUserPoolConfiguration" yaml:"cognitoUserPoolConfiguration"`
}

