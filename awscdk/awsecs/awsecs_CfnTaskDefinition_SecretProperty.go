package awsecs


// The `Secret` property specifies an object representing the secret to expose to your container.
//
// For more information, see [Specifying Sensitive Data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretProperty := &secretProperty{
//   	name: jsii.String("name"),
//   	valueFrom: jsii.String("valueFrom"),
//   }
//
type CfnTaskDefinition_SecretProperty struct {
	// The name of the secret.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The secret to expose to the container.
	//
	// The supported values are either the full ARN of the AWS Secrets Manager secret or the full ARN of the parameter in the SSM Parameter Store.
	//
	// For information about the require AWS Identity and Access Management permissions, see [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-secrets.html#secrets-iam) (for Secrets Manager) or [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-parameters.html) (for Systems Manager Parameter store) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If the SSM Parameter Store parameter exists in the same Region as the task you're launching, then you can use either the full ARN or name of the parameter. If the parameter exists in a different Region, then the full ARN must be specified.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

