package awscdk


// Options for the combine strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var arrayMergeStrategy IArrayMergeStrategy
//
//   combineStrategyOptions := &CombineStrategyOptions{
//   	Arrays: arrayMergeStrategy,
//   }
//
type CombineStrategyOptions struct {
	// Strategy for merging arrays.
	// Default: ArrayMergeStrategy.replace()
	//
	Arrays IArrayMergeStrategy `field:"optional" json:"arrays" yaml:"arrays"`
}

