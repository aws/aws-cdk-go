package mixinsawss3


// Properties for CfnMultiRegionAccessPointPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnMultiRegionAccessPointPolicyMixinProps := &CfnMultiRegionAccessPointPolicyMixinProps{
//   	MrapName: jsii.String("mrapName"),
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-multiregionaccesspointpolicy.html
//
type CfnMultiRegionAccessPointPolicyMixinProps struct {
	// The name of the Multi-Region Access Point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-multiregionaccesspointpolicy.html#cfn-s3-multiregionaccesspointpolicy-mrapname
	//
	MrapName *string `field:"optional" json:"mrapName" yaml:"mrapName"`
	// The access policy associated with the Multi-Region Access Point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-multiregionaccesspointpolicy.html#cfn-s3-multiregionaccesspointpolicy-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

