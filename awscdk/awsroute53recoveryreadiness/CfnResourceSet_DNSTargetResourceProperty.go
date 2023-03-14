package awsroute53recoveryreadiness


// A component for DNS/routing control readiness checks and architecture checks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dNSTargetResourceProperty := &DNSTargetResourceProperty{
//   	DomainName: jsii.String("domainName"),
//   	HostedZoneArn: jsii.String("hostedZoneArn"),
//   	RecordSetId: jsii.String("recordSetId"),
//   	RecordType: jsii.String("recordType"),
//   	TargetResource: &TargetResourceProperty{
//   		NlbResource: &NLBResourceProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		R53Resource: &R53ResourceRecordProperty{
//   			DomainName: jsii.String("domainName"),
//   			RecordSetId: jsii.String("recordSetId"),
//   		},
//   	},
//   }
//
type CfnResourceSet_DNSTargetResourceProperty struct {
	// The domain name that acts as an ingress point to a portion of the customer application.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.
	HostedZoneArn *string `field:"optional" json:"hostedZoneArn" yaml:"hostedZoneArn"`
	// The Amazon Route 53 record set ID that uniquely identifies a DNS record, given a name and a type.
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
	// The type of DNS record of the target resource.
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// The target resource that the Route 53 record points to.
	TargetResource interface{} `field:"optional" json:"targetResource" yaml:"targetResource"`
}

