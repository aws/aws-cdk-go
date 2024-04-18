package awsdeadline


// Properties for defining a `CfnFarm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFarmProps := &CfnFarmProps{
//   	DisplayName: jsii.String("displayName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html
//
type CfnFarmProps struct {
	// The display name of the farm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html#cfn-deadline-farm-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// A description of the farm that helps identify what the farm is used for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html#cfn-deadline-farm-description
	//
	// Default: - "".
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN for the KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html#cfn-deadline-farm-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

