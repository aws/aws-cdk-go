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
type CfnNetworkInsightsAccessScope_AccessScopePathRequestProperty struct {
	// The destination.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The source.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The through resources.
	ThroughResources interface{} `field:"optional" json:"throughResources" yaml:"throughResources"`
}

