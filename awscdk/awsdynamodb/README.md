# Amazon DynamoDB Construct Library

> The DynamoDB construct library has two table constructs - `Table` and `TableV2`. `TableV2` is the preferred construct for all use cases, including creating a single table or a table with multiple `replicas`.

[`Table` API documentation](./TABLE_V1_API.md)

Here is a minimal deployable DynamoDB table using `TableV2`:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
})
```

By default, `TableV2` will create a single table in the main deployment region referred to as the primary table. The properties of the primary table are configurable via `TableV2` properties. For example, consider the following DynamoDB table created using the `TableV2` construct defined in a `Stack` being deployed to `us-west-2`:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	ContributorInsightsSpecification: &ContributorInsightsSpecification{
		Enabled: jsii.Boolean(true),
	},
	TableClass: dynamodb.TableClass_STANDARD_INFREQUENT_ACCESS,
	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecification{
		PointInTimeRecoveryEnabled: jsii.Boolean(true),
	},
})
```

The above `TableV2` definition will result in the provisioning of a single table in `us-west-2` with properties that match the properties set on the `TableV2` instance.

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html

## Replicas

The `TableV2` construct can be configured with replica tables. This will enable you to work with your table as a global table. To do this, the `TableV2` construct must be defined in a `Stack` with a defined region. The main deployment region must not be given as a replica because this is created by default with the `TableV2` construct. The following is a minimal example of defining `TableV2` with `replicas`. This `TableV2` definition will provision three copies of the table - one in `us-west-2` (primary deployment region), one in `us-east-1`, and one in `us-east-2`.

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
		},
	},
})
```

Alternatively, you can add new `replicas` to an instance of the `TableV2` construct using the `addReplica` method:

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
	},
})

globalTable.AddReplica(&ReplicaTableProps{
	Region: jsii.String("us-east-2"),
	DeletionProtection: jsii.Boolean(true),
})
```

The following properties are configurable on a per-replica basis, but will be inherited from the `TableV2` properties if not specified:

* contributorInsightsSpecification
* deletionProtection
* pointInTimeRecoverySpecification
* tableClass
* readCapacity (only configurable if the `TableV2` billing mode is `PROVISIONED`)
* globalSecondaryIndexes (only `contributorInsightsSpecification` and `readCapacity`)

The following example shows how to define properties on a per-replica basis:

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	ContributorInsightsSpecification: &ContributorInsightsSpecification{
		Enabled: jsii.Boolean(true),
	},
	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecification{
		PointInTimeRecoveryEnabled: jsii.Boolean(true),
	},
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
			TableClass: dynamodb.TableClass_STANDARD_INFREQUENT_ACCESS,
			PointInTimeRecoverySpecification: &PointInTimeRecoverySpecification{
				PointInTimeRecoveryEnabled: jsii.Boolean(false),
			},
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
			ContributorInsightsSpecification: &ContributorInsightsSpecification{
				Enabled: jsii.Boolean(false),
			},
		},
	},
})
```

To obtain an `ITableV2` reference to a specific replica table, call the `replica` method on an instance of the `TableV2` construct and pass the replica region as an argument:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var user User


type fooStack struct {
	Stack
	globalTable TableV2
}

func newFooStack(scope Construct, id *string, props StackProps) *fooStack {
	this := &fooStack{}
	cdk.NewStack_Override(this, scope, id, props)

	this.globalTable = dynamodb.NewTableV2(this, jsii.String("GlobalTable"), &TablePropsV2{
		PartitionKey: &Attribute{
			Name: jsii.String("pk"),
			Type: dynamodb.AttributeType_STRING,
		},
		Replicas: []ReplicaTableProps{
			&ReplicaTableProps{
				Region: jsii.String("us-east-1"),
			},
			&ReplicaTableProps{
				Region: jsii.String("us-east-2"),
			},
		},
	})
	return this
}

type barStackProps struct {
	StackProps
	replicaTable ITableV2
}

type barStack struct {
	Stack
}

func newBarStack(scope Construct, id *string, props barStackProps) *barStack {
	this := &barStack{}
	cdk.NewStack_Override(this, scope, id, props)

	// user is given grantWriteData permissions to replica in us-east-1
	*props.replicaTable.GrantWriteData(user)
	return this
}

app := cdk.NewApp()

fooStack := NewFooStack(app, jsii.String("FooStack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})
barStack := NewBarStack(app, jsii.String("BarStack"), &barStackProps{
	replicaTable: fooStack.globalTable.Replica(jsii.String("us-east-1")),
	env: &Environment{
		Region: jsii.String("us-east-1"),
	},
})
```

