package mixinsawssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html
//
type CfnFunctionPropsMixin_AuthResourcePolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-awsaccountblacklist
	//
	AwsAccountBlacklist *[]*string `field:"optional" json:"awsAccountBlacklist" yaml:"awsAccountBlacklist"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-awsaccountwhitelist
	//
	AwsAccountWhitelist *[]*string `field:"optional" json:"awsAccountWhitelist" yaml:"awsAccountWhitelist"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-customstatements
	//
	CustomStatements interface{} `field:"optional" json:"customStatements" yaml:"customStatements"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-intrinsicvpcblacklist
	//
	IntrinsicVpcBlacklist *[]*string `field:"optional" json:"intrinsicVpcBlacklist" yaml:"intrinsicVpcBlacklist"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-intrinsicvpceblacklist
	//
	IntrinsicVpceBlacklist *[]*string `field:"optional" json:"intrinsicVpceBlacklist" yaml:"intrinsicVpceBlacklist"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-intrinsicvpcewhitelist
	//
	IntrinsicVpceWhitelist *[]*string `field:"optional" json:"intrinsicVpceWhitelist" yaml:"intrinsicVpceWhitelist"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-intrinsicvpcwhitelist
	//
	IntrinsicVpcWhitelist *[]*string `field:"optional" json:"intrinsicVpcWhitelist" yaml:"intrinsicVpcWhitelist"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-iprangeblacklist
	//
	IpRangeBlacklist *[]*string `field:"optional" json:"ipRangeBlacklist" yaml:"ipRangeBlacklist"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-iprangewhitelist
	//
	IpRangeWhitelist *[]*string `field:"optional" json:"ipRangeWhitelist" yaml:"ipRangeWhitelist"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-sourcevpcblacklist
	//
	SourceVpcBlacklist *[]*string `field:"optional" json:"sourceVpcBlacklist" yaml:"sourceVpcBlacklist"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-authresourcepolicy.html#cfn-serverless-function-authresourcepolicy-sourcevpcwhitelist
	//
	SourceVpcWhitelist *[]*string `field:"optional" json:"sourceVpcWhitelist" yaml:"sourceVpcWhitelist"`
}

