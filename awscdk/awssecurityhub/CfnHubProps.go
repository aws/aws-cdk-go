package awssecurityhub


// Properties for defining a `CfnHub`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnHubProps := &CfnHubProps{
//   	AutoEnableControls: jsii.Boolean(false),
//   	ControlFindingGenerator: jsii.String("controlFindingGenerator"),
//   	EnableDefaultStandards: jsii.Boolean(false),
//   	Tags: tags,
//   }
//
type CfnHubProps struct {
	// Whether to automatically enable new controls when they are added to standards that are enabled.
	//
	// By default, this is set to `true` , and new controls are enabled automatically. To not automatically enable new controls, set this to `false` .
	AutoEnableControls interface{} `field:"optional" json:"autoEnableControls" yaml:"autoEnableControls"`
	// Specifies whether an account has consolidated control findings turned on or off.
	//
	// If the value for this field is set to `SECURITY_CONTROL` , Security Hub generates a single finding for a control check even when the check applies to multiple enabled standards.
	//
	// If the value for this field is set to `STANDARD_CONTROL` , Security Hub generates separate findings for a control check when the check applies to multiple enabled standards.
	//
	// The value for this field in a member account matches the value in the administrator account. For accounts that aren't part of an organization, the default value of this field is `SECURITY_CONTROL` if you enabled Security Hub on or after February 23, 2023.
	ControlFindingGenerator *string `field:"optional" json:"controlFindingGenerator" yaml:"controlFindingGenerator"`
	// Whether to enable the security standards that Security Hub has designated as automatically enabled.
	//
	// If you don't provide a value for `EnableDefaultStandards` , it is set to `true` , and the designated standards are automatically enabled in each AWS Region where you enable Security Hub . If you don't want to enable the designated standards, set `EnableDefaultStandards` to `false` .
	//
	// Currently, the automatically enabled standards are the Center for Internet Security (CIS) AWS Foundations Benchmark v1.2.0 and AWS Foundational Security Best Practices (FSBP).
	EnableDefaultStandards interface{} `field:"optional" json:"enableDefaultStandards" yaml:"enableDefaultStandards"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