Note: You can create an instance of the `TableV2` construct with as many `replicas` as needed as long as there is only one replica per region. After table creation you can add or remove `replicas`, but you can only add or remove a single replica in each update.

## Multi-Region Strong Consistency (MRSC)

By default, DynamoDB global tables provide eventual consistency across regions. For applications requiring strong consistency across regions, you can configure Multi-Region Strong Consistency (MRSC) using the `multiRegionConsistency` property.

MRSC global tables can be configured in two ways:

* **Three replicas**: Deploy your table across three regions within the same region set
* **Two replicas + one witness**: Deploy your table across two regions with a witness region for consensus

### Region Sets

MRSC global tables must be deployed within the same region set. The supported region sets are:

* **US Region set**: `us-east-1`, `us-east-2`, `us-west-2`
* **EU Region set**: `eu-west-1`, `eu-west-2`, `eu-west-3`, `eu-central-1`
* **AP Region set**: `ap-northeast-1`, `ap-northeast-2`, `ap-northeast-3`

### Three Replicas Configuration

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

mrscTable := dynamodb.NewTableV2(stack, jsii.String("MRSCTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	MultiRegionConsistency: dynamodb.MultiRegionConsistency_STRONG,
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
		},
	},
})
```

### Two Replicas + Witness Configuration

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

mrscTable := dynamodb.NewTableV2(stack, jsii.String("MRSCTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	MultiRegionConsistency: dynamodb.MultiRegionConsistency_STRONG,
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
	},
	WitnessRegion: jsii.String("us-east-2"),
})
```

### Important Considerations

* **Witness regions** can only be used with `MultiRegionConsistency.STRONG`. Attempting to specify a witness region with eventual consistency will result in a validation error.
* **Region validation**: All regions (primary, replicas, and witness) must be within the same region set.
* **Replica count**: When using a witness region, you must have exactly 2 replicas (including the primary). Without a witness region, you must have exactly 3 replicas.
* **Performance**: MRSC provides strong consistency but may have higher latency compared to eventual consistency.

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_HowItWorks.html#V2globaltables_HowItWorks.consistency-modes-mrsc

## Billing

The `TableV2` construct can be configured with on-demand or provisioned billing:

* On-demand - The default option. This is a flexible billing option capable of serving requests without capacity planning. The billing mode will be `PAY_PER_REQUEST`.
* You can optionally specify the `maxReadRequestUnits` or `maxWriteRequestUnits` on individual tables and associated global secondary indexes (GSIs). When you configure maximum throughput for an on-demand table, throughput requests that exceed the maximum amount specified will be throttled.
* Provisioned - Specify the `readCapacity` and `writeCapacity` that you need for your application. The billing mode will be `PROVISIONED`. Capacity can be configured using one of the following modes:

  * Fixed - provisioned throughput capacity is configured with a fixed number of I/O operations per second.
  * Autoscaled - provisioned throughput capacity is dynamically adjusted on your behalf in response to actual traffic patterns.

Note: `writeCapacity` can only be configured using autoscaled capacity.

The following example shows how to configure `TableV2` with on-demand billing:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Billing: dynamodb.Billing_OnDemand(),
})
```

The following example shows how to configure `TableV2` with on-demand billing with optional maximum throughput configured:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Billing: dynamodb.Billing_OnDemand(&MaxThroughputProps{
		MaxReadRequestUnits: jsii.Number(100),
		MaxWriteRequestUnits: jsii.Number(115),
	}),
})
```

When using provisioned billing, you must also specify `readCapacity` and `writeCapacity`. You can choose to configure `readCapacity` with fixed capacity or autoscaled capacity, but `writeCapacity` can only be configured with autoscaled capacity. The following example shows how to configure `TableV2` with provisioned billing:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Billing: dynamodb.Billing_Provisioned(&ThroughputProps{
		ReadCapacity: dynamodb.Capacity_Fixed(jsii.Number(10)),
		WriteCapacity: dynamodb.Capacity_Autoscaled(&AutoscaledCapacityOptions{
			MaxCapacity: jsii.Number(15),
		}),
	}),
})
```

When using provisioned billing, you can configure the `readCapacity` on a per-replica basis:

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Billing: dynamodb.Billing_Provisioned(&ThroughputProps{
		ReadCapacity: dynamodb.Capacity_Fixed(jsii.Number(10)),
		WriteCapacity: dynamodb.Capacity_Autoscaled(&AutoscaledCapacityOptions{
			MaxCapacity: jsii.Number(15),
		}),
	}),
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
			ReadCapacity: dynamodb.Capacity_*Autoscaled(&AutoscaledCapacityOptions{
				MaxCapacity: jsii.Number(20),
				TargetUtilizationPercent: jsii.Number(50),
			}),
		},
	},
})
```

