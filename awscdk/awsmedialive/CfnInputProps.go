package awsmedialive


// Properties for defining a `CfnInput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnInputProps := &CfnInputProps{
//   	Destinations: []interface{}{
//   		&InputDestinationRequestProperty{
//   			StreamName: jsii.String("streamName"),
//   		},
//   	},
//   	InputDevices: []interface{}{
//   		&InputDeviceSettingsProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	InputSecurityGroups: []*string{
//   		jsii.String("inputSecurityGroups"),
//   	},
//   	MediaConnectFlows: []interface{}{
//   		&MediaConnectFlowRequestProperty{
//   			FlowArn: jsii.String("flowArn"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Sources: []interface{}{
//   		&InputSourceRequestProperty{
//   			PasswordParam: jsii.String("passwordParam"),
//   			Url: jsii.String("url"),
//   			Username: jsii.String("username"),
//   		},
//   	},
//   	Tags: tags,
//   	Type: jsii.String("type"),
//   	Vpc: &InputVpcRequestProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnInputProps struct {
	// Settings that apply only if the input is a push type of input.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Settings that apply only if the input is an Elemental Link input.
	InputDevices interface{} `field:"optional" json:"inputDevices" yaml:"inputDevices"`
	// The list of input security groups (referenced by IDs) to attach to the input if the input is a push type.
	InputSecurityGroups *[]*string `field:"optional" json:"inputSecurityGroups" yaml:"inputSecurityGroups"`
	// Settings that apply only if the input is a MediaConnect input.
	MediaConnectFlows interface{} `field:"optional" json:"mediaConnectFlows" yaml:"mediaConnectFlows"`
	// A name for the input.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM role for MediaLive to assume when creating a MediaConnect input or Amazon VPC input.
	//
	// This doesn't apply to other types of inputs. The role is identified by its ARN.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Settings that apply only if the input is a pull type of input.
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// A collection of tags for this input.
	//
	// Each tag is a key-value pair.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type for this input.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Settings that apply only if the input is an push input where the source is on Amazon VPC.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

