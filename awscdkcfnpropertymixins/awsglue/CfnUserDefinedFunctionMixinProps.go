package awsglue


// Properties for CfnUserDefinedFunctionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnUserDefinedFunctionMixinProps := &CfnUserDefinedFunctionMixinProps{
//   	ClassName: jsii.String("className"),
//   	DatabaseName: jsii.String("databaseName"),
//   	FunctionName: jsii.String("functionName"),
//   	FunctionType: jsii.String("functionType"),
//   	OwnerName: jsii.String("ownerName"),
//   	OwnerType: jsii.String("ownerType"),
//   	ResourceUris: []interface{}{
//   		&ResourceUriProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-userdefinedfunction.html
//
type CfnUserDefinedFunctionMixinProps struct {
	// The Java class that contains the function code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-userdefinedfunction.html#cfn-glue-userdefinedfunction-classname
	//
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// The name of the catalog database in which the function is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-userdefinedfunction.html#cfn-glue-userdefinedfunction-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-userdefinedfunction.html#cfn-glue-userdefinedfunction-functionname
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The type of the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-userdefinedfunction.html#cfn-glue-userdefinedfunction-functiontype
	//
	FunctionType *string `field:"optional" json:"functionType" yaml:"functionType"`
	// The owner of the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-userdefinedfunction.html#cfn-glue-userdefinedfunction-ownername
	//
	OwnerName *string `field:"optional" json:"ownerName" yaml:"ownerName"`
	// The owner type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-userdefinedfunction.html#cfn-glue-userdefinedfunction-ownertype
	//
	OwnerType *string `field:"optional" json:"ownerType" yaml:"ownerType"`
	// The resource URIs for the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-userdefinedfunction.html#cfn-glue-userdefinedfunction-resourceuris
	//
	ResourceUris interface{} `field:"optional" json:"resourceUris" yaml:"resourceUris"`
}

