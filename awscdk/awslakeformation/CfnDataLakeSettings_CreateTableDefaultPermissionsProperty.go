package awslakeformation


// Specifies whether access control on a newly created table is managed by Lake Formation permissions or exclusively by IAM permissions.
//
// A null value indicates that the access is controlled by Lake Formation permissions. A value that assigns `ALL` to `IAM_ALLOWED_PRINCIPALS` indicates access control by IAM permissions. This is referred to as the setting "Use only IAM access control," and is for backward compatibility with the AWS Glue permission model implemented by IAM permissions.
//
// The only permitted values are an empty array or an array that contains a single JSON object that grants `ALL` to `IAM_ALLOWED_PRINCIPALS` .
//
// For more information, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createTableDefaultPermissionsProperty := &CreateTableDefaultPermissionsProperty{
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datalakesettings-createtabledefaultpermissions.html
//
type CfnDataLakeSettings_CreateTableDefaultPermissionsProperty struct {
}

