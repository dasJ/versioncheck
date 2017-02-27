# httphash

The httphash module does not know about versions.
Instead, it will do a HTTP request and hash the response body.
The encoded hash is treated as the version.
Once the website changes, httphash gets triggered.

This is useful on directory indexes containing the source archives.

## Parameters

| Name | Optional | Type | Description              |
|:---- |:--------:|:---- |:------------------------ |
| url  | No       | str  | The URL to be requested. |
