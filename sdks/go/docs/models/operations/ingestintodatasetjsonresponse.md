# IngestIntoDatasetJSONResponse


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `HTTPMeta`                                                           | [components.HTTPMetadata](../../models/components/httpmetadata.md)   | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `IngestStatus`                                                       | [*components.IngestStatus](../../models/components/ingeststatus.md)  | :heavy_minus_sign:                                                   | ingest status                                                        | {<br/>"ingested": 2,<br/>"failed": 0,<br/>"failures": [],<br/>"processedBytes": 16<br/>} |