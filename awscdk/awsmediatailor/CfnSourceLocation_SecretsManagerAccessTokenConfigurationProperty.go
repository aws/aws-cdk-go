package awsmediatailor


// AWS Secrets Manager access token configuration parameters.
//
// For information about Secrets Manager access token authentication, see [Working with AWS Secrets Manager access token authentication](https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-access-configuration-access-token.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretsManagerAccessTokenConfigurationProperty := &SecretsManagerAccessTokenConfigurationProperty{
//   	HeaderName: jsii.String("headerName"),
//   	SecretArn: jsii.String("secretArn"),
//   	SecretStringKey: jsii.String("secretStringKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration.html
//
type CfnSourceLocation_SecretsManagerAccessTokenConfigurationProperty struct {
	// The name of the HTTP header used to supply the access token in requests to the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration.html#cfn-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration-headername
	//
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration.html#cfn-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The AWS Secrets Manager [SecretString](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html#SecretsManager-CreateSecret-request-SecretString.html) key associated with the access token. MediaTailor uses the key to look up SecretString key and value pair containing the access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration.html#cfn-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration-secretstringkey
	//
	SecretStringKey *string `field:"optional" json:"secretStringKey" yaml:"secretStringKey"`
}

