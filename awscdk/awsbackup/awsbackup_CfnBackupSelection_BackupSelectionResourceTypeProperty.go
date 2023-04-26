package awsbackup


// Specifies an object containing properties used to assign a set of resources to a backup plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var conditions interface{}
//
//   backupSelectionResourceTypeProperty := &BackupSelectionResourceTypeProperty{
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	SelectionName: jsii.String("selectionName"),
//
//   	// the properties below are optional
//   	Conditions: conditions,
//   	ListOfTags: []interface{}{
//   		&ConditionResourceTypeProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionType: jsii.String("conditionType"),
//   			ConditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	NotResources: []*string{
//   		jsii.String("notResources"),
//   	},
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   }
//
type CfnBackupSelection_BackupSelectionResourceTypeProperty struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when backing up the target resource;
	//
	// for example, `arn:aws:iam::123456789012:role/S3Access` .
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The display name of a resource selection document.
	SelectionName *string `field:"required" json:"selectionName" yaml:"selectionName"`
	// A list of conditions that you define to assign resources to your backup plans using tags.
	//
	// For example, `"StringEquals": { "ConditionKey": "aws:ResourceTag/CreatedByCryo", "ConditionValue": "true" },` . Condition operators are case sensitive.
	//
	// `Conditions` differs from `ListOfTags` as follows:
	//
	// - When you specify more than one condition, you only assign the resources that match ALL conditions (using AND logic).
	// - `Conditions` supports `StringEquals` , `StringLike` , `StringNotEquals` , and `StringNotLike` . `ListOfTags` only supports `StringEquals` .
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// A list of conditions that you define to assign resources to your backup plans using tags.
	//
	// For example, `"StringEquals": { "ConditionKey": "aws:ResourceTag/CreatedByCryo", "ConditionValue": "true" },` . Condition operators are case sensitive.
	//
	// `ListOfTags` differs from `Conditions` as follows:
	//
	// - When you specify more than one condition, you assign all resources that match AT LEAST ONE condition (using OR logic).
	// - `ListOfTags` only supports `StringEquals` . `Conditions` supports `StringEquals` , `StringLike` , `StringNotEquals` , and `StringNotLike` .
	ListOfTags interface{} `field:"optional" json:"listOfTags" yaml:"listOfTags"`
	// A list of Amazon Resource Names (ARNs) to exclude from a backup plan.
	//
	// The maximum number of ARNs is 500 without wildcards, or 30 ARNs with wildcards.
	//
	// If you need to exclude many resources from a backup plan, consider a different resource selection strategy, such as assigning only one or a few resource types or refining your resource selection using tags.
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// An array of strings that contain Amazon Resource Names (ARNs) of resources to assign to a backup plan.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

