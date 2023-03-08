package awsacmpca


// Describes an Electronic Data Interchange (EDI) entity as described in as defined in [Subject Alternative Name](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) in RFC 5280.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ediPartyNameProperty := &ediPartyNameProperty{
//   	nameAssigner: jsii.String("nameAssigner"),
//   	partyName: jsii.String("partyName"),
//   }
//
type CfnCertificateAuthority_EdiPartyNameProperty struct {
	// Specifies the name assigner.
	NameAssigner *string `field:"required" json:"nameAssigner" yaml:"nameAssigner"`
	// Specifies the party name.
	PartyName *string `field:"required" json:"partyName" yaml:"partyName"`
}

