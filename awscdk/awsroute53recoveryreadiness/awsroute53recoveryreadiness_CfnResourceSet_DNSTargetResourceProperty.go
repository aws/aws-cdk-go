package awsroute53recoveryreadiness


// A component for DNS/routing control readiness checks and architecture checks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dNSTargetResourceProperty := &dNSTargetResourceProperty{
//   	domainName: jsii.String("domainName"),
//   	hostedZoneArn: jsii.String("hostedZoneArn"),
//   	recordSetId: jsii.String("recordSetId"),
//   	recordType: jsii.String("recordType"),
//   	targetResource: &targetResourceProperty{
//   		nlbResource: &nLBResourceProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		r53Resource: &r53ResourceRecordProperty{
//   			domainName: jsii.String("domainName"),
//   			recordSetId: jsii.String("recordSetId"),
//   		},
//   	},
//   }
//
type CfnResourceSet_DNSTargetResourceProperty struct {
	// The domain name that acts as an ingress point to a portion of the customer application.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.
	HostedZoneArn *string `field:"optional" json:"hostedZoneArn" yaml:"hostedZoneArn"`
	// The Route 53 record set ID that uniquely identifies a DNS record, given a name and a type.
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
	// The type of DNS record of the target resource.
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// The target resource that the Route 53 record points to.
	TargetResource interface{} `field:"optional" json:"targetResource" yaml:"targetResource"`
}

