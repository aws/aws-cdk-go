package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainNameConfigurationProperty := &domainNameConfigurationProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	certificateName: jsii.String("certificateName"),
//   	endpointType: jsii.String("endpointType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnDomainNameV2_DomainNameConfigurationProperty struct {
	// `CfnDomainNameV2.DomainNameConfigurationProperty.CertificateArn`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html#cfn-apigatewayv2-domainname-domainnameconfiguration-certificatearn
	//
	// Deprecated: moved to package aws-apigatewayv2.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// `CfnDomainNameV2.DomainNameConfigurationProperty.CertificateName`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html#cfn-apigatewayv2-domainname-domainnameconfiguration-certificatename
	//
	// Deprecated: moved to package aws-apigatewayv2.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// `CfnDomainNameV2.DomainNameConfigurationProperty.EndpointType`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html#cfn-apigatewayv2-domainname-domainnameconfiguration-endpointtype
	//
	// Deprecated: moved to package aws-apigatewayv2.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
}

