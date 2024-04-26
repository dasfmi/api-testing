# Axiom SDK


## Overview

Axiom API: 100% of your data for every possible need: o11y, security, analytics, and new insights.


### Available Operations

* [getAnnotation](#getannotation) - Get a single annotation.
* [updateAnnotation](#updateannotation) - Update an annotation.
* [deleteAnnotation](#deleteannotation) - Delete an annotation.

## getAnnotation

Get a single annotation by id.

### Example Usage

```typescript
import { Axiom } from "axiom-js";

const axiom = new Axiom({
  bearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const id = "<value>";
  
  const result = await axiom.getAnnotation(id);

  // Handle the result
  console.log(result)
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | The id of the annotation.                                                                                                                                                      |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |


### Response

**Promise<[operations.GetAnnotationResponse](../../models/operations/getannotationresponse.md)>**
### Errors

| Error Object    | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.SDKError | 4xx-5xx         | */*             |

## updateAnnotation

Update an annotation.

### Example Usage

```typescript
import { Axiom } from "axiom-js";

const axiom = new Axiom({
  bearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const id = "<value>";
  const baseAnnotation = {
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
  };
  
  const result = await axiom.updateAnnotation(id, baseAnnotation);

  // Handle the result
  console.log(result)
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | The id of the annotation.                                                                                                                                                      |
| `baseAnnotation`                                                                                                                                                               | [components.BaseAnnotation](../../models/components/baseannotation.md)                                                                                                         | :heavy_check_mark:                                                                                                                                                             | The fields to update.                                                                                                                                                          |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |


### Response

**Promise<[operations.UpdateAnnotationResponse](../../models/operations/updateannotationresponse.md)>**
### Errors

| Error Object    | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.SDKError | 4xx-5xx         | */*             |

## deleteAnnotation

Delete an annotation by id.

### Example Usage

```typescript
import { Axiom } from "axiom-js";

const axiom = new Axiom({
  bearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
  const id = "<value>";
  
  const result = await axiom.deleteAnnotation(id);

  // Handle the result
  console.log(result)
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | The id of the annotation.                                                                                                                                                      |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |


### Response

**Promise<[operations.DeleteAnnotationResponse](../../models/operations/deleteannotationresponse.md)>**
### Errors

| Error Object    | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.SDKError | 4xx-5xx         | */*             |
