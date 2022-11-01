package awsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteProps := &cfnRouteProps{
//   	applicationIdentifier: jsii.String("applicationIdentifier"),
//   	environmentIdentifier: jsii.String("environmentIdentifier"),
//   	serviceIdentifier: jsii.String("serviceIdentifier"),
//
//   	// the properties below are optional
//   	defaultRoute: &defaultRouteInputProperty{
//   		activationState: jsii.String("activationState"),
//   	},
//   	routeType: jsii.String("routeType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	uriPathRoute: &uriPathRouteInputProperty{
//   		activationState: jsii.String("activationState"),
//
//   		// the properties below are optional
//   		includeChildPaths: jsii.Boolean(false),
//   		methods: []*string{
//   			jsii.String("methods"),
//   		},
//   		sourcePath: jsii.String("sourcePath"),
//   	},
//   }
//
type CfnRouteProps struct {
	// The unique identifier of the application.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The unique identifier of the environment.
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The unique identifier of the service.
	ServiceIdentifier *string `field:"required" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// `AWS::RefactorSpaces::Route.DefaultRoute`.
	DefaultRoute interface{} `field:"optional" json:"defaultRoute" yaml:"defaultRoute"`
	// The route type of the route.
	RouteType *string `field:"optional" json:"routeType" yaml:"routeType"`
	// The tags assigned to the route.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration for the URI path route type.
	UriPathRoute interface{} `field:"optional" json:"uriPathRoute" yaml:"uriPathRoute"`
}

