// AWS CDK Programmatic CLI library
package awscdkclilibalpha


// Configuration for creating a CLI from an AWS CDK App directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cli_lib_alpha "github.com/aws/aws-cdk-go/awscdkclilibalpha"
//
//   cdkAppDirectoryProps := &CdkAppDirectoryProps{
//   	App: jsii.String("app"),
//   	Output: jsii.String("output"),
//   }
//
// Experimental.
type CdkAppDirectoryProps struct {
	// Command-line for executing your app or a cloud assembly directory e.g. "node bin/my-app.js" or "cdk.out".
	// Experimental.
	App *string `field:"optional" json:"app" yaml:"app"`
	// Emits the synthesized cloud assembly into a directory.
	// Experimental.
	Output *string `field:"optional" json:"output" yaml:"output"`
}

