# github

The github module checks a GitHub project for the latest version.
It will connect to `https://api.github.com` to fetch the information from the V3 API.

## Parameters

| Name       | Optional | Type | Description                                       |
|:---------- |:--------:|:---- |:------------------------------------------------- |
| namespace  | No       | str  | Namespace of the project (project owner)          |
| project    | No       | str  | Name of the project in the namespace              |
| prerelease | Yes      | bool | Set to true to take the latest prerelease as well |
| draft      | Yes      | bool | Set to true to take the latest draft              |
