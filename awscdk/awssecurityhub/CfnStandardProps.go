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
type CfnStandardProps struct {
	// The ARN of the standard that you want to enable.
	//
	// To view a list of available Security Hub standards and their ARNs, use the [`DescribeStandards`](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API operation.
	StandardsArn *string `field:"required" json:"standardsArn" yaml:"standardsArn"`
	// Specifies which controls are to be disabled in a standard.
	DisabledStandardsControls interface{} `field:"optional" json:"disabledStandardsControls" yaml:"disabledStandardsControls"`
}

