package awsquicksight


// Properties for defining a `CfnLimitsProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLimitsProfileProps := &CfnLimitsProfileProps{
//   	AccountId: jsii.String("accountId"),
//   	ProfileName: jsii.String("profileName"),
//   	ResourceLimits: map[string]interface{}{
//   		"resourceLimitsKey": &ProfileLimitValueProperty{
//   			"maxValue": jsii.Number(123),
//   			"unit": jsii.String("unit"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html
//
type CfnLimitsProfileProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html#cfn-quicksight-limitsprofile-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html#cfn-quicksight-limitsprofile-profilename
	//
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html#cfn-quicksight-limitsprofile-resourcelimits
	//
	ResourceLimits interface{} `field:"required" json:"resourceLimits" yaml:"resourceLimits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-limitsprofile.html#cfn-quicksight-limitsprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

