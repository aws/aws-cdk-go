package cloudassemblyschema


// A construct that violated a policy rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   violatingConstructJson := &ViolatingConstructJson{
//   	ConstructPath: jsii.String("constructPath"),
//
//   	// the properties below are optional
//   	CloudFormationResource: &CloudFormationResourceJson{
//   		LogicalId: jsii.String("logicalId"),
//   		TemplatePath: jsii.String("templatePath"),
//
//   		// the properties below are optional
//   		PropertyPaths: []*string{
//   			jsii.String("propertyPaths"),
//   		},
//   	},
//   	ConstructFqn: jsii.String("constructFqn"),
//   	LibraryVersion: jsii.String("libraryVersion"),
//   	StackTraces: []*string{
//   		jsii.String("stackTraces"),
//   	},
//   }
//
type ViolatingConstructJson struct {
	// The construct path as defined in the application.
	// Default: - no construct path.
	//
	ConstructPath *string `field:"required" json:"constructPath" yaml:"constructPath"`
	// If this construct violation regards a CloudFormation resource, a reference to the resource details.
	CloudFormationResource *CloudFormationResourceJson `field:"optional" json:"cloudFormationResource" yaml:"cloudFormationResource"`
	// The fully qualified name of the construct class (includes the library name).
	// Default: - no construct info.
	//
	ConstructFqn *string `field:"optional" json:"constructFqn" yaml:"constructFqn"`
	// The version of the library that contains this construct.
	//
	// The library name is the first component of the construct FQN.
	// Default: - no version info.
	//
	LibraryVersion *string `field:"optional" json:"libraryVersion" yaml:"libraryVersion"`
	// Stack traces associated with this violation.
	//
	// This can be all the stack traces where a violating property got its value,
	// or just the construct creation stack trace.
	//
	// Every element of the array is a stack trace, where each stack trace is a `\n`-delimited string.
	// Default: - No stack traces.
	//
	StackTraces *[]*string `field:"optional" json:"stackTraces" yaml:"stackTraces"`
}

