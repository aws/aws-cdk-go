package awsdms


// Provides information that defines an Amazon Redshift endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Extra connection attributes when using Amazon Redshift as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redshift.html#CHAP_Target.Redshift.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftSettingsProperty := &RedshiftSettingsProperty{
//   	AcceptAnyDate: jsii.Boolean(false),
//   	AfterConnectScript: jsii.String("afterConnectScript"),
//   	BucketFolder: jsii.String("bucketFolder"),
//   	BucketName: jsii.String("bucketName"),
//   	CaseSensitiveNames: jsii.Boolean(false),
//   	CompUpdate: jsii.Boolean(false),
//   	ConnectionTimeout: jsii.Number(123),
//   	DateFormat: jsii.String("dateFormat"),
//   	EmptyAsNull: jsii.Boolean(false),
//   	EncryptionMode: jsii.String("encryptionMode"),
//   	ExplicitIds: jsii.Boolean(false),
//   	FileTransferUploadStreams: jsii.Number(123),
//   	LoadTimeout: jsii.Number(123),
//   	MapBooleanAsBoolean: jsii.Boolean(false),
//   	MaxFileSize: jsii.Number(123),
//   	RemoveQuotes: jsii.Boolean(false),
//   	ReplaceChars: jsii.String("replaceChars"),
//   	ReplaceInvalidChars: jsii.String("replaceInvalidChars"),
//   	SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	ServerSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   	ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	TimeFormat: jsii.String("timeFormat"),
//   	TrimBlanks: jsii.Boolean(false),
//   	TruncateColumns: jsii.Boolean(false),
//   	WriteBufferSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html
//
type CfnEndpoint_RedshiftSettingsProperty struct {
	// A value that indicates to allow any date format, including invalid formats such as 00/00/00 00:00:00, to be loaded without generating an error.
	//
	// You can choose `true` or `false` (the default).
	//
	// This parameter applies only to TIMESTAMP and DATE columns. Always use ACCEPTANYDATE with the DATEFORMAT parameter. If the date format for the data doesn't match the DATEFORMAT specification, Amazon Redshift inserts a NULL value into that field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-acceptanydate
	//
	AcceptAnyDate interface{} `field:"optional" json:"acceptAnyDate" yaml:"acceptAnyDate"`
	// Code to run after connecting.
	//
	// This parameter should contain the code itself, not the name of a file containing the code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-afterconnectscript
	//
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// An S3 folder where the comma-separated-value (.csv) files are stored before being uploaded to the target Redshift cluster.
	//
	// For full load mode, AWS DMS converts source records into .csv files and loads them to the *BucketFolder/TableID* path. AWS DMS uses the Redshift `COPY` command to upload the .csv files to the target table. The files are deleted once the `COPY` operation has finished. For more information, see [COPY](https://docs.aws.amazon.com/redshift/latest/dg/r_COPY.html) in the *Amazon Redshift Database Developer Guide* .
	//
	// For change-data-capture (CDC) mode, AWS DMS creates a *NetChanges* table, and loads the .csv files to this *BucketFolder/NetChangesTableID* path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-bucketfolder
	//
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// The name of the intermediate S3 bucket used to store .csv files before uploading data to Redshift.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// If Amazon Redshift is configured to support case sensitive schema names, set `CaseSensitiveNames` to `true` .
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-casesensitivenames
	//
	CaseSensitiveNames interface{} `field:"optional" json:"caseSensitiveNames" yaml:"caseSensitiveNames"`
	// If you set `CompUpdate` to `true` Amazon Redshift applies automatic compression if the table is empty.
	//
	// This applies even if the table columns already have encodings other than `RAW` . If you set `CompUpdate` to `false` , automatic compression is disabled and existing column encodings aren't changed. The default is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-compupdate
	//
	CompUpdate interface{} `field:"optional" json:"compUpdate" yaml:"compUpdate"`
	// A value that sets the amount of time to wait (in milliseconds) before timing out, beginning from when you initially establish a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-connectiontimeout
	//
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// The date format that you are using.
	//
	// Valid values are `auto` (case-sensitive), your date format string enclosed in quotes, or NULL. If this parameter is left unset (NULL), it defaults to a format of 'YYYY-MM-DD'. Using `auto` recognizes most strings, even some that aren't supported when you use a date format string.
	//
	// If your date and time values use formats different from each other, set this to `auto` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-dateformat
	//
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// A value that specifies whether AWS DMS should migrate empty CHAR and VARCHAR fields as NULL.
	//
	// A value of `true` sets empty CHAR and VARCHAR fields to null. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-emptyasnull
	//
	EmptyAsNull interface{} `field:"optional" json:"emptyAsNull" yaml:"emptyAsNull"`
	// The type of server-side encryption that you want to use for your data.
	//
	// This encryption type is part of the endpoint settings or the extra connections attributes for Amazon S3. You can choose either `SSE_S3` (the default) or `SSE_KMS` .
	//
	// > For the `ModifyEndpoint` operation, you can change the existing value of the `EncryptionMode` parameter from `SSE_KMS` to `SSE_S3` . But you can’t change the existing value from `SSE_S3` to `SSE_KMS` .
	//
	// To use `SSE_S3` , create an AWS Identity and Access Management (IAM) role with a policy that allows `"arn:aws:s3:::*"` to use the following actions: `"s3:PutObject", "s3:ListBucket"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-encryptionmode
	//
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// This setting is only valid for a full-load migration task.
	//
	// Set `ExplicitIds` to `true` to have tables with `IDENTITY` columns override their auto-generated values with explicit values loaded from the source data files used to populate the tables. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-explicitids
	//
	ExplicitIds interface{} `field:"optional" json:"explicitIds" yaml:"explicitIds"`
	// The number of threads used to upload a single file.
	//
	// This parameter accepts a value from 1 through 64. It defaults to 10.
	//
	// The number of parallel streams used to upload a single .csv file to an S3 bucket using S3 Multipart Upload. For more information, see [Multipart upload overview](https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html) .
	//
	// `FileTransferUploadStreams` accepts a value from 1 through 64. It defaults to 10.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-filetransferuploadstreams
	//
	FileTransferUploadStreams *float64 `field:"optional" json:"fileTransferUploadStreams" yaml:"fileTransferUploadStreams"`
	// The amount of time to wait (in milliseconds) before timing out of operations performed by AWS DMS on a Redshift cluster, such as Redshift COPY, INSERT, DELETE, and UPDATE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-loadtimeout
	//
	LoadTimeout *float64 `field:"optional" json:"loadTimeout" yaml:"loadTimeout"`
	// When true, lets Redshift migrate the boolean type as boolean.
	//
	// By default, Redshift migrates booleans as `varchar(1)` . You must set this setting on both the source and target endpoints for it to take effect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-mapbooleanasboolean
	//
	MapBooleanAsBoolean interface{} `field:"optional" json:"mapBooleanAsBoolean" yaml:"mapBooleanAsBoolean"`
	// The maximum size (in KB) of any .csv file used to load data on an S3 bucket and transfer data to Amazon Redshift. It defaults to 1048576KB (1 GB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-maxfilesize
	//
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// A value that specifies to remove surrounding quotation marks from strings in the incoming data.
	//
	// All characters within the quotation marks, including delimiters, are retained. Choose `true` to remove quotation marks. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-removequotes
	//
	RemoveQuotes interface{} `field:"optional" json:"removeQuotes" yaml:"removeQuotes"`
	// A value that specifies to replaces the invalid characters specified in `ReplaceInvalidChars` , substituting the specified characters instead.
	//
	// The default is `"?"` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-replacechars
	//
	ReplaceChars *string `field:"optional" json:"replaceChars" yaml:"replaceChars"`
	// A list of characters that you want to replace.
	//
	// Use with `ReplaceChars` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-replaceinvalidchars
	//
	ReplaceInvalidChars *string `field:"optional" json:"replaceInvalidChars" yaml:"replaceInvalidChars"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the Amazon Redshift endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-secretsmanageraccessrolearn
	//
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the Amazon Redshift endpoint connection details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-secretsmanagersecretid
	//
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// The AWS KMS key ID.
	//
	// If you are using `SSE_KMS` for the `EncryptionMode` , provide this key ID. The key that you use needs an attached policy that enables IAM user permissions and allows use of the key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-serversideencryptionkmskeyid
	//
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// The Amazon Resource Name (ARN) of the IAM role that has access to the Amazon Redshift service.
	//
	// The role must allow the `iam:PassRole` action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-serviceaccessrolearn
	//
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// The time format that you want to use.
	//
	// Valid values are `auto` (case-sensitive), `'timeformat_string'` , `'epochsecs'` , or `'epochmillisecs'` . It defaults to 10. Using `auto` recognizes most strings, even some that aren't supported when you use a time format string.
	//
	// If your date and time values use formats different from each other, set this parameter to `auto` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-timeformat
	//
	TimeFormat *string `field:"optional" json:"timeFormat" yaml:"timeFormat"`
	// A value that specifies to remove the trailing white space characters from a VARCHAR string.
	//
	// This parameter applies only to columns with a VARCHAR data type. Choose `true` to remove unneeded white space. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-trimblanks
	//
	TrimBlanks interface{} `field:"optional" json:"trimBlanks" yaml:"trimBlanks"`
	// A value that specifies to truncate data in columns to the appropriate number of characters, so that the data fits in the column.
	//
	// This parameter applies only to columns with a VARCHAR or CHAR data type, and rows with a size of 4 MB or less. Choose `true` to truncate data. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-truncatecolumns
	//
	TruncateColumns interface{} `field:"optional" json:"truncateColumns" yaml:"truncateColumns"`
	// The size (in KB) of the in-memory file write buffer used when generating .csv files on the local disk at the DMS replication instance. The default value is 1000 (buffer size is 1000KB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-redshiftsettings.html#cfn-dms-endpoint-redshiftsettings-writebuffersize
	//
	WriteBufferSize *float64 `field:"optional" json:"writeBufferSize" yaml:"writeBufferSize"`
}

