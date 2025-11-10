package awsec2


// Example:
//   vpc := ec2.Vpc_FromLookup(this, jsii.String("ExistingVPC"), &VpcLookupOptions{
//   	VpcId: jsii.String("vpc-12345678"),
//   })
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("CustomSG"), &SecurityGroupProps{
//   	Vpc: vpc,
//   	Description: jsii.String("Allow HTTPS outbound access"),
//   	AllowAllOutbound: jsii.Boolean(false),
//   })
//
//   securityGroup.AddEgressRule(ec2.Peer_AnyIpv4(), ec2.Port_Tcp(jsii.Number(443)), jsii.String("Allow HTTPS traffic"))
//
//   s3deploy.NewBucketDeployment(this, jsii.String("DeployWithSecurityGroup"), &BucketDeploymentProps{
//   	DestinationBucket: bucket,
//   	Vpc: vpc,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Sources: []ISource{
//   		s3deploy.Source_Asset(jsii.String("./website")),
//   	},
//   })
//
type SecurityGroupProps struct {
	// The VPC in which to create the security group.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether to allow all outbound ipv6 traffic by default.
	//
	// If this is set to true, there will only be a single egress rule which allows all
	// outbound ipv6 traffic. If this is set to false, no outbound traffic will be allowed by
	// default and all egress ipv6 traffic must be explicitly authorized.
	//
	// To allow all ipv4 traffic use allowAllOutbound.
	// Default: false.
	//
	AllowAllIpv6Outbound *bool `field:"optional" json:"allowAllIpv6Outbound" yaml:"allowAllIpv6Outbound"`
	// Whether to allow all outbound traffic by default.
	//
	// If this is set to true, there will only be a single egress rule which allows all
	// outbound traffic. If this is set to false, no outbound traffic will be allowed by
	// default and all egress traffic must be explicitly authorized.
	//
	// To allow all ipv6 traffic use allowAllIpv6Outbound.
	// Default: true.
	//
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// A description of the security group.
	// Default: The default name will be the construct's CDK path.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to disable inline ingress and egress rule optimization.
	//
	// If this is set to true, ingress and egress rules will not be declared under the
	// SecurityGroup in cloudformation, but will be separate elements.
	//
	// Inlining rules is an optimization for producing smaller stack templates. Sometimes
	// this is not desirable, for example when security group access is managed via tags.
	//
	// The default value can be overridden globally by setting the context variable
	// '@aws-cdk/aws-ec2.securityGroupDisableInlineRules'.
	// Default: false.
	//
	DisableInlineRules *bool `field:"optional" json:"disableInlineRules" yaml:"disableInlineRules"`
	// The name of the security group.
	//
	// For valid values, see the GroupName
	// parameter of the CreateSecurityGroup action in the Amazon EC2 API
	// Reference.
	//
	// It is not recommended to use an explicit group name.
	// Default: If you don't specify a GroupName, AWS CloudFormation generates a
	// unique physical ID and uses that ID for the group name.
	//
	SecurityGroupName *string `field:"optional" json:"securityGroupName" yaml:"securityGroupName"`
}

