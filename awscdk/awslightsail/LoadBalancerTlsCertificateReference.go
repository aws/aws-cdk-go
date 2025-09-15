package awslightsail


// A reference to a LoadBalancerTlsCertificate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerTlsCertificateReference := &LoadBalancerTlsCertificateReference{
//   	CertificateName: jsii.String("certificateName"),
//   	LoadBalancerName: jsii.String("loadBalancerName"),
//   	LoadBalancerTlsCertificateArn: jsii.String("loadBalancerTlsCertificateArn"),
//   }
//
type LoadBalancerTlsCertificateReference struct {
	// The CertificateName of the LoadBalancerTlsCertificate resource.
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// The LoadBalancerName of the LoadBalancerTlsCertificate resource.
	LoadBalancerName *string `field:"required" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The ARN of the LoadBalancerTlsCertificate resource.
	LoadBalancerTlsCertificateArn *string `field:"required" json:"loadBalancerTlsCertificateArn" yaml:"loadBalancerTlsCertificateArn"`
}

