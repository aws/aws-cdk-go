package awsdeadline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostConfigurationProperty := &HostConfigurationProperty{
//   	ScriptBody: jsii.String("scriptBody"),
//
//   	// the properties below are optional
//   	ScriptTimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-hostconfiguration.html
//
type CfnFleet_HostConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-hostconfiguration.html#cfn-deadline-fleet-hostconfiguration-scriptbody
	//
	ScriptBody *string `field:"required" json:"scriptBody" yaml:"scriptBody"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-hostconfiguration.html#cfn-deadline-fleet-hostconfiguration-scripttimeoutseconds
	//
	// Default: - 300.
	//
	ScriptTimeoutSeconds *float64 `field:"optional" json:"scriptTimeoutSeconds" yaml:"scriptTimeoutSeconds"`
}

