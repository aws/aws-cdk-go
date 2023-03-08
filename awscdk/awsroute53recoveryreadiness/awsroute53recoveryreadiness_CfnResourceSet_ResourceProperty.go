package awsroute53recoveryreadiness


// The resource element of a resource set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceProperty := &resourceProperty{
//   	componentId: jsii.String("componentId"),
//   	dnsTargetResource: &dNSTargetResourceProperty{
//   		domainName: jsii.String("domainName"),
//   		hostedZoneArn: jsii.String("hostedZoneArn"),
//   		recordSetId: jsii.String("recordSetId"),
//   		recordType: jsii.String("recordType"),
//   		targetResource: &targetResourceProperty{
//   			nlbResource: &nLBResourceProperty{
//   				arn: jsii.String("arn"),
//   			},
//   			r53Resource: &r53ResourceRecordProperty{
//   				domainName: jsii.String("domainName"),
//   				recordSetId: jsii.String("recordSetId"),
//   			},
//   		},
//   	},
//   	readinessScopes: []*string{
//   		jsii.String("readinessScopes"),
//   	},
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnResourceSet_ResourceProperty struct {
	// The component identifier of the resource, generated when DNS target resource is used.
	ComponentId *string `field:"optional" json:"componentId" yaml:"componentId"`
	// A component for DNS/routing control readiness checks.
	//
	// This is a required setting when `ResourceSet` `ResourceSetType` is set to `AWS::Route53RecoveryReadiness::DNSTargetResource` . Do not set it for any other `ResourceSetType` setting.
	DnsTargetResource interface{} `field:"optional" json:"dnsTargetResource" yaml:"dnsTargetResource"`
	// The recovery group Amazon Resource Name (ARN) or the cell ARN that the readiness checks for this resource set are scoped to.
	ReadinessScopes *[]*string `field:"optional" json:"readinessScopes" yaml:"readinessScopes"`
	// The Amazon Resource Name (ARN) of the AWS resource.
	//
	// This is a required setting for all `ResourceSet` `ResourceSetType` settings except `AWS::Route53RecoveryReadiness::DNSTargetResource` . Do not set this when `ResourceSetType` is set to `AWS::Route53RecoveryReadiness::DNSTargetResource` .
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

