package previewawselasticloadbalancingv2mixins


// Properties for CfnListenerCertificatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnListenerCertificateMixinProps := &CfnListenerCertificateMixinProps{
//   	Certificates: []interface{}{
//   		&CertificateProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   	},
//   	ListenerArn: jsii.String("listenerArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html
//
type CfnListenerCertificateMixinProps struct {
	// The certificate.
	//
	// You can specify one certificate per resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html#cfn-elasticloadbalancingv2-listenercertificate-certificates
	//
	Certificates interface{} `field:"optional" json:"certificates" yaml:"certificates"`
	// The Amazon Resource Name (ARN) of the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html#cfn-elasticloadbalancingv2-listenercertificate-listenerarn
	//
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
}