When changing the billing for a table from provisioned to on-demand or from on-demand to provisioned, `seedCapacity` must be configured for each autoscaled resource:

```go
globalTable := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Billing: dynamodb.Billing_Provisioned(&ThroughputProps{
		ReadCapacity: dynamodb.Capacity_Fixed(jsii.Number(10)),
		WriteCapacity: dynamodb.Capacity_Autoscaled(&AutoscaledCapacityOptions{
			MaxCapacity: jsii.Number(10),
			SeedCapacity: jsii.Number(20),
		}),
	}),
})
```

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html

## Warm Throughput

Warm throughput refers to the number of read and write operations your DynamoDB table can instantaneously support.

This optional configuration allows you to pre-warm your table or index to handle anticipated throughput, ensuring optimal performance under expected load.

The Warm Throughput configuration settings are automatically replicated across all Global Table replicas.

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	WarmThroughput: &WarmThroughput{
		ReadUnitsPerSecond: jsii.Number(15000),
		WriteUnitsPerSecond: jsii.Number(20000),
	},
})
```

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/warm-throughput.html

## Encryption

All user data stored in a DynamoDB table is fully encrypted at rest. When creating an instance of the `TableV2` construct, you can select the following table encryption options:

* AWS owned keys - Default encryption type. The keys are owned by DynamoDB (no additional charge).
* AWS managed keys - The keys are stored in your account and are managed by AWS KMS (AWS KMS charges apply).
* Customer managed keys - The keys are stored in your account and are created, owned, and managed by you. You have full control over the KMS keys (AWS KMS charges apply).

The following is an example of how to configure `TableV2` with encryption using an AWS owned key:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Encryption: dynamodb.TableEncryptionV2_DynamoOwnedKey(),
})
```

The following is an example of how to configure `TableV2` with encryption using an AWS managed key:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Encryption: dynamodb.TableEncryptionV2_AwsManagedKey(),
})
```

When configuring `TableV2` with encryption using customer managed keys, you must specify the KMS key for the primary table as the `tableKey`. A map of `replicaKeyArns` must be provided containing each replica region and the associated KMS key ARN:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import kms "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

tableKey := kms.NewKey(stack, jsii.String("Key"))
replicaKeyArns := map[string]*string{
	"us-east-1": jsii.String("arn:aws:kms:us-east-1:123456789012:key/g24efbna-az9b-42ro-m3bp-cq249l94fca6"),
	"us-east-2": jsii.String("arn:aws:kms:us-east-2:123456789012:key/h90bkasj-bs1j-92wp-s2ka-bh857d60bkj8"),
}

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Encryption: dynamodb.TableEncryptionV2_CustomerManagedKey(tableKey, replicaKeyArns),
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
		},
	},
})
```

Note: When encryption is configured with customer managed keys, you must have a key already created in each replica region.

Further reading:
https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-mgmt

## Secondary Indexes

Secondary indexes allow efficient access to data with attributes other than the `primaryKey`. DynamoDB supports two types of secondary indexes:

* Global secondary index - An index with a `partitionKey` and a `sortKey` that can be different from those on the base table. A `globalSecondaryIndex` is considered "global" because queries on the index can span all of the data in the base table, across all partitions. A `globalSecondaryIndex` is stored in its own partition space away from the base table and scales separately from the base table.
* Local secondary index - An index that has the same `partitionKey` as the base table, but a different `sortKey`. A `localSecondaryIndex` is "local" in the sense that every partition of a `localSecondaryIndex` is scoped to a base table partition that has the same `partitionKey` value.

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/SecondaryIndexes.html

### Global Secondary Indexes

`TableV2` can be configured with `globalSecondaryIndexes` by providing them as a `TableV2` property:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	GlobalSecondaryIndexes: []GlobalSecondaryIndexPropsV2{
		&GlobalSecondaryIndexPropsV2{
			IndexName: jsii.String("gsi"),
			PartitionKey: &Attribute{
				Name: jsii.String("pk"),
				Type: dynamodb.AttributeType_STRING,
			},
		},
	},
})
```

Alternatively, you can add a `globalSecondaryIndex` using the `addGlobalSecondaryIndex` method:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	GlobalSecondaryIndexes: []GlobalSecondaryIndexPropsV2{
		&GlobalSecondaryIndexPropsV2{
			IndexName: jsii.String("gsi1"),
			PartitionKey: &Attribute{
				Name: jsii.String("pk"),
				Type: dynamodb.AttributeType_STRING,
			},
		},
	},
})

table.AddGlobalSecondaryIndex(&GlobalSecondaryIndexPropsV2{
	IndexName: jsii.String("gsi2"),
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
})
```

