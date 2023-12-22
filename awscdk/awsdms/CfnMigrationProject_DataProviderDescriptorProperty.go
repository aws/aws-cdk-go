package awsdms


// Information about a data provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataProviderDescriptorProperty := &DataProviderDescriptorProperty{
//   	DataProviderArn: jsii.String("dataProviderArn"),
//   	DataProviderIdentifier: jsii.String("dataProviderIdentifier"),
//   	DataProviderName: jsii.String("dataProviderName"),
//   	SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-migrationproject-dataproviderdescriptor.html
//
type CfnMigrationProject_DataProviderDescriptorProperty struct {
	// The Amazon Resource Name (ARN) of the data provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-migrationproject-dataproviderdescriptor.html#cfn-dms-migrationproject-dataproviderdescriptor-dataproviderarn
	//
	DataProviderArn *string `field:"optional" json:"dataProviderArn" yaml:"dataProviderArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-migrationproject-dataproviderdescriptor.html#cfn-dms-migrationproject-dataproviderdescriptor-dataprovideridentifier
	//
	DataProviderIdentifier *string `field:"optional" json:"dataProviderIdentifier" yaml:"dataProviderIdentifier"`
	// The user-friendly name of the data provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-migrationproject-dataproviderdescriptor.html#cfn-dms-migrationproject-dataproviderdescriptor-dataprovidername
	//
	DataProviderName *string `field:"optional" json:"dataProviderName" yaml:"dataProviderName"`
	// The ARN of the role used to access AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-migrationproject-dataproviderdescriptor.html#cfn-dms-migrationproject-dataproviderdescriptor-secretsmanageraccessrolearn
	//
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The identifier of the AWS Secrets Manager Secret used to store access credentials for the data provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-migrationproject-dataproviderdescriptor.html#cfn-dms-migrationproject-dataproviderdescriptor-secretsmanagersecretid
	//
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

