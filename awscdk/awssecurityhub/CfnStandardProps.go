package awssecurityhub


// Properties for defining a `CfnStandard`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStandardProps := &CfnStandardProps{
//   	StandardsArn: jsii.String("standardsArn"),
//
//   	// the properties below are optional
//   	DisabledStandardsControls: []interface{}{
//   		&StandardsControlProperty{
//   			StandardsControlArn: jsii.String("standardsControlArn"),
//
//   			// the properties below are optional
//   			Reason: jsii.String("reason"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-standard.html
//
type CfnStandardProps struct {
	// The ARN of the standard that you want to enable.
	//
	// To view a list of available Security Hub CSPM standards and their ARNs, use the [`DescribeStandards`](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-standard.html#cfn-securityhub-standard-standardsarn
	//
	StandardsArn *string `field:"required" json:"standardsArn" yaml:"standardsArn"`
	// Specifies which controls are to be disabled in a standard.
	//
	// *Maximum* : `100`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-standard.html#cfn-securityhub-standard-disabledstandardscontrols
	//
	DisabledStandardsControls interface{} `field:"optional" json:"disabledStandardsControls" yaml:"disabledStandardsControls"`
}

