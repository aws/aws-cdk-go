package awsappintegrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	ApplicationSourceConfig: &ApplicationSourceConfigProperty{
//   		ExternalUrlConfig: &ExternalUrlConfigProperty{
//   			AccessUrl: jsii.String("accessUrl"),
//
//   			// the properties below are optional
//   			ApprovedOrigins: []*string{
//   				jsii.String("approvedOrigins"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Namespace: jsii.String("namespace"),
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html
//
type CfnApplicationProps struct {
	// The configuration for where the application should be loaded from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-applicationsourceconfig
	//
	ApplicationSourceConfig interface{} `field:"required" json:"applicationSourceConfig" yaml:"applicationSourceConfig"`
	// The description of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The configuration of events or requests that the application has access to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-permissions
	//
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

