package awsnetworkfirewall


// Configures one or more `IPSetReferences` for a Suricata-compatible rule group.
//
// An IP set reference is a rule variable that references a resource that you create and manage in another AWS service, such as an Amazon VPC prefix list. Network Firewall IP set references enable you to dynamically update the contents of your rules. When you create, update, or delete the IP set you are referencing in your rule, Network Firewall automatically updates the rule's content with the changes. For more information about IP set references in Network Firewall , see [Using IP set references](https://docs.aws.amazon.com/network-firewall/latest/developerguide/rule-groups-ip-set-references.html) in the *Network Firewall Developer Guide* .
//
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
	// The Amazon Resource Name (ARN) of the resource to include in the `RuleGroup.IPSetReference` .
	ReferenceArn *string `field:"optional" json:"referenceArn" yaml:"referenceArn"`
}

