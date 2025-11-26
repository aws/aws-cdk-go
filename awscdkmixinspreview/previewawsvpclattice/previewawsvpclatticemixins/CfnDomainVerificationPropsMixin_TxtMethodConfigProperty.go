package previewawsvpclatticemixins


// Configuration for TXT record-based domain verification method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   txtMethodConfigProperty := &TxtMethodConfigProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-domainverification-txtmethodconfig.html
//
type CfnDomainVerificationPropsMixin_TxtMethodConfigProperty struct {
	// The name of the TXT record that must be created for domain verification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-domainverification-txtmethodconfig.html#cfn-vpclattice-domainverification-txtmethodconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value that must be added to the TXT record for domain verification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-domainverification-txtmethodconfig.html#cfn-vpclattice-domainverification-txtmethodconfig-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

