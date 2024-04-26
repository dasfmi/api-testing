<!-- Start SDK Example Usage [usage] -->
```typescript
import { Axiom } from "axiom-js";

const axiom = new Axiom({
    bearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
});

async function run() {
    const id = "<value>";

    const result = await axiom.getAnnotation(id);

    // Handle the result
    console.log(result);
}

run();

```
<!-- End SDK Example Usage [usage] -->