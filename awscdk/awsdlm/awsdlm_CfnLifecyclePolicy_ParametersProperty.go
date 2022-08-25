package awsdlm


// Specifies optional parameters to add to a policy.
//
// The set of valid parameters depends on the combination of policy type and resource type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametersProperty := &parametersProperty{
//   	excludeBootVolume: jsii.Boolean(false),
//   	excludeDataVolumeTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	noReboot: jsii.Boolean(false),
//   }
//
type CfnLifecyclePolicy_ParametersProperty struct {
	// [EBS Snapshot Management – Instance policies only] Indicates whether to exclude the root volume from snapshots created using [CreateSnapshots](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSnapshots.html) . The default is false.
	ExcludeBootVolume interface{} `field:"optional" json:"excludeBootVolume" yaml:"excludeBootVolume"`
	// `CfnLifecyclePolicy.ParametersProperty.ExcludeDataVolumeTags`.
	ExcludeDataVolumeTags interface{} `field:"optional" json:"excludeDataVolumeTags" yaml:"excludeDataVolumeTags"`
	// Applies to AMI lifecycle policies only.
	//
	// Indicates whether targeted instances are rebooted when the lifecycle policy runs. `true` indicates that targeted instances are not rebooted when the policy runs. `false` indicates that target instances are rebooted when the policy runs. The default is `true` (instances are not rebooted).
	NoReboot interface{} `field:"optional" json:"noReboot" yaml:"noReboot"`
}

