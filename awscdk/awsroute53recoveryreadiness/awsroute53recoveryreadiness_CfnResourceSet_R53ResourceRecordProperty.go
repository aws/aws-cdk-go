package awsroute53recoveryreadiness


// The Route 53 resource that a DNS target resource record points to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   r53ResourceRecordProperty := &r53ResourceRecordProperty{
//   	domainName: jsii.String("domainName"),
//   	recordSetId: jsii.String("recordSetId"),
//   }
//
type CfnResourceSet_R53ResourceRecordProperty struct {
	// The DNS target domain name.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Route 53 Resource Record Set ID.
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
}

