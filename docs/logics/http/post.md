# POST

## /user

| プロパティ          | 型         |
|-------------------|------------|
| email             | String     |
| password          | String     |


## /userinfo

| プロパティ          | 型         |
|-------------------|------------|
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
| name          | String           |
| userIDs       | List<int>        |
| messageIDs    | List<int>        |

## /image

| プロパティ      | 型               |
|---------------|------------------|
| userID        | int              |
| url           | String           |

## /message

| プロパティ      | 型               |
|---------------|------------------|
| userID        | int              |　
| message       | String           |

## /dm

| プロパティ           | 型               |
|--------------------|------------------|
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
| name          | String           |
