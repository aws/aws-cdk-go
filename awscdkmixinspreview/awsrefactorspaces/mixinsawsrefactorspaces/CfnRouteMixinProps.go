package mixinsawsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRoutePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteMixinProps := &CfnRouteMixinProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	DefaultRoute: &DefaultRouteInputProperty{
//   		ActivationState: jsii.String("activationState"),
//   	},
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	RouteType: jsii.String("routeType"),
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UriPathRoute: &UriPathRouteInputProperty{
//   		ActivationState: jsii.String("activationState"),
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
type CfnRouteMixinProps struct {
	// The unique identifier of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-applicationidentifier
	//
	ApplicationIdentifier *string `field:"optional" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// Configuration for the default route type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-defaultroute
	//
	DefaultRoute interface{} `field:"optional" json:"defaultRoute" yaml:"defaultRoute"`
	// The unique identifier of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The route type of the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-routetype
	//
	RouteType *string `field:"optional" json:"routeType" yaml:"routeType"`
	// The unique identifier of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-serviceidentifier
	//
	ServiceIdentifier *string `field:"optional" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// The tags assigned to the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration for the URI path route type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-route.html#cfn-refactorspaces-route-uripathroute
	//
	UriPathRoute interface{} `field:"optional" json:"uriPathRoute" yaml:"uriPathRoute"`
}

