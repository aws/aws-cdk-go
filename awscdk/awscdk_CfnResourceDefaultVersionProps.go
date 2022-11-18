// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Properties for defining a `CfnResourceDefaultVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceDefaultVersionProps := &cfnResourceDefaultVersionProps{
//   	typeName: jsii.String("typeName"),
//   	typeVersionArn: jsii.String("typeVersionArn"),
//   	versionId: jsii.String("versionId"),
//   }
//
type CfnResourceDefaultVersionProps struct {
	// The name of the resource.
	//
	// Conditional: You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// The Amazon Resource Name (ARN) of the resource version.
	//
	// Conditional: You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	TypeVersionArn *string `field:"optional" json:"typeVersionArn" yaml:"typeVersionArn"`
	// The ID of a specific version of the resource.
	//
	// The version ID is the value at the end of the Amazon Resource Name (ARN) assigned to the resource version when it's registered.
	//
	// Conditional: You must specify either `TypeVersionArn` , or `TypeName` and `VersionId` .
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

