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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hub.html
//
type CfnHubProps struct {
	// Whether to automatically enable new controls when they are added to standards that are enabled.
	//
	// By default, this is set to `true` , and new controls are enabled automatically. To not automatically enable new controls, set this to `false` .
	//
	// When you automatically enable new controls, you can interact with the controls in the console and programmatically immediately after release. However, automatically enabled controls have a temporary default status of `DISABLED` . It can take up to several days for Security Hub CSPM to process the control release and designate the control as `ENABLED` in your account. During the processing period, you can manually enable or disable a control, and Security Hub CSPM will maintain that designation regardless of whether you have `AutoEnableControls` set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hub.html#cfn-securityhub-hub-autoenablecontrols
	//
	AutoEnableControls interface{} `field:"optional" json:"autoEnableControls" yaml:"autoEnableControls"`
	// Specifies whether an account has consolidated control findings turned on or off.
	//
	// If the value for this field is set to `SECURITY_CONTROL` , Security Hub CSPM generates a single finding for a control check even when the check applies to multiple enabled standards.
	//
	// If the value for this field is set to `STANDARD_CONTROL` , Security Hub CSPM generates separate findings for a control check when the check applies to multiple enabled standards.
	//
	// The value for this field in a member account matches the value in the administrator account. For accounts that aren't part of an organization, the default value of this field is `SECURITY_CONTROL` if you enabled Security Hub CSPM on or after February 23, 2023.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hub.html#cfn-securityhub-hub-controlfindinggenerator
	//
	ControlFindingGenerator *string `field:"optional" json:"controlFindingGenerator" yaml:"controlFindingGenerator"`
	// Whether to enable the security standards that Security Hub CSPM has designated as automatically enabled.
	//
	// If you don't provide a value for `EnableDefaultStandards` , it is set to `true` , and the designated standards are automatically enabled in each AWS Region where you enable Security Hub CSPM . If you don't want to enable the designated standards, set `EnableDefaultStandards` to `false` .
	//
	// Currently, the automatically enabled standards are the Center for Internet Security (CIS) AWS Foundations Benchmark v1.2.0 and AWS Foundational Security Best Practices (FSBP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hub.html#cfn-securityhub-hub-enabledefaultstandards
	//
	EnableDefaultStandards interface{} `field:"optional" json:"enableDefaultStandards" yaml:"enableDefaultStandards"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hub.html#cfn-securityhub-hub-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

