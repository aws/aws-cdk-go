package mixinsawsroute53recoveryreadiness


// A component for DNS/routing control readiness checks and architecture checks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html
//
type CfnResourceSetPropsMixin_DNSTargetResourceProperty struct {
	// The domain name that acts as an ingress point to a portion of the customer application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-hostedzonearn
	//
	HostedZoneArn *string `field:"optional" json:"hostedZoneArn" yaml:"hostedZoneArn"`
	// The Amazon Route 53 record set ID that uniquely identifies a DNS record, given a name and a type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-recordsetid
	//
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
	// The type of DNS record of the target resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-recordtype
	//
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// The target resource that the Route 53 record points to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-targetresource
	//
	TargetResource interface{} `field:"optional" json:"targetResource" yaml:"targetResource"`
}

