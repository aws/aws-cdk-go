package awscdk


// The deployment environment for a stack.
//
// Example:
//   stack1 := awscdk.Newstack(app, jsii.String("Stack1"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-east-1"),
//   	},
//   })
//   cert := acm.NewCertificate(stack1, jsii.String("Cert"), &CertificateProps{
//   	DomainName: jsii.String("*.example.com"),
//   	Validation: acm.CertificateValidation_FromDns(route53.PublicHostedZone_FromHostedZoneId(stack1, jsii.String("Zone"), jsii.String("Z0329774B51CGXTDQV3X"))),
//   })
//
//   stack2 := awscdk.Newstack(app, jsii.String("Stack2"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-east-2"),
//   	},
//   })
//   cloudfront.NewDistribution(stack2, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewHttpOrigin(jsii.String("example.com")),
//   	},
//   	DomainNames: []*string{
//   		jsii.String("dev.example.com"),
//   	},
//   	Certificate: cert,
//   })
//
type Environment struct {
	// The AWS account ID for this environment.
	//
	// This can be either a concrete value such as `585191031104` or `Aws.ACCOUNT_ID` which
	// indicates that account ID will only be determined during deployment (it
	// will resolve to the CloudFormation intrinsic `{"Ref":"AWS::AccountId"}`).
	// Note that certain features, such as cross-stack references and
	// environmental context providers require concrete region information and
	// will cause this stack to emit synthesis errors.
	// Default: Aws.ACCOUNT_ID which means that the stack will be account-agnostic.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The AWS region for this environment.
	//
	// This can be either a concrete value such as `eu-west-2` or `Aws.REGION`
	// which indicates that account ID will only be determined during deployment
	// (it will resolve to the CloudFormation intrinsic `{"Ref":"AWS::Region"}`).
	// Note that certain features, such as cross-stack references and
	// environmental context providers require concrete region information and
	// will cause this stack to emit synthesis errors.
	// Default: Aws.REGION which means that the stack will be region-agnostic.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

