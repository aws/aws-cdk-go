package awsbackup


// Contains information about which resources to include or exclude from a backup plan using their tags.
//
// Conditions are case sensitive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionsProperty := &conditionsProperty{
//   	stringEquals: []interface{}{
//   		&conditionParameterProperty{
//   			conditionKey: jsii.String("conditionKey"),
//   			conditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	stringLike: []interface{}{
//   		&conditionParameterProperty{
//   			conditionKey: jsii.String("conditionKey"),
//   			conditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	stringNotEquals: []interface{}{
//   		&conditionParameterProperty{
//   			conditionKey: jsii.String("conditionKey"),
//   			conditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	stringNotLike: []interface{}{
//   		&conditionParameterProperty{
//   			conditionKey: jsii.String("conditionKey"),
//   			conditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   }
//
type CfnBackupSelection_ConditionsProperty struct {
	// Filters the values of your tagged resources for only those resources that you tagged with the same value.
	//
	// Also called "exact matching."
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// Filters the values of your tagged resources for matching tag values with the use of a wildcard character (*) anywhere in the string.
	//
	// For example, "prod*" or "*rod*" matches the tag value "production".
	StringLike interface{} `field:"optional" json:"stringLike" yaml:"stringLike"`
	// Filters the values of your tagged resources for only those resources that you tagged that do not have the same value.
	//
	// Also called "negated matching."
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
	// Filters the values of your tagged resources for non-matching tag values with the use of a wildcard character (*) anywhere in the string.
	StringNotLike interface{} `field:"optional" json:"stringNotLike" yaml:"stringNotLike"`
}

