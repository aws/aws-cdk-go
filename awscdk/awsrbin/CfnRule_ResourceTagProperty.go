package awsrbin


// The resource tag of the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceTagProperty := &ResourceTagProperty{
//   	ResourceTagKey: jsii.String("resourceTagKey"),
//   	ResourceTagValue: jsii.String("resourceTagValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-resourcetag.html
//
type CfnRule_ResourceTagProperty struct {
	// The tag key of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-resourcetag.html#cfn-rbin-rule-resourcetag-resourcetagkey
	//
	ResourceTagKey *string `field:"required" json:"resourceTagKey" yaml:"resourceTagKey"`
	// The tag value of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-resourcetag.html#cfn-rbin-rule-resourcetag-resourcetagvalue
	//
	ResourceTagValue *string `field:"required" json:"resourceTagValue" yaml:"resourceTagValue"`
}

