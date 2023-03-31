package awselasticloadbalancingv2


// Properties for defining a `CfnListenerCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerCertificateProps := &cfnListenerCertificateProps{
//   	certificates: []interface{}{
//   		&certificateProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   		},
//   	},
//   	listenerArn: jsii.String("listenerArn"),
//   }
//
type CfnListenerCertificateProps struct {
	// The certificate.
	//
	// You can specify one certificate per resource.
	Certificates interface{} `field:"required" json:"certificates" yaml:"certificates"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
}

