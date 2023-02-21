package awsnetworkfirewall


// Configures the `ReferenceSets` for a stateful rule group.
//
// For more information, see the [Using IP set references in Suricata compatible rule groups](https://docs.aws.amazon.com/network-firewall/latest/developerguide/rule-groups-ip-set-references.html) in the *Network Firewall User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceSetsProperty := &referenceSetsProperty{
//   	ipSetReferences: map[string]interface{}{
//   		"ipSetReferencesKey": map[string]*string{
//   			"referenceArn": jsii.String("referenceArn"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_ReferenceSetsProperty struct {
	// The IP set references to use in the stateful rule group.
	IpSetReferences interface{} `field:"optional" json:"ipSetReferences" yaml:"ipSetReferences"`
}

