import 'models.dart';

//test data

final users = [
  User(id: 1, email: "test1@example.com", password: "password123", userInfoID: 1),
  User(id: 2, email: "test2@example.com", password: "password456", userInfoID: 2),
];

final userInfos = [
  UserInfo(
    id: 1,
    name: "Alice",
    birthday: 19950101,
    github: "https://github.com/alice",
    description: "Flutter developer",
    image: Image(id: 1, userID: 1, url: "https://example.com/image1.jpg"),
    tags: [1, 2],
    followIDs: [2],
    groupIDs: [1],
    dmIDs: [1],
  ),
  UserInfo(
    id: 2,
    name: "Bob",
    birthday: 19950202,
    github: "https://github.com/bob",
    description: "Backend developer",
    image: Image(id: 2, userID: 2, url: "https://example.com/image2.jpg"),
    tags: [2, 3],
    followIDs: [1],
    groupIDs: [1],
    dmIDs: [1],
  ),
];

final groups = [
  Group(id: 1, name: "Flutter Group", userIDs: [1, 2], messageIDs: [1, 2]),
];

final messages = [
  Message(id: 1, userID: 1, message: "Hello, Flutter!"),
  Message(id: 2, userID: 2, message: "Hi, Alice!"),
];

final directMessages = [
  DirectMessage(id: 1, firstID: 1, secondID: 2, messageIDs: [1, 2]),
];

final tags = [
  Tags(id: 1, tagID: 101, userIDs: [1], groupIDs: [1]),
  Tags(id: 2, tagID: 102, userIDs: [2], groupIDs: [1]),
];
