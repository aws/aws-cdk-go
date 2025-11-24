package mixinsawsroute53recoveryreadiness


// The Amazon Route 53 resource that a DNS target resource record points to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   r53ResourceRecordProperty := &R53ResourceRecordProperty{
//   	DomainName: jsii.String("domainName"),
//   	RecordSetId: jsii.String("recordSetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-r53resourcerecord.html
//
type CfnResourceSetPropsMixin_R53ResourceRecordProperty struct {
	// The DNS target domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-r53resourcerecord.html#cfn-route53recoveryreadiness-resourceset-r53resourcerecord-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Amazon Route 53 Resource Record Set ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-r53resourcerecord.html#cfn-route53recoveryreadiness-resourceset-r53resourcerecord-recordsetid
	//
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
}