You can configure `readCapacity` and `writeCapacity` on a `globalSecondaryIndex` when an `TableV2` is configured with provisioned `billing`. If `TableV2` is configured with provisioned `billing` but `readCapacity` or `writeCapacity` are not configured on a `globalSecondaryIndex`, then they will be inherited from the capacity settings specified with the `billing` configuration:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Billing: dynamodb.Billing_Provisioned(&ThroughputProps{
		ReadCapacity: dynamodb.Capacity_Fixed(jsii.Number(10)),
		WriteCapacity: dynamodb.Capacity_Autoscaled(&AutoscaledCapacityOptions{
			MaxCapacity: jsii.Number(10),
		}),
	}),
	GlobalSecondaryIndexes: []GlobalSecondaryIndexPropsV2{
		&GlobalSecondaryIndexPropsV2{
			IndexName: jsii.String("gsi1"),
			PartitionKey: &Attribute{
				Name: jsii.String("pk"),
				Type: dynamodb.AttributeType_STRING,
			},
			ReadCapacity: dynamodb.Capacity_*Fixed(jsii.Number(15)),
		},
		&GlobalSecondaryIndexPropsV2{
			IndexName: jsii.String("gsi2"),
			PartitionKey: &Attribute{
				Name: jsii.String("pk"),
				Type: dynamodb.AttributeType_STRING,
			},
			WriteCapacity: dynamodb.Capacity_*Autoscaled(&AutoscaledCapacityOptions{
				MinCapacity: jsii.Number(5),
				MaxCapacity: jsii.Number(20),
			}),
		},
	},
})
```

All `globalSecondaryIndexes` for replica tables are inherited from the primary table. You can configure `contributorInsightsSpecification` and `readCapacity` for each `globalSecondaryIndex` on a per-replica basis:

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	ContributorInsightsSpecification: &ContributorInsightsSpecification{
		Enabled: jsii.Boolean(true),
	},
	Billing: dynamodb.Billing_Provisioned(&ThroughputProps{
		ReadCapacity: dynamodb.Capacity_Fixed(jsii.Number(10)),
		WriteCapacity: dynamodb.Capacity_Autoscaled(&AutoscaledCapacityOptions{
			MaxCapacity: jsii.Number(10),
		}),
	}),
	// each global secondary index will inherit contributor insights as true
	GlobalSecondaryIndexes: []GlobalSecondaryIndexPropsV2{
		&GlobalSecondaryIndexPropsV2{
			IndexName: jsii.String("gsi1"),
			PartitionKey: &Attribute{
				Name: jsii.String("pk"),
				Type: dynamodb.AttributeType_STRING,
			},
			ReadCapacity: dynamodb.Capacity_*Fixed(jsii.Number(15)),
		},
		&GlobalSecondaryIndexPropsV2{
			IndexName: jsii.String("gsi2"),
			PartitionKey: &Attribute{
				Name: jsii.String("pk"),
				Type: dynamodb.AttributeType_STRING,
			},
			WriteCapacity: dynamodb.Capacity_*Autoscaled(&AutoscaledCapacityOptions{
				MinCapacity: jsii.Number(5),
				MaxCapacity: jsii.Number(20),
			}),
		},
	},
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
			GlobalSecondaryIndexOptions: map[string]ReplicaGlobalSecondaryIndexOptions{
				"gsi1": &ReplicaGlobalSecondaryIndexOptions{
					"readCapacity": dynamodb.Capacity_*Autoscaled(&AutoscaledCapacityOptions{
						"minCapacity": jsii.Number(1),
						"maxCapacity": jsii.Number(10),
					}),
				},
			},
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
			GlobalSecondaryIndexOptions: map[string]ReplicaGlobalSecondaryIndexOptions{
				"gsi2": &ReplicaGlobalSecondaryIndexOptions{
					"contributorInsightsSpecification": &ContributorInsightsSpecification{
						"enabled": jsii.Boolean(false),
					},
				},
			},
		},
	},
})
```

### Local Secondary Indexes

`TableV2` can only be configured with `localSecondaryIndexes` when a `sortKey` is defined as a `TableV2` property.

