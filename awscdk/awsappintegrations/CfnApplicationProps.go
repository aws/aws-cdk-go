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
//   	Namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	ApplicationConfig: &ApplicationConfigProperty{
//   		ContactHandling: &ContactHandlingProperty{
//   			Scope: jsii.String("scope"),
//   		},
//   	},
//   	IframeConfig: &IframeConfigProperty{
//   		Allow: []*string{
//   			jsii.String("allow"),
//   		},
//   		Sandbox: []*string{
//   			jsii.String("sandbox"),
//   		},
//   	},
//   	InitializationTimeout: jsii.Number(123),
//   	IsService: jsii.Boolean(false),
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
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
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-applicationconfig
	//
	ApplicationConfig interface{} `field:"optional" json:"applicationConfig" yaml:"applicationConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-iframeconfig
	//
	IframeConfig interface{} `field:"optional" json:"iframeConfig" yaml:"iframeConfig"`
	// The initialization timeout in milliseconds.
	//
	// Required when IsService is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-initializationtimeout
	//
	InitializationTimeout *float64 `field:"optional" json:"initializationTimeout" yaml:"initializationTimeout"`
	// Indicates whether the application is a service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-application.html#cfn-appintegrations-application-isservice
	//
	// Default: - false.
	//
	IsService interface{} `field:"optional" json:"isService" yaml:"isService"`
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

