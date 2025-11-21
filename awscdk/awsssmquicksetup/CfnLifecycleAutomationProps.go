package awsssmquicksetup


// Properties for defining a `CfnLifecycleAutomation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLifecycleAutomationProps := &CfnLifecycleAutomationProps{
//   	AutomationDocument: jsii.String("automationDocument"),
//   	AutomationParameters: map[string][]*string{
//   		"automationParametersKey": []*string{
//   			jsii.String("automationParameters"),
//   		},
//   	},
//   	ResourceKey: jsii.String("resourceKey"),
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-lifecycleautomation.html
//
type CfnLifecycleAutomationProps struct {
	// The name of the SSM Automation document to execute in response to CloudFormation lifecycle events (CREATE, UPDATE, DELETE).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-lifecycleautomation.html#cfn-ssmquicksetup-lifecycleautomation-automationdocument
	//
	AutomationDocument *string `field:"required" json:"automationDocument" yaml:"automationDocument"`
	// A map of key-value parameters passed to the Automation document during execution.
	//
	// Each parameter name maps to a list of values, even for single values. Parameters can include configuration-specific values for your automation workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-lifecycleautomation.html#cfn-ssmquicksetup-lifecycleautomation-automationparameters
	//
	AutomationParameters interface{} `field:"required" json:"automationParameters" yaml:"automationParameters"`
	// A unique identifier used for generating the SSM Association name.
	//
	// This ensures uniqueness when multiple lifecycle automation resources exist in the same stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-lifecycleautomation.html#cfn-ssmquicksetup-lifecycleautomation-resourcekey
	//
	ResourceKey *string `field:"required" json:"resourceKey" yaml:"resourceKey"`
	// Tags applied to the underlying SSM Association created by this resource.
	//
	// Tags help identify and organize automation executions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-lifecycleautomation.html#cfn-ssmquicksetup-lifecycleautomation-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

