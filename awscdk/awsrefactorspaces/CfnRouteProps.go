package awsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteProps := &CfnRouteProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	RouteType: jsii.String("routeType"),
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
//
//   	// the properties below are optional
//   	DefaultRoute: &DefaultRouteInputProperty{
//   		ActivationState: jsii.String("activationState"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UriPathRoute: &UriPathRouteInputProperty{
//   		ActivationState: jsii.String("activationState"),
//
//   		// the properties below are optional
//   		AppendSourcePath: jsii.Boolean(false),
//   		IncludeChildPaths: jsii.Boolean(false),
//   		Methods: []*string{
//   			jsii.String("methods"),
//   		},
//   		SourcePath: jsii.String("sourcePath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html
//
type CfnRouteProps struct {
	// The unique identifier of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-applicationidentifier
	//
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The unique identifier of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The route type of the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-routetype
	//
	RouteType *string `field:"required" json:"routeType" yaml:"routeType"`
	// The unique identifier of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-serviceidentifier
	//
	ServiceIdentifier *string `field:"required" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// Configuration for the default route type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-defaultroute
	//
	DefaultRoute interface{} `field:"optional" json:"defaultRoute" yaml:"defaultRoute"`
	// The tags assigned to the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration for the URI path route type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-uripathroute
	//
	UriPathRoute interface{} `field:"optional" json:"uriPathRoute" yaml:"uriPathRoute"`
}

