package awslogs


// This structure contains the information about one processor in a log transformer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorProperty := &ProcessorProperty{
//   	AddKeys: &AddKeysProperty{
//   		Entries: []interface{}{
//   			&AddKeyEntryProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//
//   				// the properties below are optional
//   				OverwriteIfExists: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	CopyValue: &CopyValueProperty{
//   		Entries: []interface{}{
//   			&CopyValueEntryProperty{
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//
//   				// the properties below are optional
//   				OverwriteIfExists: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Csv: &CsvProperty{
//   		Columns: []*string{
//   			jsii.String("columns"),
//   		},
//   		Delimiter: jsii.String("delimiter"),
//   		QuoteCharacter: jsii.String("quoteCharacter"),
//   		Source: jsii.String("source"),
//   	},
//   	DateTimeConverter: &DateTimeConverterProperty{
//   		MatchPatterns: []*string{
//   			jsii.String("matchPatterns"),
//   		},
//   		Source: jsii.String("source"),
//   		Target: jsii.String("target"),
//
//   		// the properties below are optional
//   		Locale: jsii.String("locale"),
//   		SourceTimezone: jsii.String("sourceTimezone"),
//   		TargetFormat: jsii.String("targetFormat"),
//   		TargetTimezone: jsii.String("targetTimezone"),
//   	},
//   	DeleteKeys: &DeleteKeysProperty{
//   		WithKeys: []*string{
//   			jsii.String("withKeys"),
//   		},
//   	},
//   	Grok: &GrokProperty{
//   		Match: jsii.String("match"),
//
//   		// the properties below are optional
//   		Source: jsii.String("source"),
//   	},
//   	ListToMap: &ListToMapProperty{
//   		Key: jsii.String("key"),
//   		Source: jsii.String("source"),
//
//   		// the properties below are optional
//   		Flatten: jsii.Boolean(false),
//   		FlattenedElement: jsii.String("flattenedElement"),
//   		Target: jsii.String("target"),
//   		ValueKey: jsii.String("valueKey"),
//   	},
//   	LowerCaseString: &LowerCaseStringProperty{
//   		WithKeys: []*string{
//   			jsii.String("withKeys"),
//   		},
//   	},
//   	MoveKeys: &MoveKeysProperty{
//   		Entries: []interface{}{
//   			&MoveKeyEntryProperty{
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//
//   				// the properties below are optional
//   				OverwriteIfExists: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	ParseCloudfront: &ParseCloudfrontProperty{
//   		Source: jsii.String("source"),
//   	},
//   	ParseJson: &ParseJSONProperty{
//   		Destination: jsii.String("destination"),
//   		Source: jsii.String("source"),
//   	},
//   	ParseKeyValue: &ParseKeyValueProperty{
//   		Destination: jsii.String("destination"),
//   		FieldDelimiter: jsii.String("fieldDelimiter"),
//   		KeyPrefix: jsii.String("keyPrefix"),
//   		KeyValueDelimiter: jsii.String("keyValueDelimiter"),
//   		NonMatchValue: jsii.String("nonMatchValue"),
//   		OverwriteIfExists: jsii.Boolean(false),
//   		Source: jsii.String("source"),
//   	},
//   	ParsePostgres: &ParsePostgresProperty{
//   		Source: jsii.String("source"),
//   	},
//   	ParseRoute53: &ParseRoute53Property{
//   		Source: jsii.String("source"),
//   	},
//   	ParseVpc: &ParseVPCProperty{
//   		Source: jsii.String("source"),
//   	},
//   	ParseWaf: &ParseWAFProperty{
//   		Source: jsii.String("source"),
//   	},
//   	RenameKeys: &RenameKeysProperty{
//   		Entries: []interface{}{
//   			&RenameKeyEntryProperty{
//   				Key: jsii.String("key"),
//   				RenameTo: jsii.String("renameTo"),
//
//   				// the properties below are optional
//   				OverwriteIfExists: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	SplitString: &SplitStringProperty{
//   		Entries: []interface{}{
//   			&SplitStringEntryProperty{
//   				Delimiter: jsii.String("delimiter"),
//   				Source: jsii.String("source"),
//   			},
//   		},
//   	},
//   	SubstituteString: &SubstituteStringProperty{
//   		Entries: []interface{}{
//   			&SubstituteStringEntryProperty{
//   				From: jsii.String("from"),
//   				Source: jsii.String("source"),
//   				To: jsii.String("to"),
//   			},
//   		},
//   	},
//   	TrimString: &TrimStringProperty{
//   		WithKeys: []*string{
//   			jsii.String("withKeys"),
//   		},
//   	},
//   	TypeConverter: &TypeConverterProperty{
//   		Entries: []interface{}{
//   			&TypeConverterEntryProperty{
//   				Key: jsii.String("key"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	UpperCaseString: &UpperCaseStringProperty{
//   		WithKeys: []*string{
//   			jsii.String("withKeys"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html
//
type CfnTransformer_ProcessorProperty struct {
	// Use this parameter to include the [addKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-addKeys) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-addkeys
	//
	AddKeys interface{} `field:"optional" json:"addKeys" yaml:"addKeys"`
	// Use this parameter to include the [copyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-copyValue) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-copyvalue
	//
	CopyValue interface{} `field:"optional" json:"copyValue" yaml:"copyValue"`
	// Use this parameter to include the [CSV](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-CSV) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-csv
	//
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
	// Use this parameter to include the [datetimeConverter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-datetimeConverter) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-datetimeconverter
	//
	DateTimeConverter interface{} `field:"optional" json:"dateTimeConverter" yaml:"dateTimeConverter"`
	// Use this parameter to include the [deleteKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-deleteKeys) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-deletekeys
	//
	DeleteKeys interface{} `field:"optional" json:"deleteKeys" yaml:"deleteKeys"`
	// Use this parameter to include the [grok](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-grok) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-grok
	//
	Grok interface{} `field:"optional" json:"grok" yaml:"grok"`
	// Use this parameter to include the [listToMap](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-listToMap) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-listtomap
	//
	ListToMap interface{} `field:"optional" json:"listToMap" yaml:"listToMap"`
	// Use this parameter to include the [lowerCaseString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-lowerCaseString) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-lowercasestring
	//
	LowerCaseString interface{} `field:"optional" json:"lowerCaseString" yaml:"lowerCaseString"`
	// Use this parameter to include the [moveKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-moveKeys) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-movekeys
	//
	MoveKeys interface{} `field:"optional" json:"moveKeys" yaml:"moveKeys"`
	// Use this parameter to include the [parseCloudfront](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseCloudfront) processor in your transformer.
	//
	// If you use this processor, it must be the first processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-parsecloudfront
	//
	ParseCloudfront interface{} `field:"optional" json:"parseCloudfront" yaml:"parseCloudfront"`
	// Use this parameter to include the [parseJSON](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseJSON) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-parsejson
	//
	ParseJson interface{} `field:"optional" json:"parseJson" yaml:"parseJson"`
	// Use this parameter to include the [parseKeyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseKeyValue) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-parsekeyvalue
	//
	ParseKeyValue interface{} `field:"optional" json:"parseKeyValue" yaml:"parseKeyValue"`
	// Use this parameter to include the [parsePostGres](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-parsePostGres) processor in your transformer.
	//
	// If you use this processor, it must be the first processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-parsepostgres
	//
	ParsePostgres interface{} `field:"optional" json:"parsePostgres" yaml:"parsePostgres"`
	// Use this parameter to include the [parseRoute53](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseRoute53) processor in your transformer.
	//
	// If you use this processor, it must be the first processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-parseroute53
	//
	ParseRoute53 interface{} `field:"optional" json:"parseRoute53" yaml:"parseRoute53"`
	// Use this parameter to include the [parseVPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseVPC) processor in your transformer.
	//
	// If you use this processor, it must be the first processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-parsevpc
	//
	ParseVpc interface{} `field:"optional" json:"parseVpc" yaml:"parseVpc"`
	// Use this parameter to include the [parseWAF](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-parseWAF) processor in your transformer.
	//
	// If you use this processor, it must be the first processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-parsewaf
	//
	ParseWaf interface{} `field:"optional" json:"parseWaf" yaml:"parseWaf"`
	// Use this parameter to include the [renameKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-renameKeys) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-renamekeys
	//
	RenameKeys interface{} `field:"optional" json:"renameKeys" yaml:"renameKeys"`
	// Use this parameter to include the [splitString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-splitString) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-splitstring
	//
	SplitString interface{} `field:"optional" json:"splitString" yaml:"splitString"`
	// Use this parameter to include the [substituteString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-substituteString) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-substitutestring
	//
	SubstituteString interface{} `field:"optional" json:"substituteString" yaml:"substituteString"`
	// Use this parameter to include the [trimString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-trimString) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-trimstring
	//
	TrimString interface{} `field:"optional" json:"trimString" yaml:"trimString"`
	// Use this parameter to include the [typeConverter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-typeConverter) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-typeconverter
	//
	TypeConverter interface{} `field:"optional" json:"typeConverter" yaml:"typeConverter"`
	// Use this parameter to include the [upperCaseString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-upperCaseString) processor in your transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-processor.html#cfn-logs-transformer-processor-uppercasestring
	//
	UpperCaseString interface{} `field:"optional" json:"upperCaseString" yaml:"upperCaseString"`
}

