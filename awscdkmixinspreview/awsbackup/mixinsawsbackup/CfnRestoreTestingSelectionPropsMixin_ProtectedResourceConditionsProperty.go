package mixinsawsbackup


// The conditions that you define for resources in your restore testing plan using tags.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   protectedResourceConditionsProperty := &ProtectedResourceConditionsProperty{
//   	StringEquals: []interface{}{
//   		&KeyValueProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	StringNotEquals: []interface{}{
//   		&KeyValueProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingselection-protectedresourceconditions.html
//
type CfnRestoreTestingSelectionPropsMixin_ProtectedResourceConditionsProperty struct {
	// Filters the values of your tagged resources for only those resources that you tagged with the same value.
	//
	// Also called "exact matching."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingselection-protectedresourceconditions.html#cfn-backup-restoretestingselection-protectedresourceconditions-stringequals
	//
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// Filters the values of your tagged resources for only those resources that you tagged that do not have the same value.
	//
	// Also called "negated matching."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingselection-protectedresourceconditions.html#cfn-backup-restoretestingselection-protectedresourceconditions-stringnotequals
	//
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
}

