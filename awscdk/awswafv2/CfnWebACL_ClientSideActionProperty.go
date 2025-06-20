package awswafv2


// Client side action config for AntiDDOS AMR.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientSideActionProperty := &ClientSideActionProperty{
//   	UsageOfAction: jsii.String("usageOfAction"),
//
//   	// the properties below are optional
//   	ExemptUriRegularExpressions: []interface{}{
//   		&RegexProperty{
//   			RegexString: jsii.String("regexString"),
//   		},
//   	},
//   	Sensitivity: jsii.String("sensitivity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideaction.html
//
type CfnWebACL_ClientSideActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideaction.html#cfn-wafv2-webacl-clientsideaction-usageofaction
	//
	UsageOfAction *string `field:"required" json:"usageOfAction" yaml:"usageOfAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideaction.html#cfn-wafv2-webacl-clientsideaction-exempturiregularexpressions
	//
	ExemptUriRegularExpressions interface{} `field:"optional" json:"exemptUriRegularExpressions" yaml:"exemptUriRegularExpressions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideaction.html#cfn-wafv2-webacl-clientsideaction-sensitivity
	//
	Sensitivity *string `field:"optional" json:"sensitivity" yaml:"sensitivity"`
}

