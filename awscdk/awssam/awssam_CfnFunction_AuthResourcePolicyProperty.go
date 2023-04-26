package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customStatements interface{}
//
//   authResourcePolicyProperty := &AuthResourcePolicyProperty{
//   	AwsAccountBlacklist: []*string{
//   		jsii.String("awsAccountBlacklist"),
//   	},
//   	AwsAccountWhitelist: []*string{
//   		jsii.String("awsAccountWhitelist"),
//   	},
//   	CustomStatements: []interface{}{
//   		customStatements,
//   	},
//   	IntrinsicVpcBlacklist: []*string{
//   		jsii.String("intrinsicVpcBlacklist"),
//   	},
//   	IntrinsicVpceBlacklist: []*string{
//   		jsii.String("intrinsicVpceBlacklist"),
//   	},
//   	IntrinsicVpceWhitelist: []*string{
//   		jsii.String("intrinsicVpceWhitelist"),
//   	},
//   	IntrinsicVpcWhitelist: []*string{
//   		jsii.String("intrinsicVpcWhitelist"),
//   	},
//   	IpRangeBlacklist: []*string{
//   		jsii.String("ipRangeBlacklist"),
//   	},
//   	IpRangeWhitelist: []*string{
//   		jsii.String("ipRangeWhitelist"),
//   	},
//   	SourceVpcBlacklist: []*string{
//   		jsii.String("sourceVpcBlacklist"),
//   	},
//   	SourceVpcWhitelist: []*string{
//   		jsii.String("sourceVpcWhitelist"),
//   	},
//   }
//
type CfnFunction_AuthResourcePolicyProperty struct {
	// `CfnFunction.AuthResourcePolicyProperty.AwsAccountBlacklist`.
	AwsAccountBlacklist *[]*string `field:"optional" json:"awsAccountBlacklist" yaml:"awsAccountBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.AwsAccountWhitelist`.
	AwsAccountWhitelist *[]*string `field:"optional" json:"awsAccountWhitelist" yaml:"awsAccountWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.CustomStatements`.
	CustomStatements interface{} `field:"optional" json:"customStatements" yaml:"customStatements"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpcBlacklist`.
	IntrinsicVpcBlacklist *[]*string `field:"optional" json:"intrinsicVpcBlacklist" yaml:"intrinsicVpcBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpceBlacklist`.
	IntrinsicVpceBlacklist *[]*string `field:"optional" json:"intrinsicVpceBlacklist" yaml:"intrinsicVpceBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpceWhitelist`.
	IntrinsicVpceWhitelist *[]*string `field:"optional" json:"intrinsicVpceWhitelist" yaml:"intrinsicVpceWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpcWhitelist`.
	IntrinsicVpcWhitelist *[]*string `field:"optional" json:"intrinsicVpcWhitelist" yaml:"intrinsicVpcWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.IpRangeBlacklist`.
	IpRangeBlacklist *[]*string `field:"optional" json:"ipRangeBlacklist" yaml:"ipRangeBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.IpRangeWhitelist`.
	IpRangeWhitelist *[]*string `field:"optional" json:"ipRangeWhitelist" yaml:"ipRangeWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.SourceVpcBlacklist`.
	SourceVpcBlacklist *[]*string `field:"optional" json:"sourceVpcBlacklist" yaml:"sourceVpcBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.SourceVpcWhitelist`.
	SourceVpcWhitelist *[]*string `field:"optional" json:"sourceVpcWhitelist" yaml:"sourceVpcWhitelist"`
}

