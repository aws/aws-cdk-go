package mixinsawsroute53recoveryreadiness


// The target resource that the Route 53 record points to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetResourceProperty := &TargetResourceProperty{
//   	NlbResource: &NLBResourceProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	R53Resource: &R53ResourceRecordProperty{
//   		DomainName: jsii.String("domainName"),
//   		RecordSetId: jsii.String("recordSetId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-targetresource.html
//
type CfnResourceSetPropsMixin_TargetResourceProperty struct {
	// The Network Load Balancer resource that a DNS target resource points to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-targetresource.html#cfn-route53recoveryreadiness-resourceset-targetresource-nlbresource
	//
	NlbResource interface{} `field:"optional" json:"nlbResource" yaml:"nlbResource"`
	// The Route 53 resource that a DNS target resource record points to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-targetresource.html#cfn-route53recoveryreadiness-resourceset-targetresource-r53resource
	//
	R53Resource interface{} `field:"optional" json:"r53Resource" yaml:"r53Resource"`
}

