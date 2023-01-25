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
//   cfnNetworkInsightsAccessScopeProps := &cfnNetworkInsightsAccessScopeProps{
//   	excludePaths: []interface{}{
//   		&accessScopePathRequestProperty{
//   			destination: &pathStatementRequestProperty{
//   				packetHeaderStatement: &packetHeaderStatementRequestProperty{
//   					destinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					destinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					destinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					sourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					sourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					sourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				resourceStatement: &resourceStatementRequestProperty{
//   					resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					resourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			source: &pathStatementRequestProperty{
//   				packetHeaderStatement: &packetHeaderStatementRequestProperty{
//   					destinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					destinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					destinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					sourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					sourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					sourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				resourceStatement: &resourceStatementRequestProperty{
//   					resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					resourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			throughResources: []interface{}{
//   				&throughResourcesStatementRequestProperty{
//   					resourceStatement: &resourceStatementRequestProperty{
//   						resources: []*string{
//   							jsii.String("resources"),
//   						},
//   						resourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	matchPaths: []interface{}{
//   		&accessScopePathRequestProperty{
//   			destination: &pathStatementRequestProperty{
//   				packetHeaderStatement: &packetHeaderStatementRequestProperty{
//   					destinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					destinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					destinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					sourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					sourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					sourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				resourceStatement: &resourceStatementRequestProperty{
//   					resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					resourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			source: &pathStatementRequestProperty{
//   				packetHeaderStatement: &packetHeaderStatementRequestProperty{
//   					destinationAddresses: []*string{
//   						jsii.String("destinationAddresses"),
//   					},
//   					destinationPorts: []*string{
//   						jsii.String("destinationPorts"),
//   					},
//   					destinationPrefixLists: []*string{
//   						jsii.String("destinationPrefixLists"),
//   					},
//   					protocols: []*string{
//   						jsii.String("protocols"),
//   					},
//   					sourceAddresses: []*string{
//   						jsii.String("sourceAddresses"),
//   					},
//   					sourcePorts: []*string{
//   						jsii.String("sourcePorts"),
//   					},
//   					sourcePrefixLists: []*string{
//   						jsii.String("sourcePrefixLists"),
//   					},
//   				},
//   				resourceStatement: &resourceStatementRequestProperty{
//   					resources: []*string{
//   						jsii.String("resources"),
//   					},
//   					resourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			throughResources: []interface{}{
//   				&throughResourcesStatementRequestProperty{
//   					resourceStatement: &resourceStatementRequestProperty{
//   						resources: []*string{
//   							jsii.String("resources"),
//   						},
//   						resourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

