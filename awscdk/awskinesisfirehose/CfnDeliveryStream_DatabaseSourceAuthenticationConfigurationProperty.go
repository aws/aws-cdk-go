package awskinesisfirehose


// The structure to configure the authentication methods for Firehose to connect to source database endpoint.
//
// Amazon Data Firehose is in preview release and is subject to change.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseSourceAuthenticationConfigurationProperty := &DatabaseSourceAuthenticationConfigurationProperty{
//   	SecretsManagerConfiguration: &SecretsManagerConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceauthenticationconfiguration.html
//
type CfnDeliveryStream_DatabaseSourceAuthenticationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceauthenticationconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceauthenticationconfiguration-secretsmanagerconfiguration
	//
	SecretsManagerConfiguration interface{} `field:"required" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
}

