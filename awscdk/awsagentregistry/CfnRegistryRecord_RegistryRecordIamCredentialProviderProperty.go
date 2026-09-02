package awsagentregistry


// IAM credential provider configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryRecordIamCredentialProviderProperty := &RegistryRecordIamCredentialProviderProperty{
//   	Region: jsii.String("region"),
//   	RoleArn: jsii.String("roleArn"),
//   	Service: jsii.String("service"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordiamcredentialprovider.html
//
type CfnRegistryRecord_RegistryRecordIamCredentialProviderProperty struct {
	// The SigV4 signing region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordiamcredentialprovider.html#cfn-agentregistry-registryrecord-registryrecordiamcredentialprovider-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The ARN of the IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordiamcredentialprovider.html#cfn-agentregistry-registryrecord-registryrecordiamcredentialprovider-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The SigV4 signing service name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-registryrecordiamcredentialprovider.html#cfn-agentregistry-registryrecord-registryrecordiamcredentialprovider-service
	//
	Service *string `field:"optional" json:"service" yaml:"service"`
}

