# ErrorDetail


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `SubCategory`                                                                                | **string*                                                                                    | :heavy_minus_sign:                                                                           | A specific category that contains more specific detail about the error                       |                                                                                              |
| `Code`                                                                                       | **string*                                                                                    | :heavy_minus_sign:                                                                           | The status code associated with the error detail                                             |                                                                                              |
| `In`                                                                                         | **string*                                                                                    | :heavy_minus_sign:                                                                           | The name of the field or parameter in which the error was found.                             |                                                                                              |
| `Context`                                                                                    | map[string][]*string*                                                                        | :heavy_minus_sign:                                                                           | Context about the error condition                                                            | {<br/>"missingScopes": [<br/>"scope1",<br/>"scope2"<br/>]<br/>}                              |
| `Message`                                                                                    | *string*                                                                                     | :heavy_check_mark:                                                                           | A human readable message describing the error along with remediation steps where appropriate |                                                                                              |