You can provide `localSecondaryIndexes` as a `TableV2` property:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	SortKey: &Attribute{
		Name: jsii.String("sk"),
		Type: dynamodb.AttributeType_NUMBER,
	},
	LocalSecondaryIndexes: []LocalSecondaryIndexProps{
		&LocalSecondaryIndexProps{
			IndexName: jsii.String("lsi"),
			SortKey: &Attribute{
				Name: jsii.String("sk"),
				Type: dynamodb.AttributeType_NUMBER,
			},
		},
	},
})
```

Alternatively, you can add a `localSecondaryIndex` using the `addLocalSecondaryIndex` method:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	SortKey: &Attribute{
		Name: jsii.String("sk"),
		Type: dynamodb.AttributeType_NUMBER,
	},
	LocalSecondaryIndexes: []LocalSecondaryIndexProps{
		&LocalSecondaryIndexProps{
			IndexName: jsii.String("lsi1"),
			SortKey: &Attribute{
				Name: jsii.String("sk"),
				Type: dynamodb.AttributeType_NUMBER,
			},
		},
	},
})

table.AddLocalSecondaryIndex(&LocalSecondaryIndexProps{
	IndexName: jsii.String("lsi2"),
	SortKey: &Attribute{
		Name: jsii.String("sk"),
		Type: dynamodb.AttributeType_NUMBER,
	},
})
```

## Streams

Each DynamoDB table produces an independent stream based on all its writes, regardless of the origination point for those writes. DynamoDB supports two stream types:

* DynamoDB streams - Capture item-level changes in your table, and push the changes to a DynamoDB stream. You then can access the change information through the DynamoDB Streams API.
* Kinesis streams - Amazon Kinesis Data Streams for DynamoDB captures item-level changes in your table, and replicates the changes to a Kinesis data stream. You then can consume and manage the change information from Kinesis.

### DynamoDB Streams

A `dynamoStream` can be configured as a `TableV2` property. If the `TableV2` instance has replica tables, then all replica tables will inherit the `dynamoStream` setting from the primary table.  If replicas are configured, but `dynamoStream` is not configured, then the primary table and all replicas will be automatically configured with the `NEW_AND_OLD_IMAGES` stream view type.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import kinesis "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

globalTable := dynamodb.NewTableV2(this, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	DynamoStream: dynamodb.StreamViewType_OLD_IMAGE,
	// tables in us-west-2, us-east-1, and us-east-2 all have dynamo stream type of OLD_IMAGES
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
		},
	},
})
```

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Streams.html

### Kinesis Streams

A `kinesisStream` can be configured as a `TableV2` property. Replica tables will not inherit the `kinesisStream` configured for the primary table and should added on a per-replica basis.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

stream1 := kinesis.NewStream(stack, jsii.String("Stream1"))
stream2 := kinesis.Stream_FromStreamArn(stack, jsii.String("Stream2"), jsii.String("arn:aws:kinesis:us-east-2:123456789012:stream/my-stream"))

globalTable := dynamodb.NewTableV2(this, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	KinesisStream: stream1,
	 // for table in us-west-2
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
			KinesisStream: stream2,
		},
	},
})
```

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/kds.html

## Keys

When an instance of the `TableV2` construct is defined, you must define its schema using the `partitionKey` (required) and `sortKey` (optional) properties.

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	SortKey: &Attribute{
		Name: jsii.String("sk"),
		Type: dynamodb.AttributeType_NUMBER,
	},
})
```

## Contributor Insights

Enabling `contributorInsightSpecification` for `TableV2` will provide information about the most accessed and throttled or throttled only items in a table or `globalSecondaryIndex`. DynamoDB delivers this information to you via CloudWatch Contributor Insights rules, reports, and graphs of report data.

By default, Contributor Insights for DynamoDB monitors all requests, including both the most accessed and most throttled items.
To limit the scope to only the most accessed or only the most throttled items, use the optional `mode` parameter.

* To monitor all traffic on a table or index, set `mode` to `ContributorInsightsMode.ACCESSED_AND_THROTTLED_KEYS`.
* To monitor only throttled traffic on a table or index, set `mode` to `ContributorInsightsMode.THROTTLED_KEYS`.

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	ContributorInsightsSpecification: &ContributorInsightsSpecification{
		Enabled: jsii.Boolean(true),
		Mode: dynamodb.ContributorInsightsMode_ACCESSED_AND_THROTTLED_KEYS,
	},
})
```

When you use `Table`, you can enable contributor insights for a table or specific global secondary index by setting `contributorInsightsSpecification` parameter `enabled` to `true`.

