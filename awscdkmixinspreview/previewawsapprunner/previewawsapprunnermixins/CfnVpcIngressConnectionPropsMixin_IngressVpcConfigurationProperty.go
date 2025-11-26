package previewawsapprunnermixins


// Specifications for the customer’s VPC and related PrivateLink VPC endpoint that are used to associate with the VPC Ingress Connection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ingressVpcConfigurationProperty := &IngressVpcConfigurationProperty{
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-vpcingressconnection-ingressvpcconfiguration.html
//
type CfnVpcIngressConnectionPropsMixin_IngressVpcConfigurationProperty struct {
	// The ID of the VPC endpoint that your App Runner service connects to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-vpcingressconnection-ingressvpcconfiguration.html#cfn-apprunner-vpcingressconnection-ingressvpcconfiguration-vpcendpointid
	//
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The ID of the VPC that is used for the VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-vpcingressconnection-ingressvpcconfiguration.html#cfn-apprunner-vpcingressconnection-ingressvpcconfiguration-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

