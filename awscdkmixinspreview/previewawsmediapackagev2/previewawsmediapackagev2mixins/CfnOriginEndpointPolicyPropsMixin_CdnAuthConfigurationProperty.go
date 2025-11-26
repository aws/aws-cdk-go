package previewawsmediapackagev2mixins


// The settings to enable CDN authorization headers in MediaPackage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cdnAuthConfigurationProperty := &CdnAuthConfigurationProperty{
//   	CdnIdentifierSecretArns: []*string{
//   		jsii.String("cdnIdentifierSecretArns"),
//   	},
//   	SecretsRoleArn: jsii.String("secretsRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpointpolicy-cdnauthconfiguration.html
//
type CfnOriginEndpointPolicyPropsMixin_CdnAuthConfigurationProperty struct {
	// The ARN for the secret in Secrets Manager that your CDN uses for authorization to access the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpointpolicy-cdnauthconfiguration.html#cfn-mediapackagev2-originendpointpolicy-cdnauthconfiguration-cdnidentifiersecretarns
	//
	CdnIdentifierSecretArns *[]*string `field:"optional" json:"cdnIdentifierSecretArns" yaml:"cdnIdentifierSecretArns"`
	// The ARN for the IAM role that gives MediaPackage read access to Secrets Manager and AWS  for CDN authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpointpolicy-cdnauthconfiguration.html#cfn-mediapackagev2-originendpointpolicy-cdnauthconfiguration-secretsrolearn
	//
	SecretsRoleArn *string `field:"optional" json:"secretsRoleArn" yaml:"secretsRoleArn"`
}

