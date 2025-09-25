package awsverifiedpermissions


// The type of entity that a policy store maps to groups from an Amazon Cognito user pool identity source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoGroupConfigurationProperty := &CognitoGroupConfigurationProperty{
//   	GroupEntityType: jsii.String("groupEntityType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-cognitogroupconfiguration.html
//
type CfnIdentitySource_CognitoGroupConfigurationProperty struct {
	// The name of the schema entity type that's mapped to the user pool group.
	//
	// Defaults to `AWS::CognitoGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-cognitogroupconfiguration.html#cfn-verifiedpermissions-identitysource-cognitogroupconfiguration-groupentitytype
	//
	GroupEntityType *string `field:"required" json:"groupEntityType" yaml:"groupEntityType"`
}

