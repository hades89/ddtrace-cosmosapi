# ddtrace-cosmosapi

## Overview
This is a wrapper around the [Vippsas Azure CosmosDB Go SDK](https://github.com/vippsas/go-cosmosdb) that can be used to trace all the operations that are being executed by the CosmosDB client. This library is built on top of the [Datadog Go Tracer](https://github.com/DataDog/dd-trace-go).

## Installation
go get -u github.com/hades89/ddtrace-cosmosapi

## Getting Started
In order to start using ddtrace-cosmosapi, you can configure the tracer with the following code:

```
type CosmosStorageConfig struct {
	Endpoint string `yaml:"endpoint"`
	Database string `yaml:"database"`
	client   *tcc.TracedCosmosClient
}

func (c *CosmosStorageConfig) Open(password string) error {
	log.Debugf("Opening Cosmos storage adapter")

	cosmosCfg := cosmosapi.Config{
		MasterKey: password,
	}

	url := fmt.Sprintf("https://%s.documents.azure.com:443/", c.Endpoint)
	client := cosmosapi.New(url, cosmosCfg, nil, log.StandardLogger())

	c.client = ddtrace.NewTracedCosmosClient(client)
	return nil
}
```
By creating a `NewTracedCosmosClient`, you will be able to trace all the operations that are being executed by the CosmosDB client.
## Usage
Using ddtrace-cosmosapi is straightforward. You simply need to pass the context to your database methods, and the library will handle the rest, including setting up the necessary tracing.

Here's how you can use the library to trace a database query operation:
```
func (c *CosmosStorageConfig) GetSubscriptions(ctx context.Context) ([]datamodel.Subscription, error) {
	qops := cosmosapi.DefaultQueryDocumentOptions()
	qry := cosmosapi.Query{
		Query: selectAllQuery,
	}

	var docs []datamodel.Subscription
	for {
		var x []datamodel.Subscription
		
		// pass in the context here
		res, err := c.client.QueryDocuments(ctx, c.Database, subscriptionsColl, qry, &x, qops)
        ...
	}

	return docs, nil
}
```
## Contributions
Contributions are welcome! If you encounter any bugs, issues, or have feature requests, please open an issue. Pull requests are also appreciated.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
