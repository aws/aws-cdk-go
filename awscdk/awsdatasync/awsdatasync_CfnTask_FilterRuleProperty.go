package awsdatasync


// Specifies which files, folders, and objects to include or exclude when transferring files from source to destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterRuleProperty := &filterRuleProperty{
//   	filterType: jsii.String("filterType"),
//   	value: jsii.String("value"),
//   }
//
type CfnTask_FilterRuleProperty struct {
	// The type of filter rule to apply.
	//
	// AWS DataSync only supports the SIMPLE_PATTERN rule type.
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// A single filter string that consists of the patterns to include or exclude.
	//
	// The patterns are delimited by "|" (that is, a pipe), for example: `/folder1|/folder2`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

