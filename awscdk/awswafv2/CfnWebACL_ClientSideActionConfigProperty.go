package awswafv2


// Client side action config for AntiDDOS AMR.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientSideActionConfigProperty := &ClientSideActionConfigProperty{
//   	Challenge: &ClientSideActionProperty{
//   		UsageOfAction: jsii.String("usageOfAction"),
//
//   		// the properties below are optional
//   		ExemptUriRegularExpressions: []interface{}{
//   			&RegexProperty{
//   				RegexString: jsii.String("regexString"),
//   			},
//   		},
//   		Sensitivity: jsii.String("sensitivity"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideactionconfig.html
//
type CfnWebACL_ClientSideActionConfigProperty struct {
	// Client side action config for AntiDDOS AMR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideactionconfig.html#cfn-wafv2-webacl-clientsideactionconfig-challenge
	//
	Challenge interface{} `field:"required" json:"challenge" yaml:"challenge"`
}

