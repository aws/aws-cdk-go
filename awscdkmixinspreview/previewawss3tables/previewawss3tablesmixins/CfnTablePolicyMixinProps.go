package previewawss3tablesmixins


// Properties for CfnTablePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var resourcePolicy interface{}
//
//   cfnTablePolicyMixinProps := &CfnTablePolicyMixinProps{
//   	ResourcePolicy: resourcePolicy,
//   	TableArn: jsii.String("tableArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablepolicy.html
//
type CfnTablePolicyMixinProps struct {
	// The `JSON` that defines the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablepolicy.html#cfn-s3tables-tablepolicy-resourcepolicy
	//
	ResourcePolicy interface{} `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// The Amazon Resource Name (ARN) of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-tablepolicy.html#cfn-s3tables-tablepolicy-tablearn
	//
	TableArn *string `field:"optional" json:"tableArn" yaml:"tableArn"`
}