```go
table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	ContributorInsightsSpecification: &ContributorInsightsSpecification{
		 // for a table
		Enabled: jsii.Boolean(true),
		Mode: dynamodb.ContributorInsightsMode_THROTTLED_KEYS,
	},
})

table.AddGlobalSecondaryIndex(&GlobalSecondaryIndexProps{
	ContributorInsightsSpecification: &ContributorInsightsSpecification{
		 // for a specific global secondary index
		Enabled: jsii.Boolean(true),
	},
	IndexName: jsii.String("gsi"),
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
})
```

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/contributorinsights_HowItWorks.html

## Deletion Protection

`deletionProtection` determines if your DynamoDB table is protected from deletion and is configurable as a `TableV2` property. When enabled, the table cannot be deleted by any user or process.

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	DeletionProtection: jsii.Boolean(true),
})
```

You can also specify the `removalPolicy` as a property of the `TableV2` construct. This property allows you to control what happens to tables provisioned using `TableV2` during `stack` deletion. By default, the `removalPolicy` is `RETAIN` which will cause all tables provisioned using `TableV2` to be retained in the account, but orphaned from the `stack` they were created in. You can also set the `removalPolicy` to `DESTROY` which will delete all tables created using `TableV2` during `stack` deletion:

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	// applies to all replicas, i.e., us-west-2, us-east-1, us-east-2
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
		},
	},
})
```

`deletionProtection` is configurable on a per-replica basis. If the `removalPolicy` is set to `DESTROY`, but some `replicas` have `deletionProtection` enabled, then only the `replicas` without `deletionProtection` will be deleted during `stack` deletion:

```go
import "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
	DeletionProtection: jsii.Boolean(true),
	// only the replica in us-east-1 will be deleted during stack deletion
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
			DeletionProtection: jsii.Boolean(false),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
			DeletionProtection: jsii.Boolean(true),
		},
	},
})
```

## Point-in-Time Recovery

`pointInTimeRecoverySpecifcation` provides automatic backups of your DynamoDB table data which helps protect your tables from accidental write or delete operations.

You can also choose to set `recoveryPeriodInDays` to a value between `1` and `35` which dictates how many days of recoverable data is stored. If no value is provided, the recovery period defaults to `35` days.

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecification{
		PointInTimeRecoveryEnabled: jsii.Boolean(true),
		RecoveryPeriodInDays: jsii.Number(4),
	},
})
```

## Table Class

You can configure a `TableV2` instance with table classes:

* STANDARD - the default mode, and is recommended for the vast majority of workloads.
* STANDARD_INFREQUENT_ACCESS - optimized for tables where storage is the dominant cost.

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	TableClass: dynamodb.TableClass_STANDARD_INFREQUENT_ACCESS,
})
```

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.TableClasses.html

## Tags

You can add tags to a `TableV2` in several ways. By adding the tags to the construct itself it will apply the tags to the
primary table.

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Tags: []CfnTag{
		&CfnTag{
			Key: jsii.String("primaryTableTagKey"),
			Value: jsii.String("primaryTableTagValue"),
		},
	},
})
```

You can also add tags to replica tables by specifying them within the replica table properties.

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-west-1"),
			Tags: []CfnTag{
				&CfnTag{
					Key: jsii.String("replicaTableTagKey"),
					Value: jsii.String("replicaTableTagValue"),
				},
			},
		},
	},
})
```

## Referencing Existing Global Tables

To reference an existing DynamoDB table in your CDK application, use the `TableV2.fromTableName`, `TableV2.fromTableArn`, or `TableV2.fromTableAttributes`
factory methods:

```go
var user User


table := dynamodb.TableV2_FromTableArn(this, jsii.String("ImportedTable"), jsii.String("arn:aws:dynamodb:us-east-1:123456789012:table/my-table"))
// now you can call methods on the referenced table
table.GrantReadWriteData(user)
```

If you intend to use the `tableStreamArn` (including indirectly, for example by creating an
`aws-cdk-lib/aws-lambda-event-sources.DynamoEventSource` on the referenced table), you *must* use the
`TableV2.fromTableAttributes` method and the `tableStreamArn` property *must* be populated.

To grant permissions to indexes for a referenced table you can either set `grantIndexPermissions` to `true`, or you can provide the indexes via the `globalIndexes` or `localIndexes` properties. This will enable `grant*` methods to also grant permissions to *all* table indexes.

## Resource Policy

Using `resourcePolicy` you can add a [resource policy](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/access-control-resource-based.html) to a table in the form of a `PolicyDocument`:

