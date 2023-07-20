package awslakeformation


// A wildcard object, consisting of an optional list of excluded column names or indexes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnWildcardProperty := &ColumnWildcardProperty{
//   	ExcludedColumnNames: []*string{
//   		jsii.String("excludedColumnNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-columnwildcard.html
//
type CfnPermissions_ColumnWildcardProperty struct {
	// Excludes column names.
	//
	// Any column with this name will be excluded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-columnwildcard.html#cfn-lakeformation-permissions-columnwildcard-excludedcolumnnames
	//
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
}

