package previewawsiotanalyticsevents


// Type definition for RequestParametersItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParametersItem := &RequestParametersItem{
//   	Address: []*string{
//   		jsii.String("address"),
//   	},
//   	BigEndian: []*string{
//   		jsii.String("bigEndian"),
//   	},
//   	Capacity: []*string{
//   		jsii.String("capacity"),
//   	},
//   	Hb: []*f64{
//   		jsii.Number(123),
//   	},
//   	IsReadOnly: []*string{
//   		jsii.String("isReadOnly"),
//   	},
//   	Limit: []*string{
//   		jsii.String("limit"),
//   	},
//   	Mark: []*string{
//   		jsii.String("mark"),
//   	},
//   	NativeByteOrder: []*string{
//   		jsii.String("nativeByteOrder"),
//   	},
//   	Offset: []*string{
//   		jsii.String("offset"),
//   	},
//   	Position: []*string{
//   		jsii.String("position"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParametersItem struct {
	// address property.
	//
	// Specify an array of string values to match this event if the actual value of address is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Address *[]*string `field:"optional" json:"address" yaml:"address"`
	// bigEndian property.
	//
	// Specify an array of string values to match this event if the actual value of bigEndian is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BigEndian *[]*string `field:"optional" json:"bigEndian" yaml:"bigEndian"`
	// capacity property.
	//
	// Specify an array of string values to match this event if the actual value of capacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Capacity *[]*string `field:"optional" json:"capacity" yaml:"capacity"`
	// hb property.
	//
	// Specify an array of string values to match this event if the actual value of hb is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Hb *[]*float64 `field:"optional" json:"hb" yaml:"hb"`
	// isReadOnly property.
	//
	// Specify an array of string values to match this event if the actual value of isReadOnly is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IsReadOnly *[]*string `field:"optional" json:"isReadOnly" yaml:"isReadOnly"`
	// limit property.
	//
	// Specify an array of string values to match this event if the actual value of limit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Limit *[]*string `field:"optional" json:"limit" yaml:"limit"`
	// mark property.
	//
	// Specify an array of string values to match this event if the actual value of mark is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Mark *[]*string `field:"optional" json:"mark" yaml:"mark"`
	// nativeByteOrder property.
	//
	// Specify an array of string values to match this event if the actual value of nativeByteOrder is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NativeByteOrder *[]*string `field:"optional" json:"nativeByteOrder" yaml:"nativeByteOrder"`
	// offset property.
	//
	// Specify an array of string values to match this event if the actual value of offset is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Offset *[]*string `field:"optional" json:"offset" yaml:"offset"`
	// position property.
	//
	// Specify an array of string values to match this event if the actual value of position is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Position *[]*string `field:"optional" json:"position" yaml:"position"`
}

