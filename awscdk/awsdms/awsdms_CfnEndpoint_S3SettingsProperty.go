package awsdms


// Provides information that defines an Amazon S3 endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about the available settings, see [Extra connection attributes when using Amazon S3 as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.S3.html#CHAP_Source.S3.Configuring) and [Extra connection attributes when using Amazon S3 as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SettingsProperty := &s3SettingsProperty{
//   	addColumnName: jsii.Boolean(false),
//   	bucketFolder: jsii.String("bucketFolder"),
//   	bucketName: jsii.String("bucketName"),
//   	cannedAclForObjects: jsii.String("cannedAclForObjects"),
//   	cdcInsertsAndUpdates: jsii.Boolean(false),
//   	cdcInsertsOnly: jsii.Boolean(false),
//   	cdcMaxBatchInterval: jsii.Number(123),
//   	cdcMinFileSize: jsii.Number(123),
//   	cdcPath: jsii.String("cdcPath"),
//   	compressionType: jsii.String("compressionType"),
//   	csvDelimiter: jsii.String("csvDelimiter"),
//   	csvNoSupValue: jsii.String("csvNoSupValue"),
//   	csvNullValue: jsii.String("csvNullValue"),
//   	csvRowDelimiter: jsii.String("csvRowDelimiter"),
//   	dataFormat: jsii.String("dataFormat"),
//   	dataPageSize: jsii.Number(123),
//   	datePartitionDelimiter: jsii.String("datePartitionDelimiter"),
//   	datePartitionEnabled: jsii.Boolean(false),
//   	datePartitionSequence: jsii.String("datePartitionSequence"),
//   	datePartitionTimezone: jsii.String("datePartitionTimezone"),
//   	dictPageSizeLimit: jsii.Number(123),
//   	enableStatistics: jsii.Boolean(false),
//   	encodingType: jsii.String("encodingType"),
//   	encryptionMode: jsii.String("encryptionMode"),
//   	externalTableDefinition: jsii.String("externalTableDefinition"),
//   	ignoreHeaderRows: jsii.Number(123),
//   	includeOpForFullLoad: jsii.Boolean(false),
//   	maxFileSize: jsii.Number(123),
//   	parquetTimestampInMillisecond: jsii.Boolean(false),
//   	parquetVersion: jsii.String("parquetVersion"),
//   	preserveTransactions: jsii.Boolean(false),
//   	rfc4180: jsii.Boolean(false),
//   	rowGroupLength: jsii.Number(123),
//   	serverSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   	serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	timestampColumnName: jsii.String("timestampColumnName"),
//   	useCsvNoSupValue: jsii.Boolean(false),
//   	useTaskStartTimeForFullLoadTimestamp: jsii.Boolean(false),
//   }
//
type CfnEndpoint_S3SettingsProperty struct {
	// An optional parameter that, when set to `true` or `y` , you can use to add column name information to the .csv output file.
	//
	// The default value is `false` . Valid values are `true` , `false` , `y` , and `n` .
	AddColumnName interface{} `field:"optional" json:"addColumnName" yaml:"addColumnName"`
	// An optional parameter to set a folder name in the S3 bucket.
	//
	// If provided, tables are created in the path `*bucketFolder* / *schema_name* / *table_name* /` . If this parameter isn't specified, the path used is `*schema_name* / *table_name* /` .
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// The name of the S3 bucket.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// A value that enables AWS DMS to specify a predefined (canned) access control list (ACL) for objects created in an Amazon S3 bucket as .csv or .parquet files. For more information about Amazon S3 canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 Developer Guide* .
	//
	// The default value is NONE. Valid values include NONE, PRIVATE, PUBLIC_READ, PUBLIC_READ_WRITE, AUTHENTICATED_READ, AWS_EXEC_READ, BUCKET_OWNER_READ, and BUCKET_OWNER_FULL_CONTROL.
	CannedAclForObjects *string `field:"optional" json:"cannedAclForObjects" yaml:"cannedAclForObjects"`
	// A value that enables a change data capture (CDC) load to write INSERT and UPDATE operations to .csv or .parquet (columnar storage) output files. The default setting is `false` , but when `CdcInsertsAndUpdates` is set to `true` or `y` , only INSERTs and UPDATEs from the source database are migrated to the .csv or .parquet file.
	//
	// For .csv file format only, how these INSERTs and UPDATEs are recorded depends on the value of the `IncludeOpForFullLoad` parameter. If `IncludeOpForFullLoad` is set to `true` , the first field of every CDC record is set to either `I` or `U` to indicate INSERT and UPDATE operations at the source. But if `IncludeOpForFullLoad` is set to `false` , CDC records are written without an indication of INSERT or UPDATE operations at the source. For more information about how these settings work together, see [Indicating Source DB Operations in Migrated S3 Data](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring.InsertOps) in the *AWS Database Migration Service User Guide* .
	//
	// > AWS DMS supports the use of the `CdcInsertsAndUpdates` parameter in versions 3.3.1 and later.
	// >
	// > `CdcInsertsOnly` and `CdcInsertsAndUpdates` can't both be set to `true` for the same endpoint. Set either `CdcInsertsOnly` or `CdcInsertsAndUpdates` to `true` for the same endpoint, but not both.
	CdcInsertsAndUpdates interface{} `field:"optional" json:"cdcInsertsAndUpdates" yaml:"cdcInsertsAndUpdates"`
	// A value that enables a change data capture (CDC) load to write only INSERT operations to .csv or columnar storage (.parquet) output files. By default (the `false` setting), the first field in a .csv or .parquet record contains the letter I (INSERT), U (UPDATE), or D (DELETE). These values indicate whether the row was inserted, updated, or deleted at the source database for a CDC load to the target.
	//
	// If `CdcInsertsOnly` is set to `true` or `y` , only INSERTs from the source database are migrated to the .csv or .parquet file. For .csv format only, how these INSERTs are recorded depends on the value of `IncludeOpForFullLoad` . If `IncludeOpForFullLoad` is set to `true` , the first field of every CDC record is set to I to indicate the INSERT operation at the source. If `IncludeOpForFullLoad` is set to `false` , every CDC record is written without a first field to indicate the INSERT operation at the source. For more information about how these settings work together, see [Indicating Source DB Operations in Migrated S3 Data](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring.InsertOps) in the *AWS Database Migration Service User Guide* .
	//
	// > AWS DMS supports the interaction described preceding between the `CdcInsertsOnly` and `IncludeOpForFullLoad` parameters in versions 3.1.4 and later.
	// >
	// > `CdcInsertsOnly` and `CdcInsertsAndUpdates` can't both be set to `true` for the same endpoint. Set either `CdcInsertsOnly` or `CdcInsertsAndUpdates` to `true` for the same endpoint, but not both.
	CdcInsertsOnly interface{} `field:"optional" json:"cdcInsertsOnly" yaml:"cdcInsertsOnly"`
	// Maximum length of the interval, defined in seconds, after which to output a file to Amazon S3.
	//
	// When `CdcMaxBatchInterval` and `CdcMinFileSize` are both specified, the file write is triggered by whichever parameter condition is met first within an AWS DMS CloudFormation template.
	//
	// The default value is 60 seconds.
	CdcMaxBatchInterval *float64 `field:"optional" json:"cdcMaxBatchInterval" yaml:"cdcMaxBatchInterval"`
	// Minimum file size, defined in megabytes, to reach for a file output to Amazon S3.
	//
	// When `CdcMinFileSize` and `CdcMaxBatchInterval` are both specified, the file write is triggered by whichever parameter condition is met first within an AWS DMS CloudFormation template.
	//
	// The default value is 32 MB.
	CdcMinFileSize *float64 `field:"optional" json:"cdcMinFileSize" yaml:"cdcMinFileSize"`
	// Specifies the folder path of CDC files.
	//
	// For an S3 source, this setting is required if a task captures change data; otherwise, it's optional. If `CdcPath` is set, AWS DMS reads CDC files from this path and replicates the data changes to the target endpoint. For an S3 target if you set [`PreserveTransactions`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-PreserveTransactions) to `true` , AWS DMS verifies that you have set this parameter to a folder path on your S3 target where AWS DMS can save the transaction order for the CDC load. AWS DMS creates this CDC folder path in either your S3 target working directory or the S3 target location specified by [`BucketFolder`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-BucketFolder) and [`BucketName`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-BucketName) .
	//
	// For example, if you specify `CdcPath` as `MyChangedData` , and you specify `BucketName` as `MyTargetBucket` but do not specify `BucketFolder` , AWS DMS creates the CDC folder path following: `MyTargetBucket/MyChangedData` .
	//
	// If you specify the same `CdcPath` , and you specify `BucketName` as `MyTargetBucket` and `BucketFolder` as `MyTargetData` , AWS DMS creates the CDC folder path following: `MyTargetBucket/MyTargetData/MyChangedData` .
	//
	// For more information on CDC including transaction order on an S3 target, see [Capturing data changes (CDC) including transaction order on the S3 target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.EndpointSettings.CdcPath) .
	//
	// > This setting is supported in AWS DMS versions 3.4.2 and later.
	CdcPath *string `field:"optional" json:"cdcPath" yaml:"cdcPath"`
	// An optional parameter.
	//
	// When set to GZIP it enables the service to compress the target files. To allow the service to write the target files uncompressed, either set this parameter to NONE (the default) or don't specify the parameter at all. This parameter applies to both .csv and .parquet file formats.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The delimiter used to separate columns in the .csv file for both source and target. The default is a comma.
	CsvDelimiter *string `field:"optional" json:"csvDelimiter" yaml:"csvDelimiter"`
	// This setting only applies if your Amazon S3 output files during a change data capture (CDC) load are written in .csv format. If [`UseCsvNoSupValue`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-UseCsvNoSupValue) is set to true, specify a string value that you want AWS DMS to use for all columns not included in the supplemental log. If you do not specify a string value, AWS DMS uses the null value for these columns regardless of the `UseCsvNoSupValue` setting.
	//
	// > This setting is supported in AWS DMS versions 3.4.1 and later.
	CsvNoSupValue *string `field:"optional" json:"csvNoSupValue" yaml:"csvNoSupValue"`
	// An optional parameter that specifies how AWS DMS treats null values.
	//
	// While handling the null value, you can use this parameter to pass a user-defined string as null when writing to the target. For example, when target columns are not nullable, you can use this option to differentiate between the empty string value and the null value. So, if you set this parameter value to the empty string ("" or ''), AWS DMS treats the empty string as the null value instead of `NULL` .
	//
	// The default value is `NULL` . Valid values include any valid string.
	CsvNullValue *string `field:"optional" json:"csvNullValue" yaml:"csvNullValue"`
	// The delimiter used to separate rows in the .csv file for both source and target.
	//
	// The default is a carriage return ( `\n` ).
	CsvRowDelimiter *string `field:"optional" json:"csvRowDelimiter" yaml:"csvRowDelimiter"`
	// The format of the data that you want to use for output. You can choose one of the following:.
	//
	// - `csv` : This is a row-based file format with comma-separated values (.csv).
	// - `parquet` : Apache Parquet (.parquet) is a columnar storage file format that features efficient compression and provides faster query response.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// The size of one data page in bytes.
	//
	// This parameter defaults to 1024 * 1024 bytes (1 MiB). This number is used for .parquet file format only.
	DataPageSize *float64 `field:"optional" json:"dataPageSize" yaml:"dataPageSize"`
	// Specifies a date separating delimiter to use during folder partitioning.
	//
	// The default value is `SLASH` . Use this parameter when `DatePartitionedEnabled` is set to `true` .
	DatePartitionDelimiter *string `field:"optional" json:"datePartitionDelimiter" yaml:"datePartitionDelimiter"`
	// When set to `true` , this parameter partitions S3 bucket folders based on transaction commit dates.
	//
	// The default value is `false` . For more information about date-based folder partitioning, see [Using date-based folder partitioning](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.DatePartitioning) .
	DatePartitionEnabled interface{} `field:"optional" json:"datePartitionEnabled" yaml:"datePartitionEnabled"`
	// Identifies the sequence of the date format to use during folder partitioning.
	//
	// The default value is `YYYYMMDD` . Use this parameter when `DatePartitionedEnabled` is set to `true` .
	DatePartitionSequence *string `field:"optional" json:"datePartitionSequence" yaml:"datePartitionSequence"`
	// When creating an S3 target endpoint, set `DatePartitionTimezone` to convert the current UTC time into a specified time zone.
	//
	// The conversion occurs when a date partition folder is created and a change data capture (CDC) file name is generated. The time zone format is Area/Location. Use this parameter when `DatePartitionedEnabled` is set to `true` , as shown in the following example.
	//
	// `s3-settings='{"DatePartitionEnabled": true, "DatePartitionSequence": "YYYYMMDDHH", "DatePartitionDelimiter": "SLASH", "DatePartitionTimezone":" *Asia/Seoul* ", "BucketName": "dms-nattarat-test"}'`.
	DatePartitionTimezone *string `field:"optional" json:"datePartitionTimezone" yaml:"datePartitionTimezone"`
	// The maximum size of an encoded dictionary page of a column.
	//
	// If the dictionary page exceeds this, this column is stored using an encoding type of `PLAIN` . This parameter defaults to 1024 * 1024 bytes (1 MiB), the maximum size of a dictionary page before it reverts to `PLAIN` encoding. This size is used for .parquet file format only.
	DictPageSizeLimit *float64 `field:"optional" json:"dictPageSizeLimit" yaml:"dictPageSizeLimit"`
	// A value that enables statistics for Parquet pages and row groups.
	//
	// Choose `true` to enable statistics, `false` to disable. Statistics include `NULL` , `DISTINCT` , `MAX` , and `MIN` values. This parameter defaults to `true` . This value is used for .parquet file format only.
	EnableStatistics interface{} `field:"optional" json:"enableStatistics" yaml:"enableStatistics"`
	// The type of encoding that you're using:.
	//
	// - `RLE_DICTIONARY` uses a combination of bit-packing and run-length encoding to store repeated values more efficiently. This is the default.
	// - `PLAIN` doesn't use encoding at all. Values are stored as they are.
	// - `PLAIN_DICTIONARY` builds a dictionary of the values encountered in a given column. The dictionary is stored in a dictionary page for each column chunk.
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// The type of server-side encryption that you want to use for your data.
	//
	// This encryption type is part of the endpoint settings or the extra connections attributes for Amazon S3. You can choose either `SSE_S3` (the default) or `SSE_KMS` .
	//
	// > For the `ModifyEndpoint` operation, you can change the existing value of the `EncryptionMode` parameter from `SSE_KMS` to `SSE_S3` . But you can’t change the existing value from `SSE_S3` to `SSE_KMS` .
	//
	// To use `SSE_S3` , you need an IAM role with permission to allow `"arn:aws:s3:::dms-*"` to use the following actions:
	//
	// - `s3:CreateBucket`
	// - `s3:ListBucket`
	// - `s3:DeleteBucket`
	// - `s3:GetBucketLocation`
	// - `s3:GetObject`
	// - `s3:PutObject`
	// - `s3:DeleteObject`
	// - `s3:GetObjectVersion`
	// - `s3:GetBucketPolicy`
	// - `s3:PutBucketPolicy`
	// - `s3:DeleteBucketPolicy`.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// The external table definition.
	//
	// Conditional: If `S3` is used as a source then `ExternalTableDefinition` is required.
	ExternalTableDefinition *string `field:"optional" json:"externalTableDefinition" yaml:"externalTableDefinition"`
	// When this value is set to 1, AWS DMS ignores the first row header in a .csv file. A value of 1 turns on the feature; a value of 0 turns off the feature.
	//
	// The default is 0.
	IgnoreHeaderRows *float64 `field:"optional" json:"ignoreHeaderRows" yaml:"ignoreHeaderRows"`
	// A value that enables a full load to write INSERT operations to the comma-separated value (.csv) output files only to indicate how the rows were added to the source database.
	//
	// > AWS DMS supports the `IncludeOpForFullLoad` parameter in versions 3.1.4 and later.
	//
	// For full load, records can only be inserted. By default (the `false` setting), no information is recorded in these output files for a full load to indicate that the rows were inserted at the source database. If `IncludeOpForFullLoad` is set to `true` or `y` , the INSERT is recorded as an I annotation in the first field of the .csv file. This allows the format of your target records from a full load to be consistent with the target records from a CDC load.
	//
	// > This setting works together with the `CdcInsertsOnly` and the `CdcInsertsAndUpdates` parameters for output to .csv files only. For more information about how these settings work together, see [Indicating Source DB Operations in Migrated S3 Data](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring.InsertOps) in the *AWS Database Migration Service User Guide* .
	IncludeOpForFullLoad interface{} `field:"optional" json:"includeOpForFullLoad" yaml:"includeOpForFullLoad"`
	// A value that specifies the maximum size (in KB) of any .csv file to be created while migrating to an S3 target during full load.
	//
	// The default value is 1,048,576 KB (1 GB). Valid values include 1 to 1,048,576.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// A value that specifies the precision of any `TIMESTAMP` column values that are written to an Amazon S3 object file in .parquet format.
	//
	// > AWS DMS supports the `ParquetTimestampInMillisecond` parameter in versions 3.1.4 and later.
	//
	// When `ParquetTimestampInMillisecond` is set to `true` or `y` , AWS DMS writes all `TIMESTAMP` columns in a .parquet formatted file with millisecond precision. Otherwise, DMS writes them with microsecond precision.
	//
	// Currently, Amazon Athena and AWS Glue can handle only millisecond precision for `TIMESTAMP` values. Set this parameter to `true` for S3 endpoint object files that are .parquet formatted only if you plan to query or process the data with Athena or AWS Glue .
	//
	// > AWS DMS writes any `TIMESTAMP` column values written to an S3 file in .csv format with microsecond precision.
	// >
	// > Setting `ParquetTimestampInMillisecond` has no effect on the string format of the timestamp column value that is inserted by setting the `TimestampColumnName` parameter.
	ParquetTimestampInMillisecond interface{} `field:"optional" json:"parquetTimestampInMillisecond" yaml:"parquetTimestampInMillisecond"`
	// The version of the Apache Parquet format that you want to use: `parquet_1_0` (the default) or `parquet_2_0` .
	ParquetVersion *string `field:"optional" json:"parquetVersion" yaml:"parquetVersion"`
	// If this setting is set to `true` , AWS DMS saves the transaction order for a change data capture (CDC) load on the Amazon S3 target specified by [`CdcPath`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-CdcPath) . For more information, see [Capturing data changes (CDC) including transaction order on the S3 target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.EndpointSettings.CdcPath) .
	//
	// > This setting is supported in AWS DMS versions 3.4.2 and later.
	PreserveTransactions interface{} `field:"optional" json:"preserveTransactions" yaml:"preserveTransactions"`
	// For an S3 source, when this value is set to `true` or `y` , each leading double quotation mark has to be followed by an ending double quotation mark.
	//
	// This formatting complies with RFC 4180. When this value is set to `false` or `n` , string literals are copied to the target as is. In this case, a delimiter (row or column) signals the end of the field. Thus, you can't use a delimiter as part of the string, because it signals the end of the value.
	//
	// For an S3 target, an optional parameter used to set behavior to comply with RFC 4180 for data migrated to Amazon S3 using .csv file format only. When this value is set to `true` or `y` using Amazon S3 as a target, if the data has quotation marks or newline characters in it, AWS DMS encloses the entire column with an additional pair of double quotation marks ("). Every quotation mark within the data is repeated twice.
	//
	// The default value is `true` . Valid values include `true` , `false` , `y` , and `n` .
	Rfc4180 interface{} `field:"optional" json:"rfc4180" yaml:"rfc4180"`
	// The number of rows in a row group.
	//
	// A smaller row group size provides faster reads. But as the number of row groups grows, the slower writes become. This parameter defaults to 10,000 rows. This number is used for .parquet file format only.
	//
	// If you choose a value larger than the maximum, `RowGroupLength` is set to the max row group length in bytes (64 * 1024 * 1024).
	RowGroupLength *float64 `field:"optional" json:"rowGroupLength" yaml:"rowGroupLength"`
	// If you are using `SSE_KMS` for the `EncryptionMode` , provide the AWS KMS key ID.
	//
	// The key that you use needs an attached policy that enables IAM user permissions and allows use of the key.
	//
	// Here is a CLI example: `aws dms create-endpoint --endpoint-identifier *value* --endpoint-type target --engine-name s3 --s3-settings ServiceAccessRoleArn= *value* ,BucketFolder= *value* ,BucketName= *value* ,EncryptionMode=SSE_KMS,ServerSideEncryptionKmsKeyId= *value*`.
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// A required parameter that specifies the Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the `iam:PassRole` action. It enables AWS DMS to read and write objects from an S3 bucket.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// A value that when nonblank causes AWS DMS to add a column with timestamp information to the endpoint data for an Amazon S3 target.
	//
	// > AWS DMS supports the `TimestampColumnName` parameter in versions 3.1.4 and later.
	//
	// AWS DMS includes an additional `STRING` column in the .csv or .parquet object files of your migrated data when you set `TimestampColumnName` to a nonblank value.
	//
	// For a full load, each row of this timestamp column contains a timestamp for when the data was transferred from the source to the target by DMS.
	//
	// For a change data capture (CDC) load, each row of the timestamp column contains the timestamp for the commit of that row in the source database.
	//
	// The string format for this timestamp column value is `yyyy-MM-dd HH:mm:ss.SSSSSS` . By default, the precision of this value is in microseconds. For a CDC load, the rounding of the precision depends on the commit timestamp supported by DMS for the source database.
	//
	// When the `AddColumnName` parameter is set to `true` , DMS also includes a name for the timestamp column that you set with `TimestampColumnName` .
	TimestampColumnName *string `field:"optional" json:"timestampColumnName" yaml:"timestampColumnName"`
	// This setting applies if the S3 output files during a change data capture (CDC) load are written in .csv format. If this setting is set to `true` for columns not included in the supplemental log, AWS DMS uses the value specified by [`CsvNoSupValue`](https://docs.aws.amazon.com/dms/latest/APIReference/API_S3Settings.html#DMS-Type-S3Settings-CsvNoSupValue) . If this setting isn't set or is set to `false` , AWS DMS uses the null value for these columns.
	//
	// > This setting is supported in AWS DMS versions 3.4.1 and later.
	UseCsvNoSupValue interface{} `field:"optional" json:"useCsvNoSupValue" yaml:"useCsvNoSupValue"`
	// When set to true, this parameter uses the task start time as the timestamp column value instead of the time data is written to target.
	//
	// For full load, when `useTaskStartTimeForFullLoadTimestamp` is set to `true` , each row of the timestamp column contains the task start time. For CDC loads, each row of the timestamp column contains the transaction commit time.
	//
	// When `useTaskStartTimeForFullLoadTimestamp` is set to `false` , the full load timestamp in the timestamp column increments with the time data arrives at the target.
	UseTaskStartTimeForFullLoadTimestamp interface{} `field:"optional" json:"useTaskStartTimeForFullLoadTimestamp" yaml:"useTaskStartTimeForFullLoadTimestamp"`
}

