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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-iamrolearn
	//
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-protectedresourcetype
	//
	ProtectedResourceType *string `field:"required" json:"protectedResourceType" yaml:"protectedResourceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-restoretestingplanname
	//
	RestoreTestingPlanName *string `field:"required" json:"restoreTestingPlanName" yaml:"restoreTestingPlanName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-restoretestingselectionname
	//
	RestoreTestingSelectionName *string `field:"required" json:"restoreTestingSelectionName" yaml:"restoreTestingSelectionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-protectedresourcearns
	//
	ProtectedResourceArns *[]*string `field:"optional" json:"protectedResourceArns" yaml:"protectedResourceArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-protectedresourceconditions
	//
	ProtectedResourceConditions interface{} `field:"optional" json:"protectedResourceConditions" yaml:"protectedResourceConditions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-restoremetadataoverrides
	//
	RestoreMetadataOverrides interface{} `field:"optional" json:"restoreMetadataOverrides" yaml:"restoreMetadataOverrides"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-validationwindowhours
	//
	ValidationWindowHours *float64 `field:"optional" json:"validationWindowHours" yaml:"validationWindowHours"`
}

