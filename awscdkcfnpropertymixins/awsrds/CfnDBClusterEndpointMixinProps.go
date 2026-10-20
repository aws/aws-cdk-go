package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDBClusterEndpointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDBClusterEndpointMixinProps := &CfnDBClusterEndpointMixinProps{
//   	CustomEndpointType: jsii.String("customEndpointType"),
//   	DbClusterEndpointIdentifier: jsii.String("dbClusterEndpointIdentifier"),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	ExcludedMembers: []*string{
//   		jsii.String("excludedMembers"),
//   	},
//   	StaticMembers: []*string{
//   		jsii.String("staticMembers"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterendpoint.html
//
type CfnDBClusterEndpointMixinProps struct {
	// The type of the custom endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterendpoint.html#cfn-rds-dbclusterendpoint-customendpointtype
	//
	CustomEndpointType *string `field:"optional" json:"customEndpointType" yaml:"customEndpointType"`
	// The identifier to use for the new endpoint.
	//
	// This parameter is stored as a lowercase string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterendpoint.html#cfn-rds-dbclusterendpoint-dbclusterendpointidentifier
	//
	DbClusterEndpointIdentifier *string `field:"optional" json:"dbClusterEndpointIdentifier" yaml:"dbClusterEndpointIdentifier"`
	// The DB cluster identifier of the DB cluster associated with the endpoint.
	//
	// This parameter is stored as a lowercase string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterendpoint.html#cfn-rds-dbclusterendpoint-dbclusteridentifier
	//
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// List of DB instance identifiers that aren't part of the custom endpoint group.
	//
	// All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Once either member list is set, it can be changed or replaced by the other list, but both lists cannot be removed in place; removing them requires replacing the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterendpoint.html#cfn-rds-dbclusterendpoint-excludedmembers
	//
	ExcludedMembers *[]*string `field:"optional" json:"excludedMembers" yaml:"excludedMembers"`
	// List of DB instance identifiers that are part of the custom endpoint group.
	//
	// Once either member list is set, it can be changed or replaced by the other list, but both lists cannot be removed in place; removing them requires replacing the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterendpoint.html#cfn-rds-dbclusterendpoint-staticmembers
	//
	StaticMembers *[]*string `field:"optional" json:"staticMembers" yaml:"staticMembers"`
	// The tags to be assigned to the DB cluster endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterendpoint.html#cfn-rds-dbclusterendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

