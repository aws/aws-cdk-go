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
//   			Network: jsii.String("network"),
//   			NetworkRoutes: []interface{}{
//   				&InputRequestDestinationRouteProperty{
//   					Cidr: jsii.String("cidr"),
//   					Gateway: jsii.String("gateway"),
//   				},
//   			},
//   			StaticIpAddress: jsii.String("staticIpAddress"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   	},
//   	InputDevices: []interface{}{
//   		&InputDeviceSettingsProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	InputNetworkLocation: jsii.String("inputNetworkLocation"),
//   	InputSecurityGroups: []*string{
//   		jsii.String("inputSecurityGroups"),
//   	},
//   	MediaConnectFlows: []interface{}{
//   		&MediaConnectFlowRequestProperty{
//   			FlowArn: jsii.String("flowArn"),
//   		},
//   	},
//   	MulticastSettings: &MulticastSettingsCreateRequestProperty{
//   		Sources: []interface{}{
//   			&MulticastSourceCreateRequestProperty{
//   				SourceIp: jsii.String("sourceIp"),
//   				Url: jsii.String("url"),
//   			},
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
//   	SrtSettings: &SrtSettingsRequestProperty{
//   		SrtCallerSources: []interface{}{
//   			&SrtCallerSourceRequestProperty{
//   				Decryption: &SrtCallerDecryptionRequestProperty{
//   					Algorithm: jsii.String("algorithm"),
//   					PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   				},
//   				MinimumLatency: jsii.Number(123),
//   				SrtListenerAddress: jsii.String("srtListenerAddress"),
//   				SrtListenerPort: jsii.String("srtListenerPort"),
//   				StreamId: jsii.String("streamId"),
//   			},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html
//
type CfnInputProps struct {
	// Settings that apply only if the input is a push type of input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Settings that apply only if the input is an Elemental Link input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-inputdevices
	//
	InputDevices interface{} `field:"optional" json:"inputDevices" yaml:"inputDevices"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-inputnetworklocation
	//
	InputNetworkLocation *string `field:"optional" json:"inputNetworkLocation" yaml:"inputNetworkLocation"`
	// The list of input security groups (referenced by IDs) to attach to the input if the input is a push type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-inputsecuritygroups
	//
	InputSecurityGroups *[]*string `field:"optional" json:"inputSecurityGroups" yaml:"inputSecurityGroups"`
	// Settings that apply only if the input is a MediaConnect input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-mediaconnectflows
	//
	MediaConnectFlows interface{} `field:"optional" json:"mediaConnectFlows" yaml:"mediaConnectFlows"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-multicastsettings
	//
	MulticastSettings interface{} `field:"optional" json:"multicastSettings" yaml:"multicastSettings"`
	// A name for the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM role for MediaLive to assume when creating a MediaConnect input or Amazon VPC input.
	//
	// This doesn't apply to other types of inputs. The role is identified by its ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Settings that apply only if the input is a pull type of input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-srtsettings
	//
	SrtSettings interface{} `field:"optional" json:"srtSettings" yaml:"srtSettings"`
	// A collection of tags for this input.
	//
	// Each tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type for this input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Settings that apply only if the input is an push input where the source is on Amazon VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html#cfn-medialive-input-vpc
	//
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

