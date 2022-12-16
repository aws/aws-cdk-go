package awsrefactorspaces


// The configuration for the URI path route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uriPathRouteInputProperty := &uriPathRouteInputProperty{
//   	activationState: jsii.String("activationState"),
//
//   	// the properties below are optional
//   	includeChildPaths: jsii.Boolean(false),
//   	methods: []*string{
//   		jsii.String("methods"),
//   	},
//   	sourcePath: jsii.String("sourcePath"),
//   }
//
type CfnRoute_UriPathRouteInputProperty struct {
	// Indicates whether traffic is forwarded to this route’s service after the route is created.
	ActivationState *string `field:"required" json:"activationState" yaml:"activationState"`
	// Indicates whether to match all subpaths of the given source path.
	//
	// If this value is `false` , requests must match the source path exactly before they are forwarded to this route's service.
	IncludeChildPaths interface{} `field:"optional" json:"includeChildPaths" yaml:"includeChildPaths"`
	// A list of HTTP methods to match.
	//
	// An empty list matches all values. If a method is present, only HTTP requests using that method are forwarded to this route’s service.
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// The path to use to match traffic.
	//
	// Paths must start with `/` and are relative to the base of the application.
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

