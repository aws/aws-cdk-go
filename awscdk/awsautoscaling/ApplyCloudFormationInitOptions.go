package awsautoscaling


// Options for applying CloudFormation init to an instance or instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applyCloudFormationInitOptions := &ApplyCloudFormationInitOptions{
//   	ConfigSets: []*string{
//   		jsii.String("configSets"),
//   	},
//   	EmbedFingerprint: jsii.Boolean(false),
//   	IgnoreFailures: jsii.Boolean(false),
//   	IncludeRole: jsii.Boolean(false),
//   	IncludeUrl: jsii.Boolean(false),
//   	PrintLog: jsii.Boolean(false),
//   }
//
type ApplyCloudFormationInitOptions struct {
	// ConfigSet to activate.
	// Default: ['default'].
	//
	ConfigSets *[]*string `field:"optional" json:"configSets" yaml:"configSets"`
	// Force instance replacement by embedding a config fingerprint.
	//
	// If `true` (the default), a hash of the config will be embedded into the
	// UserData, so that if the config changes, the UserData changes and
	// instances will be replaced (given an UpdatePolicy has been configured on
	// the AutoScalingGroup).
	//
	// If `false`, no such hash will be embedded, and if the CloudFormation Init
	// config changes nothing will happen to the running instances. If a
	// config update introduces errors, you will not notice until after the
	// CloudFormation deployment successfully finishes and the next instance
	// fails to launch.
	// Default: true.
	//
	EmbedFingerprint *bool `field:"optional" json:"embedFingerprint" yaml:"embedFingerprint"`
	// Don't fail the instance creation when cfn-init fails.
	//
	// You can use this to prevent CloudFormation from rolling back when
	// instances fail to start up, to help in debugging.
	// Default: false.
	//
	IgnoreFailures *bool `field:"optional" json:"ignoreFailures" yaml:"ignoreFailures"`
	// Include --role argument when running cfn-init and cfn-signal commands.
	//
	// This will be the IAM instance profile attached to the EC2 instance.
	// Default: false.
	//
	IncludeRole *bool `field:"optional" json:"includeRole" yaml:"includeRole"`
	// Include --url argument when running cfn-init and cfn-signal commands.
	//
	// This will be the cloudformation endpoint in the deployed region
	// e.g. https://cloudformation.us-east-1.amazonaws.com
	// Default: false.
	//
	IncludeUrl *bool `field:"optional" json:"includeUrl" yaml:"includeUrl"`
	// Print the results of running cfn-init to the Instance System Log.
	//
	// By default, the output of running cfn-init is written to a log file
	// on the instance. Set this to `true` to print it to the System Log
	// (visible from the EC2 Console), `false` to not print it.
	//
	// (Be aware that the system log is refreshed at certain points in
	// time of the instance life cycle, and successful execution may
	// not always show up).
	// Default: true.
	//
	PrintLog *bool `field:"optional" json:"printLog" yaml:"printLog"`
}

