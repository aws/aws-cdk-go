package awsbackup


// Properties for defining a `CfnRestoreTestingSelection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRestoreTestingSelectionProps := &CfnRestoreTestingSelectionProps{
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	ProtectedResourceType: jsii.String("protectedResourceType"),
//   	RestoreTestingPlanName: jsii.String("restoreTestingPlanName"),
//   	RestoreTestingSelectionName: jsii.String("restoreTestingSelectionName"),
//
//   	// the properties below are optional
//   	ProtectedResourceArns: []*string{
//   		jsii.String("protectedResourceArns"),
//   	},
//   	ProtectedResourceConditions: &ProtectedResourceConditionsProperty{
//   		StringEquals: []interface{}{
//   			&KeyValueProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		StringNotEquals: []interface{}{
//   			&KeyValueProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	RestoreMetadataOverrides: map[string]*string{
//   		"restoreMetadataOverridesKey": jsii.String("restoreMetadataOverrides"),
//   	},
//   	ValidationWindowHours: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html
//
type CfnRestoreTestingSelectionProps struct {
	// The Amazon Resource Name (ARN) of the IAM role that AWS Backup uses to create the target resource;
	//
	// for example: `arn:aws:iam::123456789012:role/S3Access` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-iamrolearn
	//
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The type of AWS resource included in a resource testing selection;
	//
	// for example, an Amazon EBS volume or an Amazon RDS database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-protectedresourcetype
	//
	ProtectedResourceType *string `field:"required" json:"protectedResourceType" yaml:"protectedResourceType"`
	// Unique string that is the name of the restore testing plan.
	//
	// The name cannot be changed after creation. The name must consist of only alphanumeric characters and underscores. Maximum length is 50.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-restoretestingplanname
	//
	RestoreTestingPlanName *string `field:"required" json:"restoreTestingPlanName" yaml:"restoreTestingPlanName"`
	// The unique name of the restore testing selection that belongs to the related restore testing plan.
	//
	// The name consists of only alphanumeric characters and underscores. Maximum length is 50.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-restoretestingselectionname
	//
	RestoreTestingSelectionName *string `field:"required" json:"restoreTestingSelectionName" yaml:"restoreTestingSelectionName"`
	// You can include specific ARNs, such as `ProtectedResourceArns: ["arn:aws:...", "arn:aws:..."]` or you can include a wildcard: `ProtectedResourceArns: ["*"]` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-protectedresourcearns
	//
	ProtectedResourceArns *[]*string `field:"optional" json:"protectedResourceArns" yaml:"protectedResourceArns"`
	// In a resource testing selection, this parameter filters by specific conditions such as `StringEquals` or `StringNotEquals` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-protectedresourceconditions
	//
	ProtectedResourceConditions interface{} `field:"optional" json:"protectedResourceConditions" yaml:"protectedResourceConditions"`
	// You can override certain restore metadata keys by including the parameter `RestoreMetadataOverrides` in the body of `RestoreTestingSelection` .
	//
	// Key values are not case sensitive.
	//
	// See the complete list of [restore testing inferred metadata](https://docs.aws.amazon.com/aws-backup/latest/devguide/restore-testing-inferred-metadata.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-restoremetadataoverrides
	//
	RestoreMetadataOverrides interface{} `field:"optional" json:"restoreMetadataOverrides" yaml:"restoreMetadataOverrides"`
	// This is amount of hours (1 to 168) available to run a validation script on the data.
	//
	// The data will be deleted upon the completion of the validation script or the end of the specified retention period, whichever comes first.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-validationwindowhours
	//
	ValidationWindowHours *float64 `field:"optional" json:"validationWindowHours" yaml:"validationWindowHours"`
}

