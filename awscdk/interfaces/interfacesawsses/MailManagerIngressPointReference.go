package interfacesawsses


// A reference to a MailManagerIngressPoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerIngressPointReference := &MailManagerIngressPointReference{
//   	IngressPointArn: jsii.String("ingressPointArn"),
//   	IngressPointId: jsii.String("ingressPointId"),
//   }
//
type MailManagerIngressPointReference struct {
	// The ARN of the MailManagerIngressPoint resource.
	IngressPointArn *string `field:"required" json:"ingressPointArn" yaml:"ingressPointArn"`
	// The IngressPointId of the MailManagerIngressPoint resource.
	IngressPointId *string `field:"required" json:"ingressPointId" yaml:"ingressPointId"`
}

