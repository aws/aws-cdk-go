package awscertificatemanager


// The expiration configuration for the external account binding.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   expirationProperty := &ExpirationProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeexternalaccountbinding-expiration.html
//
type CfnAcmeExternalAccountBindingPropsMixin_ExpirationProperty struct {
	// The time unit for the expiration value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeexternalaccountbinding-expiration.html#cfn-certificatemanager-acmeexternalaccountbinding-expiration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The expiration value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeexternalaccountbinding-expiration.html#cfn-certificatemanager-acmeexternalaccountbinding-expiration-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

