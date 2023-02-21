package awsquicksight


// A rule defined to grant access on one or more restricted columns.
//
// Each dataset can have multiple rules. To create a restricted column, you add it to one or more rules. Each rule must contain at least one column and at least one user or group. To be able to see a restricted column, a user or group needs to be added to a rule for that column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnLevelPermissionRuleProperty := &columnLevelPermissionRuleProperty{
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	principals: []*string{
//   		jsii.String("principals"),
//   	},
//   }
//
type CfnDataSet_ColumnLevelPermissionRuleProperty struct {
	// An array of column names.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// An array of Amazon Resource Names (ARNs) for Amazon QuickSight users or groups.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}

