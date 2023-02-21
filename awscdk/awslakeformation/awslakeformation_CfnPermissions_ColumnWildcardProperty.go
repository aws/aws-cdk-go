package awslakeformation


// A wildcard object, consisting of an optional list of excluded column names or indexes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnWildcardProperty := &columnWildcardProperty{
//   	excludedColumnNames: []*string{
//   		jsii.String("excludedColumnNames"),
//   	},
//   }
//
type CfnPermissions_ColumnWildcardProperty struct {
	// Excludes column names.
	//
	// Any column with this name will be excluded.
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
}

