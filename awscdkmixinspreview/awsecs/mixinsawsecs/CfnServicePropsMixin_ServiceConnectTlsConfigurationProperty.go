package mixinsawsecs


// The key that encrypts and decrypts your resources for Service Connect TLS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceConnectTlsConfigurationProperty := &ServiceConnectTlsConfigurationProperty{
//   	IssuerCertificateAuthority: &ServiceConnectTlsCertificateAuthorityProperty{
//   		AwsPcaAuthorityArn: jsii.String("awsPcaAuthorityArn"),
//   	},
//   	KmsKey: jsii.String("kmsKey"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttlsconfiguration.html
//
type CfnServicePropsMixin_ServiceConnectTlsConfigurationProperty struct {
	// The signer certificate authority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttlsconfiguration.html#cfn-ecs-service-serviceconnecttlsconfiguration-issuercertificateauthority
	//
	IssuerCertificateAuthority interface{} `field:"optional" json:"issuerCertificateAuthority" yaml:"issuerCertificateAuthority"`
	// The AWS Key Management Service key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttlsconfiguration.html#cfn-ecs-service-serviceconnecttlsconfiguration-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The Amazon Resource Name (ARN) of the IAM role that's associated with the Service Connect TLS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttlsconfiguration.html#cfn-ecs-service-serviceconnecttlsconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