```
    // resource policy document
    const policy = new iam.PolicyDocument({
      statements: [
        new iam.PolicyStatement({
          actions: ['dynamodb:GetItem'],
          principals: [new iam.AccountRootPrincipal()],
          resources: ['*'],
        }),
      ],
    });

    // table with resource policy
    new dynamodb.TableV2(this, 'TableTestV2-1', {
      partitionKey: {
        name: 'id',
        type: dynamodb.AttributeType.STRING,
      },
      removalPolicy: RemovalPolicy.DESTROY,
      resourcePolicy: policy,
    });
```

### Adding Resource Policy Statements Dynamically

You can also add resource policy statements to a table after it's created using the `addToResourcePolicy` method. Following the same pattern as KMS, resource policies use wildcard resources to avoid circular dependencies:

```go
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
})

// Standard resource policy (recommended approach)
table.AddToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("dynamodb:GetItem"),
		jsii.String("dynamodb:PutItem"),
		jsii.String("dynamodb:Query"),
	},
	Principals: []IPrincipal{
		iam.NewAccountRootPrincipal(),
	},
	Resources: []*string{
		jsii.String("*"),
	},
}))

// Allow specific service access
table.AddToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("dynamodb:Query"),
	},
	Principals: []IPrincipal{
		iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
	},
	Resources: []*string{
		jsii.String("*"),
	},
}))
```

#### Scoped Resource Policies (Advanced)

For scoped resource policies that reference specific table ARNs, you must specify an explicit table name:

```go
import "github.com/aws/aws-cdk-go/awscdk"


// Table with explicit name enables scoped resource policies
table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
	TableName: jsii.String("my-explicit-table-name"),
	 // Required for scoped resources
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
})

// Now you can use scoped resources
table.AddToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("dynamodb:GetItem"),
	},
	Principals: []IPrincipal{
		iam.NewAccountRootPrincipal(),
	},
	Resources: []*string{
		awscdk.Fn_Sub(jsii.String("arn:aws:dynamodb:${AWS::Region}:${AWS::AccountId}:table/my-explicit-table-name")),
		awscdk.Fn_*Sub(jsii.String("arn:aws:dynamodb:${AWS::Region}:${AWS::AccountId}:table/my-explicit-table-name/index/*")),
	},
}))
```

**Important Limitations:**

* **Auto-generated table names**: Must use `resources: ['*']` to avoid circular dependencies
* **Explicit table names**: Enable scoped resources but lose CDK's automatic naming benefits
* **CloudFormation constraint**: Resource policies cannot reference the resource they're attached to during creation

TableV2 doesn’t support creating a replica and adding a resource-based policy to that replica in the same stack update in Regions other than the Region where you deploy the stack update.
To incorporate a resource-based policy into a replica, you'll need to initially deploy the replica without the policy, followed by a subsequent update to include the desired policy.

## Grants

Using any of the `grant*` methods on an instance of the `TableV2` construct will only apply to the primary table, its indexes, and any associated `encryptionKey`. As an example, `grantReadData` used below will only apply the table in `us-west-2`:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import kms "github.com/aws/aws-cdk-go/awscdk"

var user User


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

tableKey := kms.NewKey(stack, jsii.String("Key"))
replicaKeyArns := map[string]*string{
	"us-east-1": jsii.String("arn:aws:kms:us-east-1:123456789012:key/g24efbna-az9b-42ro-m3bp-cq249l94fca6"),
	"us-east-2": jsii.String("arn:aws:kms:us-east-2:123456789012:key/g24efbna-az9b-42ro-m3bp-cq249l94fca6"),
}

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Encryption: dynamodb.TableEncryptionV2_CustomerManagedKey(tableKey, replicaKeyArns),
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
		},
	},
})

// grantReadData only applies to the table in us-west-2 and the tableKey
globalTable.GrantReadData(user)
```

The `replica` method can be used to grant to a specific replica table:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import kms "github.com/aws/aws-cdk-go/awscdk"

var user User


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

tableKey := kms.NewKey(stack, jsii.String("Key"))
replicaKeyArns := map[string]*string{
	"us-east-1": jsii.String("arn:aws:kms:us-east-1:123456789012:key/g24efbna-az9b-42ro-m3bp-cq249l94fca6"),
	"us-east-2": jsii.String("arn:aws:kms:us-east-2:123456789012:key/g24efbna-az9b-42ro-m3bp-cq249l94fca6"),
}

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Encryption: dynamodb.TableEncryptionV2_CustomerManagedKey(tableKey, replicaKeyArns),
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
		},
	},
})

// grantReadData applies to the table in us-east-2 and the key arn for the key in us-east-2
globalTable.Replica(jsii.String("us-east-2")).GrantReadData(user)
```

## Metrics

