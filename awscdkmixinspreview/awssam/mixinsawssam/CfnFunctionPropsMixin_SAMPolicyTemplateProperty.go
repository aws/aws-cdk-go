package mixinsawssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sAMPolicyTemplateProperty := &SAMPolicyTemplateProperty{
//   	AmiDescribePolicy: &EmptySAMPTProperty{
//   	},
//   	AwsSecretsManagerGetSecretValuePolicy: &SecretArnSAMPTProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	CloudFormationDescribeStacksPolicy: &EmptySAMPTProperty{
//   	},
//   	CloudWatchPutMetricPolicy: &EmptySAMPTProperty{
//   	},
//   	DynamoDbCrudPolicy: &TableSAMPTProperty{
//   		TableName: jsii.String("tableName"),
//   	},
//   	DynamoDbReadPolicy: &TableSAMPTProperty{
//   		TableName: jsii.String("tableName"),
//   	},
//   	DynamoDbStreamReadPolicy: &TableStreamSAMPTProperty{
//   		StreamName: jsii.String("streamName"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	DynamoDbWritePolicy: &TableSAMPTProperty{
//   		TableName: jsii.String("tableName"),
//   	},
//   	Ec2DescribePolicy: &EmptySAMPTProperty{
//   	},
//   	ElasticsearchHttpPostPolicy: &DomainSAMPTProperty{
//   		DomainName: jsii.String("domainName"),
//   	},
//   	FilterLogEventsPolicy: &LogGroupSAMPTProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   	},
//   	KinesisCrudPolicy: &StreamSAMPTProperty{
//   		StreamName: jsii.String("streamName"),
//   	},
//   	KinesisStreamReadPolicy: &StreamSAMPTProperty{
//   		StreamName: jsii.String("streamName"),
//   	},
//   	KmsDecryptPolicy: &KeySAMPTProperty{
//   		KeyId: jsii.String("keyId"),
//   	},
//   	LambdaInvokePolicy: &FunctionSAMPTProperty{
//   		FunctionName: jsii.String("functionName"),
//   	},
//   	RekognitionDetectOnlyPolicy: &EmptySAMPTProperty{
//   	},
//   	RekognitionLabelsPolicy: &EmptySAMPTProperty{
//   	},
//   	RekognitionNoDataAccessPolicy: &CollectionSAMPTProperty{
//   		CollectionId: jsii.String("collectionId"),
//   	},
//   	RekognitionReadPolicy: &CollectionSAMPTProperty{
//   		CollectionId: jsii.String("collectionId"),
//   	},
//   	RekognitionWriteOnlyAccessPolicy: &CollectionSAMPTProperty{
//   		CollectionId: jsii.String("collectionId"),
//   	},
//   	S3CrudPolicy: &BucketSAMPTProperty{
//   		BucketName: jsii.String("bucketName"),
//   	},
//   	S3ReadPolicy: &BucketSAMPTProperty{
//   		BucketName: jsii.String("bucketName"),
//   	},
//   	S3WritePolicy: &BucketSAMPTProperty{
//   		BucketName: jsii.String("bucketName"),
//   	},
//   	SesBulkTemplatedCrudPolicy: &IdentitySAMPTProperty{
//   		IdentityName: jsii.String("identityName"),
//   	},
//   	SesCrudPolicy: &IdentitySAMPTProperty{
//   		IdentityName: jsii.String("identityName"),
//   	},
//   	SesEmailTemplateCrudPolicy: &EmptySAMPTProperty{
//   	},
//   	SesSendBouncePolicy: &IdentitySAMPTProperty{
//   		IdentityName: jsii.String("identityName"),
//   	},
//   	SnsCrudPolicy: &TopicSAMPTProperty{
//   		TopicName: jsii.String("topicName"),
//   	},
//   	SnsPublishMessagePolicy: &TopicSAMPTProperty{
//   		TopicName: jsii.String("topicName"),
//   	},
//   	SqsPollerPolicy: &QueueSAMPTProperty{
//   		QueueName: jsii.String("queueName"),
//   	},
//   	SqsSendMessagePolicy: &QueueSAMPTProperty{
//   		QueueName: jsii.String("queueName"),
//   	},
//   	SsmParameterReadPolicy: &ParameterNameSAMPTProperty{
//   		ParameterName: jsii.String("parameterName"),
//   	},
//   	StepFunctionsExecutionPolicy: &StateMachineSAMPTProperty{
//   		StateMachineName: jsii.String("stateMachineName"),
//   	},
//   	VpcAccessPolicy: &EmptySAMPTProperty{
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html
//
type CfnFunctionPropsMixin_SAMPolicyTemplateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-amidescribepolicy
	//
	AmiDescribePolicy interface{} `field:"optional" json:"amiDescribePolicy" yaml:"amiDescribePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-awssecretsmanagergetsecretvaluepolicy
	//
	AwsSecretsManagerGetSecretValuePolicy interface{} `field:"optional" json:"awsSecretsManagerGetSecretValuePolicy" yaml:"awsSecretsManagerGetSecretValuePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-cloudformationdescribestackspolicy
	//
	CloudFormationDescribeStacksPolicy interface{} `field:"optional" json:"cloudFormationDescribeStacksPolicy" yaml:"cloudFormationDescribeStacksPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-cloudwatchputmetricpolicy
	//
	CloudWatchPutMetricPolicy interface{} `field:"optional" json:"cloudWatchPutMetricPolicy" yaml:"cloudWatchPutMetricPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-dynamodbcrudpolicy
	//
	DynamoDbCrudPolicy interface{} `field:"optional" json:"dynamoDbCrudPolicy" yaml:"dynamoDbCrudPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-dynamodbreadpolicy
	//
	DynamoDbReadPolicy interface{} `field:"optional" json:"dynamoDbReadPolicy" yaml:"dynamoDbReadPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-dynamodbstreamreadpolicy
	//
	DynamoDbStreamReadPolicy interface{} `field:"optional" json:"dynamoDbStreamReadPolicy" yaml:"dynamoDbStreamReadPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-dynamodbwritepolicy
	//
	DynamoDbWritePolicy interface{} `field:"optional" json:"dynamoDbWritePolicy" yaml:"dynamoDbWritePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-ec2describepolicy
	//
	Ec2DescribePolicy interface{} `field:"optional" json:"ec2DescribePolicy" yaml:"ec2DescribePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-elasticsearchhttppostpolicy
	//
	ElasticsearchHttpPostPolicy interface{} `field:"optional" json:"elasticsearchHttpPostPolicy" yaml:"elasticsearchHttpPostPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-filterlogeventspolicy
	//
	FilterLogEventsPolicy interface{} `field:"optional" json:"filterLogEventsPolicy" yaml:"filterLogEventsPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-kinesiscrudpolicy
	//
	KinesisCrudPolicy interface{} `field:"optional" json:"kinesisCrudPolicy" yaml:"kinesisCrudPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-kinesisstreamreadpolicy
	//
	KinesisStreamReadPolicy interface{} `field:"optional" json:"kinesisStreamReadPolicy" yaml:"kinesisStreamReadPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-kmsdecryptpolicy
	//
	KmsDecryptPolicy interface{} `field:"optional" json:"kmsDecryptPolicy" yaml:"kmsDecryptPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-lambdainvokepolicy
	//
	LambdaInvokePolicy interface{} `field:"optional" json:"lambdaInvokePolicy" yaml:"lambdaInvokePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-rekognitiondetectonlypolicy
	//
	RekognitionDetectOnlyPolicy interface{} `field:"optional" json:"rekognitionDetectOnlyPolicy" yaml:"rekognitionDetectOnlyPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-rekognitionlabelspolicy
	//
	RekognitionLabelsPolicy interface{} `field:"optional" json:"rekognitionLabelsPolicy" yaml:"rekognitionLabelsPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-rekognitionnodataaccesspolicy
	//
	RekognitionNoDataAccessPolicy interface{} `field:"optional" json:"rekognitionNoDataAccessPolicy" yaml:"rekognitionNoDataAccessPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-rekognitionreadpolicy
	//
	RekognitionReadPolicy interface{} `field:"optional" json:"rekognitionReadPolicy" yaml:"rekognitionReadPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-rekognitionwriteonlyaccesspolicy
	//
	RekognitionWriteOnlyAccessPolicy interface{} `field:"optional" json:"rekognitionWriteOnlyAccessPolicy" yaml:"rekognitionWriteOnlyAccessPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-s3crudpolicy
	//
	S3CrudPolicy interface{} `field:"optional" json:"s3CrudPolicy" yaml:"s3CrudPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-s3readpolicy
	//
	S3ReadPolicy interface{} `field:"optional" json:"s3ReadPolicy" yaml:"s3ReadPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-s3writepolicy
	//
	S3WritePolicy interface{} `field:"optional" json:"s3WritePolicy" yaml:"s3WritePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-sesbulktemplatedcrudpolicy
	//
	SesBulkTemplatedCrudPolicy interface{} `field:"optional" json:"sesBulkTemplatedCrudPolicy" yaml:"sesBulkTemplatedCrudPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-sescrudpolicy
	//
	SesCrudPolicy interface{} `field:"optional" json:"sesCrudPolicy" yaml:"sesCrudPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-sesemailtemplatecrudpolicy
	//
	SesEmailTemplateCrudPolicy interface{} `field:"optional" json:"sesEmailTemplateCrudPolicy" yaml:"sesEmailTemplateCrudPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-sessendbouncepolicy
	//
	SesSendBouncePolicy interface{} `field:"optional" json:"sesSendBouncePolicy" yaml:"sesSendBouncePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-snscrudpolicy
	//
	SnsCrudPolicy interface{} `field:"optional" json:"snsCrudPolicy" yaml:"snsCrudPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-snspublishmessagepolicy
	//
	SnsPublishMessagePolicy interface{} `field:"optional" json:"snsPublishMessagePolicy" yaml:"snsPublishMessagePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-sqspollerpolicy
	//
	SqsPollerPolicy interface{} `field:"optional" json:"sqsPollerPolicy" yaml:"sqsPollerPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-sqssendmessagepolicy
	//
	SqsSendMessagePolicy interface{} `field:"optional" json:"sqsSendMessagePolicy" yaml:"sqsSendMessagePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-ssmparameterreadpolicy
	//
	SsmParameterReadPolicy interface{} `field:"optional" json:"ssmParameterReadPolicy" yaml:"ssmParameterReadPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-stepfunctionsexecutionpolicy
	//
	StepFunctionsExecutionPolicy interface{} `field:"optional" json:"stepFunctionsExecutionPolicy" yaml:"stepFunctionsExecutionPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-sampolicytemplate.html#cfn-serverless-function-sampolicytemplate-vpcaccesspolicy
	//
	VpcAccessPolicy interface{} `field:"optional" json:"vpcAccessPolicy" yaml:"vpcAccessPolicy"`
}

