package awsbatch


// An object representing the secret to expose to your container.
//
// Secrets can be exposed to a container in the following ways:
//
// - To inject sensitive data into your containers as environment variables, use the `secrets` container definition parameter.
// - To reference sensitive information in the log configuration of a container, use the `secretOptions` container definition parameter.
//
// For more information, see [Specifying sensitive data](https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data.html) in the *AWS Batch User Guide* .
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
type CfnJobDefinition_SecretProperty struct {
	// The name of the secret.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The secret to expose to the container.
	//
	// The supported values are either the full ARN of the AWS Secrets Manager secret or the full ARN of the parameter in the AWS Systems Manager Parameter Store.
	//
	// > If the AWS Systems Manager Parameter Store parameter exists in the same Region as the job you're launching, then you can use either the full ARN or name of the parameter. If the parameter exists in a different Region, then the full ARN must be specified.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

