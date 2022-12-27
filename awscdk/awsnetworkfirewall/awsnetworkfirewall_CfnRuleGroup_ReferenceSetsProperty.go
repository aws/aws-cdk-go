package awsnetworkfirewall


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
	// `CfnRuleGroup.ReferenceSetsProperty.IPSetReferences`.
	IpSetReferences interface{} `field:"optional" json:"ipSetReferences" yaml:"ipSetReferences"`
}

