package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceSelfManagedKafkaParametersProperty := &PipeSourceSelfManagedKafkaParametersProperty{
//   	TopicName: jsii.String("topicName"),
//
//   	// the properties below are optional
//   	AdditionalBootstrapServers: []*string{
//   		jsii.String("additionalBootstrapServers"),
//   	},
//   	BatchSize: jsii.Number(123),
//   	ConsumerGroupId: jsii.String("consumerGroupId"),
//   	Credentials: &SelfManagedKafkaAccessConfigurationCredentialsProperty{
//   		BasicAuth: jsii.String("basicAuth"),
//   		ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   		SaslScram256Auth: jsii.String("saslScram256Auth"),
//   		SaslScram512Auth: jsii.String("saslScram512Auth"),
//   	},
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	ServerRootCaCertificate: jsii.String("serverRootCaCertificate"),
//   	StartingPosition: jsii.String("startingPosition"),
//   	Vpc: &SelfManagedKafkaAccessConfigurationVpcProperty{
//   		SecurityGroup: []*string{
//   			jsii.String("securityGroup"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html
//
type CfnPipe_PipeSourceSelfManagedKafkaParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-topicname
	//
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// An array of server URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-additionalbootstrapservers
	//
	AdditionalBootstrapServers *[]*string `field:"optional" json:"additionalBootstrapServers" yaml:"additionalBootstrapServers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-consumergroupid
	//
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-credentials
	//
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-maximumbatchingwindowinseconds
	//
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// Optional SecretManager ARN which stores the database credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-serverrootcacertificate
	//
	ServerRootCaCertificate *string `field:"optional" json:"serverRootCaCertificate" yaml:"serverRootCaCertificate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-startingposition
	//
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html#cfn-pipes-pipe-pipesourceselfmanagedkafkaparameters-vpc
	//
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

