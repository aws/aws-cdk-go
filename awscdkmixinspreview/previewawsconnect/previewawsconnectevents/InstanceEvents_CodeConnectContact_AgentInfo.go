package previewawsconnectevents


// Type definition for AgentInfo.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   agentInfo := &AgentInfo{
//   	AgentArn: []*string{
//   		jsii.String("agentArn"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_CodeConnectContact_AgentInfo struct {
	// agentArn property.
	//
	// Specify an array of string values to match this event if the actual value of agentArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AgentArn *[]*string `field:"optional" json:"agentArn" yaml:"agentArn"`
}

