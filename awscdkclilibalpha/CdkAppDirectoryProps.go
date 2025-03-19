package awscdkclilibalpha


// Configuration for creating a CLI from an AWS CDK App directory.
// Experimental.
type CdkAppDirectoryProps struct {
	// Command-line for executing your app or a cloud assembly directory e.g. "node bin/my-app.js" or "cdk.out".
	// Default: - read from cdk.json
	//
	// Experimental.
	App *string `field:"optional" json:"app" yaml:"app"`
	// Emits the synthesized cloud assembly into a directory.
	// Default: cdk.out
	//
	// Experimental.
	Output *string `field:"optional" json:"output" yaml:"output"`
}