You can use `metric*` methods to generate metrics for a table that can be used when configuring an `Alarm` or `Graphs`. The `metric*` methods only apply to the primary table provisioned using the `TableV2` construct. As an example, `metricConsumedReadCapacityUnits` used below is only for the table in `us-west-2`:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})

globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
		},
		&ReplicaTableProps{
			Region: jsii.String("us-east-2"),
		},
	},
})

// metric is only for the table in us-west-2
metric := globalTable.MetricConsumedReadCapacityUnits()

cloudwatch.NewAlarm(this, jsii.String("Alarm"), &AlarmProps{
	Metric: metric,
	EvaluationPeriods: jsii.Number(1),
	Threshold: jsii.Number(1),
})
```

The `replica` method can be used to generate a metric for a specific replica table:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"


type fooStack struct {
	Stack
	globalTable TableV2
}

func newFooStack(scope Construct, id *string, props StackProps) *fooStack {
	this := &fooStack{}
	cdk.NewStack_Override(this, scope, id, props)

	this.globalTable = dynamodb.NewTableV2(this, jsii.String("GlobalTable"), &TablePropsV2{
		PartitionKey: &Attribute{
			Name: jsii.String("pk"),
			Type: dynamodb.AttributeType_STRING,
		},
		Replicas: []ReplicaTableProps{
			&ReplicaTableProps{
				Region: jsii.String("us-east-1"),
			},
			&ReplicaTableProps{
				Region: jsii.String("us-east-2"),
			},
		},
	})
	return this
}

type barStackProps struct {
	StackProps
	replicaTable ITableV2
}

type barStack struct {
	Stack
}

func newBarStack(scope Construct, id *string, props barStackProps) *barStack {
	this := &barStack{}
	cdk.NewStack_Override(this, scope, id, props)

	// metric is only for the table in us-east-1
	metric := *props.replicaTable.MetricConsumedReadCapacityUnits()

	cloudwatch.NewAlarm(this, jsii.String("Alarm"), &AlarmProps{
		Metric: metric,
		EvaluationPeriods: jsii.Number(1),
		Threshold: jsii.Number(1),
	})
	return this
}

app := cdk.NewApp()
fooStack := NewFooStack(app, jsii.String("FooStack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-west-2"),
	},
})
barStack := NewBarStack(app, jsii.String("BarStack"), &barStackProps{
	replicaTable: fooStack.globalTable.Replica(jsii.String("us-east-1")),
	env: &Environment{
		Region: jsii.String("us-east-1"),
	},
})
```

## import from S3 Bucket

You can import data in S3 when creating a Table using the `Table` construct.
To import data into DynamoDB, it is required that your data is in a CSV, DynamoDB JSON, or Amazon Ion format within an Amazon S3 bucket.
The data may be compressed using ZSTD or GZIP formats, or you may choose to import it without compression.
The data source can be a single S3 object or multiple S3 objects sharing a common prefix.

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/S3DataImport.HowItWorks.html

### use CSV format

The `InputFormat.csv` method accepts `delimiter` and `headerList` options as arguments.
If `delimiter` is not specified, `,` is used by default.
And if `headerList` is specified, the first line of CSV is treated as data instead of header.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"

var bucket IBucket


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"))

dynamodb.NewTable(stack, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	ImportSource: &ImportSourceSpecification{
		CompressionType: dynamodb.InputCompressionType_GZIP,
		InputFormat: dynamodb.InputFormat_Csv(&CsvOptions{
			Delimiter: jsii.String(","),
			HeaderList: []*string{
				jsii.String("id"),
				jsii.String("name"),
			},
		}),
		Bucket: *Bucket,
		KeyPrefix: jsii.String("prefix"),
	},
})
```

### use DynamoDB JSON format

Use the `InputFormat.dynamoDBJson()` method to specify the `inputFormat` property.
There are currently no options available.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"

var bucket IBucket


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"))

dynamodb.NewTable(stack, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	ImportSource: &ImportSourceSpecification{
		CompressionType: dynamodb.InputCompressionType_GZIP,
		InputFormat: dynamodb.InputFormat_DynamoDBJson(),
		Bucket: *Bucket,
		KeyPrefix: jsii.String("prefix"),
	},
})
```

### use Amazon Ion format

Use the `InputFormat.ion()` method to specify the `inputFormat` property.
There are currently no options available.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"

var bucket IBucket


app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("Stack"))

dynamodb.NewTable(stack, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	ImportSource: &ImportSourceSpecification{
		CompressionType: dynamodb.InputCompressionType_GZIP,
		InputFormat: dynamodb.InputFormat_Ion(),
		Bucket: *Bucket,
		KeyPrefix: jsii.String("prefix"),
	},
})
```
