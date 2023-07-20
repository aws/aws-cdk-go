package awsec2


// Describes a path.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessScopePathRequestProperty := &AccessScopePathRequestProperty{
//   	Destination: &PathStatementRequestProperty{
//   		PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   			DestinationAddresses: []*string{
//   				jsii.String("destinationAddresses"),
//   			},
//   			DestinationPorts: []*string{
//   				jsii.String("destinationPorts"),
//   			},
//   			DestinationPrefixLists: []*string{
//   				jsii.String("destinationPrefixLists"),
//   			},
//   			Protocols: []*string{
//   				jsii.String("protocols"),
//   			},
//   			SourceAddresses: []*string{
//   				jsii.String("sourceAddresses"),
//   			},
//   			SourcePorts: []*string{
//   				jsii.String("sourcePorts"),
//   			},
//   			SourcePrefixLists: []*string{
//   				jsii.String("sourcePrefixLists"),
//   			},
//   		},
//   		ResourceStatement: &ResourceStatementRequestProperty{
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			ResourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   		},
//   	},
//   	Source: &PathStatementRequestProperty{
//   		PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   			DestinationAddresses: []*string{
//   				jsii.String("destinationAddresses"),
//   			},
//   			DestinationPorts: []*string{
//   				jsii.String("destinationPorts"),
//   			},
//   			DestinationPrefixLists: []*string{
//   				jsii.String("destinationPrefixLists"),
//   			},
//   			Protocols: []*string{
//   				jsii.String("protocols"),
//   			},
//   			SourceAddresses: []*string{
//   				jsii.String("sourceAddresses"),
//   			},
//   			SourcePorts: []*string{
//   				jsii.String("sourcePorts"),
//   			},
//   			SourcePrefixLists: []*string{
//   				jsii.String("sourcePrefixLists"),
//   			},
//   		},
//   		ResourceStatement: &ResourceStatementRequestProperty{
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			ResourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   		},
//   	},
//   	ThroughResources: []interface{}{
//   		&ThroughResourcesStatementRequestProperty{
//   			ResourceStatement: &ResourceStatementRequestProperty{
//   				Resources: []*string{
//   					jsii.String("resources"),
//   				},
//   				ResourceTypes: []*string{
//   					jsii.String("resourceTypes"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-accessscopepathrequest.html
//
type CfnNetworkInsightsAccessScope_AccessScopePathRequestProperty struct {
	// The destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-accessscopepathrequest.html#cfn-ec2-networkinsightsaccessscope-accessscopepathrequest-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-accessscopepathrequest.html#cfn-ec2-networkinsightsaccessscope-accessscopepathrequest-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The through resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-accessscopepathrequest.html#cfn-ec2-networkinsightsaccessscope-accessscopepathrequest-throughresources
	//
	ThroughResources interface{} `field:"optional" json:"throughResources" yaml:"throughResources"`
}

