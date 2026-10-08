package awsnetworksecuritymanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnScopePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnScopeMixinProps := &CfnScopeMixinProps{
//   	ScopeConfiguration: jsii.String("scopeConfiguration"),
//   	ScopeDescription: jsii.String("scopeDescription"),
//   	ScopeName: jsii.String("scopeName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-scope.html
//
type CfnScopeMixinProps struct {
	// The scope configuration as a JSON string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-scope.html#cfn-networksecuritymanager-scope-scopeconfiguration
	//
	ScopeConfiguration *string `field:"optional" json:"scopeConfiguration" yaml:"scopeConfiguration"`
	// A description of the scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-scope.html#cfn-networksecuritymanager-scope-scopedescription
	//
	ScopeDescription *string `field:"optional" json:"scopeDescription" yaml:"scopeDescription"`
	// The name of the scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-scope.html#cfn-networksecuritymanager-scope-scopename
	//
	ScopeName *string `field:"optional" json:"scopeName" yaml:"scopeName"`
	// The tags associated with the scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-scope.html#cfn-networksecuritymanager-scope-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

