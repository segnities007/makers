# POST

## /user

curl -X POST "http://localhost:1323/user?email=email&password=password" | jq .


| プロパティ          | 型         |
|-------------------|------------|
| email             | String     |
| password          | String     |

## /userinfo

curl -X POST \
  'http://localhost:1323/userinfo?id=1&name=name&birthday=10101010&github=url&description=desc' \
  -F "image=@/home/segnities007/Pictures/image/segnities007.png" | jq .


| プロパティ          | 型         |
|-------------------|------------|
| name              | String     |
| birthday          | int?       |
| github            | String     |
| description       | String     |
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
| firstID             | int              |
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
