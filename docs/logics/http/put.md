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
| selfIntro         | String     |
| image             | Image?     |
| preference        | String     |
| followIDs         | List<int>  |
| groupIDs          | List<int>  |
| DirectMessageIDs  | List<int>  |

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

//クライアントは使用しない。

## /tags

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
