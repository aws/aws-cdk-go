package mixinsawsroute53recoveryreadiness


// The resource element of a resource set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceProperty := &ResourceProperty{
//   	ComponentId: jsii.String("componentId"),
//   	DnsTargetResource: &DNSTargetResourceProperty{
//   		DomainName: jsii.String("domainName"),
//   		HostedZoneArn: jsii.String("hostedZoneArn"),
//   		RecordSetId: jsii.String("recordSetId"),
//   		RecordType: jsii.String("recordType"),
//   		TargetResource: &TargetResourceProperty{
//   			NlbResource: &NLBResourceProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			R53Resource: &R53ResourceRecordProperty{
//   				DomainName: jsii.String("domainName"),
//   				RecordSetId: jsii.String("recordSetId"),
//   			},
//   		},
//   	},
//   	ReadinessScopes: []*string{
//   		jsii.String("readinessScopes"),
//   	},
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html
//
type CfnResourceSetPropsMixin_ResourceProperty struct {
	// The component identifier of the resource, generated when DNS target resource is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html#cfn-route53recoveryreadiness-resourceset-resource-componentid
	//
	ComponentId *string `field:"optional" json:"componentId" yaml:"componentId"`
	// A component for DNS/routing control readiness checks.
	//
	// This is a required setting when `ResourceSet` `ResourceSetType` is set to `AWS::Route53RecoveryReadiness::DNSTargetResource` . Do not set it for any other `ResourceSetType` setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html#cfn-route53recoveryreadiness-resourceset-resource-dnstargetresource
	//
	DnsTargetResource interface{} `field:"optional" json:"dnsTargetResource" yaml:"dnsTargetResource"`
	// The recovery group Amazon Resource Name (ARN) or the cell ARN that the readiness checks for this resource set are scoped to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html#cfn-route53recoveryreadiness-resourceset-resource-readinessscopes
	//
	ReadinessScopes *[]*string `field:"optional" json:"readinessScopes" yaml:"readinessScopes"`
	// The Amazon Resource Name (ARN) of the AWS resource.
	//
	// This is a required setting for all `ResourceSet` `ResourceSetType` settings except `AWS::Route53RecoveryReadiness::DNSTargetResource` . Do not set this when `ResourceSetType` is set to `AWS::Route53RecoveryReadiness::DNSTargetResource` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html#cfn-route53recoveryreadiness-resourceset-resource-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

