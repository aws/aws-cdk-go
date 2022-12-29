package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterCriteriaProperty := &filterCriteriaProperty{
//   	filters: []interface{}{
//   		&filterProperty{
//   			pattern: jsii.String("pattern"),
//   		},
//   	},
//   }
//
type CfnPipe_FilterCriteriaProperty struct {
	// `CfnPipe.FilterCriteriaProperty.Filters`.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

