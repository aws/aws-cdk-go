package awsnetworkfirewall


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetReferenceProperty := map[string]*string{
//   	"referenceArn": jsii.String("referenceArn"),
//   }
//
type CfnRuleGroup_IPSetReferenceProperty struct {
	// `CfnRuleGroup.IPSetReferenceProperty.ReferenceArn`.
	ReferenceArn *string `field:"optional" json:"referenceArn" yaml:"referenceArn"`
}

