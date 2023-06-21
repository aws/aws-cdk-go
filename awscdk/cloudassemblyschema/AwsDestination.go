package cloudassemblyschema


// Destination for assets that need to be uploaded to AWS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsDestination := &AwsDestination{
//   	AssumeRoleArn: jsii.String("assumeRoleArn"),
//   	AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	Region: jsii.String("region"),
//   }
//
type AwsDestination struct {
	// The role that needs to be assumed while publishing this asset.
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	AssumeRoleExternalId *string `field:"optional" json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

