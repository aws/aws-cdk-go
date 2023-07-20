package awsrefactorspaces


// The configuration for the URI path route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uriPathRouteInputProperty := &UriPathRouteInputProperty{
//   	ActivationState: jsii.String("activationState"),
//
//   	// the properties below are optional
//   	AppendSourcePath: jsii.Boolean(false),
//   	IncludeChildPaths: jsii.Boolean(false),
//   	Methods: []*string{
//   		jsii.String("methods"),
//   	},
//   	SourcePath: jsii.String("sourcePath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html
//
type CfnRoute_UriPathRouteInputProperty struct {
	// If set to `ACTIVE` , traffic is forwarded to this route’s service after the route is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html#cfn-refactorspaces-route-uripathrouteinput-activationstate
	//
	ActivationState *string `field:"required" json:"activationState" yaml:"activationState"`
	// If set to `true` , this option appends the source path to the service URL endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html#cfn-refactorspaces-route-uripathrouteinput-appendsourcepath
	//
	AppendSourcePath interface{} `field:"optional" json:"appendSourcePath" yaml:"appendSourcePath"`
	// Indicates whether to match all subpaths of the given source path.
	//
	// If this value is `false` , requests must match the source path exactly before they are forwarded to this route's service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html#cfn-refactorspaces-route-uripathrouteinput-includechildpaths
	//
	IncludeChildPaths interface{} `field:"optional" json:"includeChildPaths" yaml:"includeChildPaths"`
	// A list of HTTP methods to match.
	//
	// An empty list matches all values. If a method is present, only HTTP requests using that method are forwarded to this route’s service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html#cfn-refactorspaces-route-uripathrouteinput-methods
	//
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// This is the path that Refactor Spaces uses to match traffic.
	//
	// Paths must start with `/` and are relative to the base of the application. To use path parameters in the source path, add a variable in curly braces. For example, the resource path {user} represents a path parameter called 'user'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-uripathrouteinput.html#cfn-refactorspaces-route-uripathrouteinput-sourcepath
	//
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

