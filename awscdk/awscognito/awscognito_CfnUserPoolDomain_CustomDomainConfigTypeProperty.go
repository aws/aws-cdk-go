package awscognito


// The configuration for a custom domain that hosts the sign-up and sign-in webpages for your application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customDomainConfigTypeProperty := &customDomainConfigTypeProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   }
//
type CfnUserPoolDomain_CustomDomainConfigTypeProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Certificate Manager SSL certificate.
	//
	// You use this certificate for the subdomain of your custom domain.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
}

