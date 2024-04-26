# Annotations
(*annotations*)

## Overview

Annotations are used to mark important events in your data.

### Available Operations

* [getAnnotations](#getannotations) - Get annotations
* [createAnnotation](#createannotation) - Create a new annotation.

## getAnnotations

Get all annotations and filter by datasets or timerange.

### Example Usage

```typescript
import { Axiom } from "axiom-js";

const axiom = new Axiom({
  bearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const datasets = [
    "o",
    "n",
    "e",
    "-",
    "d",
    "a",
    "t",
    "a",
    "s",
    "e",
    "t",
    ",",
    "a",
    "n",
    "o",
    "t",
    "h",
    "e",
    "r",
    "-",
    "d",
    "a",
    "t",
    "a",
    "s",
    "e",
    "t",
  ];
  const start = new Date("2024-04-19T15:00:00Z");
  const end = new Date("2024-04-19T16:00:00Z");
  
  const result = await axiom.annotations.getAnnotations(datasets, start, end);

  // Handle the result
  console.log(result)
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    | Example                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `datasets`                                                                                                                                                                     | *string*[]                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                             | The datasets to filter by.                                                                                                                                                     | [object Object]                                                                                                                                                                |
| `start`                                                                                                                                                                        | [Date](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date)                                                                                  | :heavy_minus_sign:                                                                                                                                                             | If set, will filter to events after this date. Should be in RFC3339.                                                                                                           | [object Object]                                                                                                                                                                |
| `end`                                                                                                                                                                          | [Date](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date)                                                                                  | :heavy_minus_sign:                                                                                                                                                             | If set, will filter to events before this date. Should be in RFC3339.                                                                                                          | [object Object]                                                                                                                                                                |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |                                                                                                                                                                                |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |                                                                                                                                                                                |


### Response

**Promise<[operations.GetAnnotationsResponse](../../models/operations/getannotationsresponse.md)>**
### Errors

| Error Object    | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.SDKError | 4xx-5xx         | */*             |

## createAnnotation

Create a new annotation.

### Example Usage

```typescript
import { Axiom } from "axiom-js";

const axiom = new Axiom({
  bearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const result = await axiom.annotations.createAnnotation({
    title: "Production deployment",
    description: "A production deployment happened.",
    time: new Date("2024-04-19T15:00:00Z"),
    endTime: new Date("2024-04-19T16:00:00Z"),
    type: "deploy",
    url: "https://deployments.example.com/42",
    datasets: [
      "o",
      "n",
      "e",
      "-",
      "d",
      "a",
      "t",
      "a",
      "s",
      "e",
      "t",
      ",",
      "a",
      "n",
      "o",
      "t",
      "h",
      "e",
      "r",
      "-",
      "d",
      "a",
      "t",
      "a",
      "s",
      "e",
      "t",
    ],
  });

  // Handle the result
  console.log(result)
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [components.NewAnnotation](../../models/components/newannotation.md)                                                                                                           | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |


### Response

**Promise<[operations.CreateAnnotationResponse](../../models/operations/createannotationresponse.md)>**
### Errors

| Error Object    | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.SDKError | 4xx-5xx         | */*             |
