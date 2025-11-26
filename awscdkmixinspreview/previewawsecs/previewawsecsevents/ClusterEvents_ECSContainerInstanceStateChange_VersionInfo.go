package previewawsecsevents


// Type definition for VersionInfo.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   versionInfo := &VersionInfo{
//   	AgentHash: []*string{
//   		jsii.String("agentHash"),
//   	},
//   	AgentVersion: []*string{
//   		jsii.String("agentVersion"),
//   	},
//   	DockerVersion: []*string{
//   		jsii.String("dockerVersion"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSContainerInstanceStateChange_VersionInfo struct {
	// agentHash property.
	//
	// Specify an array of string values to match this event if the actual value of agentHash is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AgentHash *[]*string `field:"optional" json:"agentHash" yaml:"agentHash"`
	// agentVersion property.
	//
	// Specify an array of string values to match this event if the actual value of agentVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AgentVersion *[]*string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// dockerVersion property.
	//
	// Specify an array of string values to match this event if the actual value of dockerVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DockerVersion *[]*string `field:"optional" json:"dockerVersion" yaml:"dockerVersion"`
}

