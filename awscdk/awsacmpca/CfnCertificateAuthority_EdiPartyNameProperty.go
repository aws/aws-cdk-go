package awsacmpca


// Describes an Electronic Data Interchange (EDI) entity as described in as defined in [Subject Alternative Name](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) in RFC 5280.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ediPartyNameProperty := &EdiPartyNameProperty{
//   	PartyName: jsii.String("partyName"),
//
//   	// the properties below are optional
//   	NameAssigner: jsii.String("nameAssigner"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-edipartyname.html
//
type CfnCertificateAuthority_EdiPartyNameProperty struct {
	// Specifies the party name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-edipartyname.html#cfn-acmpca-certificateauthority-edipartyname-partyname
	//
	PartyName *string `field:"required" json:"partyName" yaml:"partyName"`
	// Specifies the name assigner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-edipartyname.html#cfn-acmpca-certificateauthority-edipartyname-nameassigner
	//
	NameAssigner *string `field:"optional" json:"nameAssigner" yaml:"nameAssigner"`
}

