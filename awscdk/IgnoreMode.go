package awscdk


// Determines the ignore behavior to use.
//
// Example:
//   lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("my-python-handler")), &AssetOptions{
//   		Exclude: []*string{
//   			jsii.String("*.ignore"),
//   		},
//   		IgnoreMode: awscdk.IgnoreMode_DOCKER,
//   	}),
//   	Runtime: lambda.Runtime_PYTHON_3_9(),
//   	Handler: jsii.String("index.handler"),
//   })
//
type IgnoreMode string

const (
	// Ignores file paths based on simple glob patterns.
	//
	// This is the default for file assets.
	//
	// It is also the default for Docker image assets, unless the '@aws-cdk/aws-ecr-assets:dockerIgnoreSupport'
	// context flag is set.
	IgnoreMode_GLOB IgnoreMode = "GLOB"
	// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
	IgnoreMode_GIT IgnoreMode = "GIT"
	// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
	//
	// This is the default for Docker image assets if the '@aws-cdk/aws-ecr-assets:dockerIgnoreSupport'
	// context flag is set.
	IgnoreMode_DOCKER IgnoreMode = "DOCKER"
)

