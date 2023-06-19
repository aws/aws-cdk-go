package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnNetworkInsightsAccessScope`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInsightsAccessScopeProps := &CfnNetworkInsightsAccessScopeProps{
//   	ExcludePaths: []interface{}{
//   		&AccessScopePathRequestProperty{
//   			Destination: &PathStatementRequestProperty{
//   				PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   					DestinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					DestinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					DestinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					Protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					SourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					SourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					SourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				ResourceStatement: &ResourceStatementRequestProperty{
//   					Resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			Source: &PathStatementRequestProperty{
//   				PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   					DestinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					DestinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					DestinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					Protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					SourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					SourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					SourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				ResourceStatement: &ResourceStatementRequestProperty{
//   					Resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			ThroughResources: []interface{}{
//   				&ThroughResourcesStatementRequestProperty{
//   					ResourceStatement: &ResourceStatementRequestProperty{
//   						Resources: []*string{
//   							jsii.String("resources"),
//   						},
//   						ResourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	MatchPaths: []interface{}{
//   		&AccessScopePathRequestProperty{
//   			Destination: &PathStatementRequestProperty{
//   				PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   					DestinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					DestinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					DestinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					Protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					SourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					SourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					SourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				ResourceStatement: &ResourceStatementRequestProperty{
//   					Resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			Source: &PathStatementRequestProperty{
//   				PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   					DestinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					DestinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					DestinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					Protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					SourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					SourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					SourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				ResourceStatement: &ResourceStatementRequestProperty{
//   					Resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			ThroughResources: []interface{}{
//   				&ThroughResourcesStatementRequestProperty{
//   					ResourceStatement: &ResourceStatementRequestProperty{
//   						Resources: []*string{
//   							jsii.String("resources"),
//   						},
//   						ResourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnNetworkInsightsAccessScopeProps struct {
	// The paths to exclude.
	ExcludePaths interface{} `field:"optional" json:"excludePaths" yaml:"excludePaths"`
	// The paths to match.
	MatchPaths interface{} `field:"optional" json:"matchPaths" yaml:"matchPaths"`
	// The tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

