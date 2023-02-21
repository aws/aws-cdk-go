package pipelines


// Additional options for adding a stack deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addStackOptions := &addStackOptions{
//   	executeRunOrder: jsii.Number(123),
//   	runOrder: jsii.Number(123),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type AddStackOptions struct {
	// Base runorder.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder *float64 `field:"optional" json:"executeRunOrder" yaml:"executeRunOrder"`
	// Base runorder.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
}

