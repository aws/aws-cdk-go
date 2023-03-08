package awssimspaceweaver


// Properties for defining a `CfnSimulation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSimulationProps := &cfnSimulationProps{
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	schemaS3Location: &s3LocationProperty{
//   		bucketName: jsii.String("bucketName"),
//   		objectKey: jsii.String("objectKey"),
//   	},
//   }
//
type CfnSimulationProps struct {
	// The name of the simulation.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management ( IAM ) role that the simulation assumes to perform actions.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* . For more information about IAM roles, see [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the *AWS Identity and Access Management User Guide* .
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The location of the simulation schema in Amazon Simple Storage Service ( Amazon S3 ).
	//
	// For more information about Amazon S3 , see the [*Amazon Simple Storage Service User Guide*](https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html) .
	SchemaS3Location interface{} `field:"optional" json:"schemaS3Location" yaml:"schemaS3Location"`
}

