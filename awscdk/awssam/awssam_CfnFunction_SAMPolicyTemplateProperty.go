package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAMPolicyTemplateProperty := &sAMPolicyTemplateProperty{
//   	amiDescribePolicy: &emptySAMPTProperty{
//   	},
//   	awsSecretsManagerGetSecretValuePolicy: &secretArnSAMPTProperty{
//   		secretArn: jsii.String("secretArn"),
//   	},
//   	cloudFormationDescribeStacksPolicy: &emptySAMPTProperty{
//   	},
//   	cloudWatchPutMetricPolicy: &emptySAMPTProperty{
//   	},
//   	dynamoDbCrudPolicy: &tableSAMPTProperty{
//   		tableName: jsii.String("tableName"),
//   	},
//   	dynamoDbReadPolicy: &tableSAMPTProperty{
//   		tableName: jsii.String("tableName"),
//   	},
//   	dynamoDbStreamReadPolicy: &tableStreamSAMPTProperty{
//   		streamName: jsii.String("streamName"),
//   		tableName: jsii.String("tableName"),
//   	},
//   	dynamoDbWritePolicy: &tableSAMPTProperty{
//   		tableName: jsii.String("tableName"),
//   	},
//   	ec2DescribePolicy: &emptySAMPTProperty{
//   	},
//   	elasticsearchHttpPostPolicy: &domainSAMPTProperty{
//   		domainName: jsii.String("domainName"),
//   	},
//   	filterLogEventsPolicy: &logGroupSAMPTProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	kinesisCrudPolicy: &streamSAMPTProperty{
//   		streamName: jsii.String("streamName"),
//   	},
//   	kinesisStreamReadPolicy: &streamSAMPTProperty{
//   		streamName: jsii.String("streamName"),
//   	},
//   	kmsDecryptPolicy: &keySAMPTProperty{
//   		keyId: jsii.String("keyId"),
//   	},
//   	lambdaInvokePolicy: &functionSAMPTProperty{
//   		functionName: jsii.String("functionName"),
//   	},
//   	rekognitionDetectOnlyPolicy: &emptySAMPTProperty{
//   	},
//   	rekognitionLabelsPolicy: &emptySAMPTProperty{
//   	},
//   	rekognitionNoDataAccessPolicy: &collectionSAMPTProperty{
//   		collectionId: jsii.String("collectionId"),
//   	},
//   	rekognitionReadPolicy: &collectionSAMPTProperty{
//   		collectionId: jsii.String("collectionId"),
//   	},
//   	rekognitionWriteOnlyAccessPolicy: &collectionSAMPTProperty{
//   		collectionId: jsii.String("collectionId"),
//   	},
//   	s3CrudPolicy: &bucketSAMPTProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   	s3ReadPolicy: &bucketSAMPTProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   	s3WritePolicy: &bucketSAMPTProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   	sesBulkTemplatedCrudPolicy: &identitySAMPTProperty{
//   		identityName: jsii.String("identityName"),
//   	},
//   	sesCrudPolicy: &identitySAMPTProperty{
//   		identityName: jsii.String("identityName"),
//   	},
//   	sesEmailTemplateCrudPolicy: &emptySAMPTProperty{
//   	},
//   	sesSendBouncePolicy: &identitySAMPTProperty{
//   		identityName: jsii.String("identityName"),
//   	},
//   	snsCrudPolicy: &topicSAMPTProperty{
//   		topicName: jsii.String("topicName"),
//   	},
//   	snsPublishMessagePolicy: &topicSAMPTProperty{
//   		topicName: jsii.String("topicName"),
//   	},
//   	sqsPollerPolicy: &queueSAMPTProperty{
//   		queueName: jsii.String("queueName"),
//   	},
//   	sqsSendMessagePolicy: &queueSAMPTProperty{
//   		queueName: jsii.String("queueName"),
//   	},
//   	ssmParameterReadPolicy: &parameterNameSAMPTProperty{
//   		parameterName: jsii.String("parameterName"),
//   	},
//   	stepFunctionsExecutionPolicy: &stateMachineSAMPTProperty{
//   		stateMachineName: jsii.String("stateMachineName"),
//   	},
//   	vpcAccessPolicy: &emptySAMPTProperty{
//   	},
//   }
//
type CfnFunction_SAMPolicyTemplateProperty struct {
	// `CfnFunction.SAMPolicyTemplateProperty.AMIDescribePolicy`.
	AmiDescribePolicy interface{} `field:"optional" json:"amiDescribePolicy" yaml:"amiDescribePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.AWSSecretsManagerGetSecretValuePolicy`.
	AwsSecretsManagerGetSecretValuePolicy interface{} `field:"optional" json:"awsSecretsManagerGetSecretValuePolicy" yaml:"awsSecretsManagerGetSecretValuePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.CloudFormationDescribeStacksPolicy`.
	CloudFormationDescribeStacksPolicy interface{} `field:"optional" json:"cloudFormationDescribeStacksPolicy" yaml:"cloudFormationDescribeStacksPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.CloudWatchPutMetricPolicy`.
	CloudWatchPutMetricPolicy interface{} `field:"optional" json:"cloudWatchPutMetricPolicy" yaml:"cloudWatchPutMetricPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBCrudPolicy`.
	DynamoDbCrudPolicy interface{} `field:"optional" json:"dynamoDbCrudPolicy" yaml:"dynamoDbCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBReadPolicy`.
	DynamoDbReadPolicy interface{} `field:"optional" json:"dynamoDbReadPolicy" yaml:"dynamoDbReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBStreamReadPolicy`.
	DynamoDbStreamReadPolicy interface{} `field:"optional" json:"dynamoDbStreamReadPolicy" yaml:"dynamoDbStreamReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBWritePolicy`.
	DynamoDbWritePolicy interface{} `field:"optional" json:"dynamoDbWritePolicy" yaml:"dynamoDbWritePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.EC2DescribePolicy`.
	Ec2DescribePolicy interface{} `field:"optional" json:"ec2DescribePolicy" yaml:"ec2DescribePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.ElasticsearchHttpPostPolicy`.
	ElasticsearchHttpPostPolicy interface{} `field:"optional" json:"elasticsearchHttpPostPolicy" yaml:"elasticsearchHttpPostPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.FilterLogEventsPolicy`.
	FilterLogEventsPolicy interface{} `field:"optional" json:"filterLogEventsPolicy" yaml:"filterLogEventsPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.KinesisCrudPolicy`.
	KinesisCrudPolicy interface{} `field:"optional" json:"kinesisCrudPolicy" yaml:"kinesisCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.KinesisStreamReadPolicy`.
	KinesisStreamReadPolicy interface{} `field:"optional" json:"kinesisStreamReadPolicy" yaml:"kinesisStreamReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.KMSDecryptPolicy`.
	KmsDecryptPolicy interface{} `field:"optional" json:"kmsDecryptPolicy" yaml:"kmsDecryptPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.LambdaInvokePolicy`.
	LambdaInvokePolicy interface{} `field:"optional" json:"lambdaInvokePolicy" yaml:"lambdaInvokePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionDetectOnlyPolicy`.
	RekognitionDetectOnlyPolicy interface{} `field:"optional" json:"rekognitionDetectOnlyPolicy" yaml:"rekognitionDetectOnlyPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionLabelsPolicy`.
	RekognitionLabelsPolicy interface{} `field:"optional" json:"rekognitionLabelsPolicy" yaml:"rekognitionLabelsPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionNoDataAccessPolicy`.
	RekognitionNoDataAccessPolicy interface{} `field:"optional" json:"rekognitionNoDataAccessPolicy" yaml:"rekognitionNoDataAccessPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionReadPolicy`.
	RekognitionReadPolicy interface{} `field:"optional" json:"rekognitionReadPolicy" yaml:"rekognitionReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionWriteOnlyAccessPolicy`.
	RekognitionWriteOnlyAccessPolicy interface{} `field:"optional" json:"rekognitionWriteOnlyAccessPolicy" yaml:"rekognitionWriteOnlyAccessPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.S3CrudPolicy`.
	S3CrudPolicy interface{} `field:"optional" json:"s3CrudPolicy" yaml:"s3CrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.S3ReadPolicy`.
	S3ReadPolicy interface{} `field:"optional" json:"s3ReadPolicy" yaml:"s3ReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.S3WritePolicy`.
	S3WritePolicy interface{} `field:"optional" json:"s3WritePolicy" yaml:"s3WritePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESBulkTemplatedCrudPolicy`.
	SesBulkTemplatedCrudPolicy interface{} `field:"optional" json:"sesBulkTemplatedCrudPolicy" yaml:"sesBulkTemplatedCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESCrudPolicy`.
	SesCrudPolicy interface{} `field:"optional" json:"sesCrudPolicy" yaml:"sesCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESEmailTemplateCrudPolicy`.
	SesEmailTemplateCrudPolicy interface{} `field:"optional" json:"sesEmailTemplateCrudPolicy" yaml:"sesEmailTemplateCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESSendBouncePolicy`.
	SesSendBouncePolicy interface{} `field:"optional" json:"sesSendBouncePolicy" yaml:"sesSendBouncePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SNSCrudPolicy`.
	SnsCrudPolicy interface{} `field:"optional" json:"snsCrudPolicy" yaml:"snsCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SNSPublishMessagePolicy`.
	SnsPublishMessagePolicy interface{} `field:"optional" json:"snsPublishMessagePolicy" yaml:"snsPublishMessagePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SQSPollerPolicy`.
	SqsPollerPolicy interface{} `field:"optional" json:"sqsPollerPolicy" yaml:"sqsPollerPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SQSSendMessagePolicy`.
	SqsSendMessagePolicy interface{} `field:"optional" json:"sqsSendMessagePolicy" yaml:"sqsSendMessagePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SSMParameterReadPolicy`.
	SsmParameterReadPolicy interface{} `field:"optional" json:"ssmParameterReadPolicy" yaml:"ssmParameterReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.StepFunctionsExecutionPolicy`.
	StepFunctionsExecutionPolicy interface{} `field:"optional" json:"stepFunctionsExecutionPolicy" yaml:"stepFunctionsExecutionPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.VPCAccessPolicy`.
	VpcAccessPolicy interface{} `field:"optional" json:"vpcAccessPolicy" yaml:"vpcAccessPolicy"`
}

