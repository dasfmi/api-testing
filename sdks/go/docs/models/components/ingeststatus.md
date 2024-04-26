# IngestStatus


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `BlocksCreated`                                                        | **int64*                                                               | :heavy_minus_sign:                                                     | Number of blocks created                                               |
| `Failed`                                                               | *int64*                                                                | :heavy_check_mark:                                                     | Number of failures that occurred during ingest                         |
| `Failures`                                                             | [][components.IngestFailure](../../models/components/ingestfailure.md) | :heavy_minus_sign:                                                     | List of failures that occurred during ingest                           |
| `Ingested`                                                             | *int64*                                                                | :heavy_check_mark:                                                     | Number of events ingested                                              |
| `ProcessedBytes`                                                       | *int64*                                                                | :heavy_check_mark:                                                     | Number of bytes processed                                              |
| `WalLength`                                                            | **int64*                                                               | :heavy_minus_sign:                                                     | Length of the WAL                                                      |