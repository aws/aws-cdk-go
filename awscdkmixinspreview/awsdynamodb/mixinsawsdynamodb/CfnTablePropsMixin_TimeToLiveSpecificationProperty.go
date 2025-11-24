package mixinsawsdynamodb


// Represents the settings used to enable or disable Time to Live (TTL) for the specified table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeToLiveSpecificationProperty := &TimeToLiveSpecificationProperty{
//   	AttributeName: jsii.String("attributeName"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-timetolivespecification.html
//
type CfnTablePropsMixin_TimeToLiveSpecificationProperty struct {
	// The name of the TTL attribute used to store the expiration time for items in the table.
	//
	// > - The `AttributeName` property is required when enabling the TTL, or when TTL is already enabled.
	// > - To update this property, you must first disable TTL and then enable TTL with the new attribute name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-timetolivespecification.html#cfn-dynamodb-table-timetolivespecification-attributename
	//
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// Indicates whether TTL is to be enabled (true) or disabled (false) on the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-timetolivespecification.html#cfn-dynamodb-table-timetolivespecification-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

