package previewawssecurityhubmixins


// Properties for CfnStandardPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStandardMixinProps := &CfnStandardMixinProps{
//   	DisabledStandardsControls: []interface{}{
//   		&StandardsControlProperty{
//   			Reason: jsii.String("reason"),
//   			StandardsControlArn: jsii.String("standardsControlArn"),
//   		},
//   	},
//   	StandardsArn: jsii.String("standardsArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-standard.html
//
type CfnStandardMixinProps struct {
	// Specifies which controls are to be disabled in a standard.
	//
	// *Maximum* : `100`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-standard.html#cfn-securityhub-standard-disabledstandardscontrols
	//
	DisabledStandardsControls interface{} `field:"optional" json:"disabledStandardsControls" yaml:"disabledStandardsControls"`
	// The ARN of the standard that you want to enable.
	//
	// To view a list of available Security Hub standards and their ARNs, use the [`DescribeStandards`](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-standard.html#cfn-securityhub-standard-standardsarn
	//
	StandardsArn *string `field:"optional" json:"standardsArn" yaml:"standardsArn"`
}

