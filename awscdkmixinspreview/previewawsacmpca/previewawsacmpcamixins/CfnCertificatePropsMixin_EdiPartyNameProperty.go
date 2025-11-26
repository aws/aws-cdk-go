package previewawsacmpcamixins


// Describes an Electronic Data Interchange (EDI) entity as described in as defined in [Subject Alternative Name](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) in RFC 5280.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ediPartyNameProperty := &EdiPartyNameProperty{
//   	NameAssigner: jsii.String("nameAssigner"),
//   	PartyName: jsii.String("partyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-edipartyname.html
//
type CfnCertificatePropsMixin_EdiPartyNameProperty struct {
	// Specifies the name assigner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-edipartyname.html#cfn-acmpca-certificate-edipartyname-nameassigner
	//
	NameAssigner *string `field:"optional" json:"nameAssigner" yaml:"nameAssigner"`
	// Specifies the party name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-edipartyname.html#cfn-acmpca-certificate-edipartyname-partyname
	//
	PartyName *string `field:"optional" json:"partyName" yaml:"partyName"`
}

