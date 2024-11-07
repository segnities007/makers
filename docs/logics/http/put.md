# PUT

## /user

| プロパティ          | 型         |
|-------------------|------------|
| id                | int        |
| email             | String     |
| password          | String     |


## /userinfo

| プロパティ          | 型         |
|-------------------|------------|
| id                | int        |
| name              | String     |
| birthday          | int?       |
| github            | String     |
| description       | String     |
| image             | Image?     |
| tags              | List<int>  |
| followIDs         | List<int>  |
| groupIDs          | List<int>  |
| directMessageIDs  | List<int>  |

## /group

| プロパティ      | 型               |
|---------------|------------------|
| id            | int              |
| name          | String           |
| userIDs       | List<int>        |
| messageIDs    | List<int>        |

## /image

| プロパティ      | 型               |
|---------------|------------------|
| id            | int              |
| userID        | int              |
| url           | String           |

## /message

| プロパティ      | 型               |
|---------------|------------------|
| id            | int              |
| userID        | int              |　
| message       | String           |

## /dm

| プロパティ           | 型               |
|--------------------|------------------|
| id                 | int              |
| firsID             | int              |
| secondID           | int              |
| messageIDs         | List<int>        |



## /tags
//多分これはいらない

| プロパティ      | 型               |
|---------------|------------------|
| id            | int              |
| tag           | Tag              |
| userIDs       | List<int>        |
| groupIDs      | List<int>        |

## /tag

| プロパティ      | 型               |
|---------------|------------------|
| id            | int              |
| name          | String           |
