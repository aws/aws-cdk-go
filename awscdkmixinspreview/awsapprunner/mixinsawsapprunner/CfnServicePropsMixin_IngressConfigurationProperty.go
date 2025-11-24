package mixinsawsapprunner


// Network configuration settings for inbound network traffic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ingressConfigurationProperty := &IngressConfigurationProperty{
//   	IsPubliclyAccessible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-ingressconfiguration.html
//
type CfnServicePropsMixin_IngressConfigurationProperty struct {
	// Specifies whether your App Runner service is publicly accessible.
	//
	// To make the service publicly accessible set it to `True` . To make the service privately accessible, from only within an Amazon VPC set it to `False` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-ingressconfiguration.html#cfn-apprunner-service-ingressconfiguration-ispubliclyaccessible
	//
	IsPubliclyAccessible interface{} `field:"optional" json:"isPubliclyAccessible" yaml:"isPubliclyAccessible"`
}

