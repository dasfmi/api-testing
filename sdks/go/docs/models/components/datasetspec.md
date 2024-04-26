# DatasetSpec


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Created`                                                | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | The RFC3339-formatted time when the dataset was created. |
| `Description`                                            | *string*                                                 | :heavy_check_mark:                                       | Dataset description                                      |
| `ID`                                                     | *string*                                                 | :heavy_check_mark:                                       | Dataset ID                                               |
| `Name`                                                   | *string*                                                 | :heavy_check_mark:                                       | Unique dataset name                                      |
| `Who`                                                    | *string*                                                 | :heavy_check_mark:                                       | Name of the dataset creator                              |