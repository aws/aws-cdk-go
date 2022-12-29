package awsacmpca


// Specifies the X.509 extension information for a certificate.
//
// Extensions present in `CustomExtensions` follow the `ApiPassthrough` [template rules](https://docs.aws.amazon.com/acm-pca/latest/userguide/UsingTemplates.html#template-order-of-operations) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customExtensionProperty := &customExtensionProperty{
//   	objectIdentifier: jsii.String("objectIdentifier"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	critical: jsii.Boolean(false),
//   }
//
type CfnCertificate_CustomExtensionProperty struct {
	// Specifies the object identifier (OID) of the X.509 extension. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29).
	ObjectIdentifier *string `field:"required" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Specifies the base64-encoded value of the X.509 extension.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Specifies the critical flag of the X.509 extension.
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
}

