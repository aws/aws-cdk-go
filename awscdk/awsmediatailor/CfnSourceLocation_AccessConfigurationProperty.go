package awsmediatailor


// <p>Access configuration parameters.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessConfigurationProperty := &AccessConfigurationProperty{
//   	AccessType: jsii.String("accessType"),
//   	SecretsManagerAccessTokenConfiguration: &SecretsManagerAccessTokenConfigurationProperty{
//   		HeaderName: jsii.String("headerName"),
//   		SecretArn: jsii.String("secretArn"),
//   		SecretStringKey: jsii.String("secretStringKey"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-accessconfiguration.html
//
type CfnSourceLocation_AccessConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-accessconfiguration.html#cfn-mediatailor-sourcelocation-accessconfiguration-accesstype
	//
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// <p>AWS Secrets Manager access token configuration parameters.
	//
	// For information about Secrets Manager access token authentication, see <a href="https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-access-configuration-access-token.html">Working with AWS Secrets Manager access token authentication</a>.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-accessconfiguration.html#cfn-mediatailor-sourcelocation-accessconfiguration-secretsmanageraccesstokenconfiguration
	//
	SecretsManagerAccessTokenConfiguration interface{} `field:"optional" json:"secretsManagerAccessTokenConfiguration" yaml:"secretsManagerAccessTokenConfiguration"`
}

