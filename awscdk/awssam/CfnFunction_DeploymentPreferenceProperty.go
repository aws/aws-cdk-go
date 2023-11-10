package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentPreferenceProperty := &DeploymentPreferenceProperty{
//   	Alarms: []*string{
//   		jsii.String("alarms"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	Hooks: &HooksProperty{
//   		PostTraffic: jsii.String("postTraffic"),
//   		PreTraffic: jsii.String("preTraffic"),
//   	},
//   	Role: jsii.String("role"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-deploymentpreference.html
//
type CfnFunction_DeploymentPreferenceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-deploymentpreference.html#cfn-serverless-function-deploymentpreference-alarms
	//
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-deploymentpreference.html#cfn-serverless-function-deploymentpreference-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-deploymentpreference.html#cfn-serverless-function-deploymentpreference-hooks
	//
	Hooks interface{} `field:"optional" json:"hooks" yaml:"hooks"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-deploymentpreference.html#cfn-serverless-function-deploymentpreference-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-deploymentpreference.html#cfn-serverless-function-deploymentpreference-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

