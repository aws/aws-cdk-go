package previewawsroute53resolvermixins


// Properties for CfnResolverDNSSECConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResolverDNSSECConfigMixinProps := &CfnResolverDNSSECConfigMixinProps{
//   	ResourceId: jsii.String("resourceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverdnssecconfig.html
//
type CfnResolverDNSSECConfigMixinProps struct {
	// The ID of the virtual private cloud (VPC) that you're configuring the DNSSEC validation status for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverdnssecconfig.html#cfn-route53resolver-resolverdnssecconfig-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

